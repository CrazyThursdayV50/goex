package models

import (
	"encoding/json"
	"os"
	"testing"
)

func TestFilters(t *testing.T) {
	// 读取 exchangeInfo.json 文件
	data, err := os.ReadFile("exchangeInfo.json")
	if err != nil {
		t.Fatalf("Failed to read exchangeInfo.json: %v", err)
	}

	// 定义一个结构体来解析交易所信息
	var exchangeInfo struct {
		Symbols []struct {
			Symbol  string  `json:"symbol"`
			Filters Filters `json:"filters"`
		} `json:"symbols"`
	}

	// 解析 JSON 数据
	if err := json.Unmarshal(data, &exchangeInfo); err != nil {
		t.Fatalf("Failed to parse exchangeInfo.json: %v", err)
	}

	// 遍历所有交易对的过滤器
	for _, symbol := range exchangeInfo.Symbols {
		t.Logf("Testing filters for symbol: %s", symbol.Symbol)
		
		// 检查是否成功解析了过滤器
		if len(symbol.Filters) == 0 {
			t.Errorf("No filters found for symbol %s", symbol.Symbol)
			continue
		}

		// 测试过滤器的序列化和反序列化
		for i, filter := range symbol.Filters {
			// 序列化
			data, err := json.Marshal(filter)
			if err != nil {
				t.Errorf("Failed to marshal filter %d for symbol %s: %v", i, symbol.Symbol, err)
				continue
			}

			// 反序列化
			parsedFilter, err := ParseFilter(data)
			if err != nil {
				t.Errorf("Failed to parse filter %d for symbol %s: %v", i, symbol.Symbol, err)
				continue
			}

			// 验证过滤器类型
			if parsedFilter.GetFilterType() != filter.GetFilterType() {
				t.Errorf("Filter type mismatch for symbol %s: expected %s, got %s",
					symbol.Symbol, filter.GetFilterType(), parsedFilter.GetFilterType())
			}

			// 再次序列化以验证数据完整性
			newData, err := json.Marshal(parsedFilter)
			if err != nil {
				t.Errorf("Failed to marshal parsed filter %d for symbol %s: %v", i, symbol.Symbol, err)
				continue
			}

			// 比较原始数据和重新序列化的数据
			var original, parsed map[string]interface{}
			if err := json.Unmarshal(data, &original); err != nil {
				t.Errorf("Failed to unmarshal original data for comparison: %v", err)
				continue
			}
			if err := json.Unmarshal(newData, &parsed); err != nil {
				t.Errorf("Failed to unmarshal new data for comparison: %v", err)
				continue
			}

			// 检查所有字段是否匹配
			for key, value := range original {
				if parsedValue, ok := parsed[key]; !ok {
					t.Errorf("Missing field %s in parsed filter for symbol %s", key, symbol.Symbol)
				} else if value != parsedValue {
					t.Errorf("Field value mismatch for %s in symbol %s: expected %v, got %v",
						key, symbol.Symbol, value, parsedValue)
				}
			}
		}
	}
} 