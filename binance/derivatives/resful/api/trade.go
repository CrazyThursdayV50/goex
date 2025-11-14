package api

import (
	"context"

	"github.com/CrazyThursdayV50/goex/binance/derivatives/resful/models/trade"
)

const (
	marginTypePath = "/fapi/v1/marginType"
)

// 变换逐仓全仓模式
func (api *API) SetMarginType(ctx context.Context, params *trade.MarginTypeParams) (result *trade.MarginTypeResult, err error) {
	err = api.trade().post(ctx, marginTypePath, params, &result)
	return
}

const (
	openLeveragePath = "/fapi/v1/leverage"
)

// 设置开仓杠杆
func (api *API) SetOpenLeverage(ctx context.Context, params *trade.OpenLeverageParams) (result *trade.OpenLeverageResult, err error) {
	err = api.trade().post(ctx, openLeveragePath, params, &result)
	return
}
