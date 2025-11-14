package api

import (
	"context"

	userdata "github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/useData"
)

// 获取交易对配置

const (
	symbolConfigPath = "/fapi/v1/symbolConfig"
)

func (api *API) SymbolConfig(ctx context.Context, params *userdata.SymbolConfigData) (result userdata.SymbolConfigResult, err error) {
	api.userdata().get(ctx, symbolConfigPath, params, &result)
	return
}
