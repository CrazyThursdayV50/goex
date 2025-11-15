package markprice

import (
	"fmt"
	"strings"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/websocket/stream/market/models"
)

const (
	Event = "markPriceUpdate"
)

func StreamName(symbol string) string {
	return fmt.Sprintf("%s@markPrice@1s", strings.ToLower(symbol))
}

// Result 表示标记价格更新事件的数据结构
type Result struct {
	models.BaseResult
	MarkPriceData
}

type MarkPriceData struct {
	MarkPrice            string `json:"p"` // 标记价格
	IndexPrice           string `json:"i"` // 现货指数价格
	EstimatedSettlePrice string `json:"P"` // 预估结算价, 仅在结算前最后一小时有参考价值
	FundingRate          string `json:"r"` // 资金费率
	NextFundingTime      int64  `json:"T"` // 下次资金时间 (毫秒时间戳)
}
