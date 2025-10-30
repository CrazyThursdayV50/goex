package account

import "github.com/CrazyThursdayV50/goex/binance/websocket-api/models"

// ParamsData 账户状态查询请求参数
type ParamsData struct {
	OmitZeroBalances bool `json:"omitZeroBalances"`
	models.Sign
}

func (d *ParamsData) Map() map[string]string {
	params := d.Sign.Map()
	if d.OmitZeroBalances {
		params["omitZeroBalances"] = "true"
	}

	return params
}

type Params = models.WsAPIParams[*ParamsData]

func NewParams() *Params {
	return &Params{
		Method: "account.status",
	}
}

// ResultData 账户状态查询响应数据
type ResultData struct {
	MakerCommission            int             `json:"makerCommission"`
	TakerCommission            int             `json:"takerCommission"`
	BuyerCommission            int             `json:"buyerCommission"`
	SellerCommission           int             `json:"sellerCommission"`
	CanTrade                   bool            `json:"canTrade"`
	CanWithdraw                bool            `json:"canWithdraw"`
	CanDeposit                 bool            `json:"canDeposit"`
	CommissionRates            CommissionRates `json:"commissionRates"`
	Brokered                   bool            `json:"brokered"`
	RequireSelfTradePrevention bool            `json:"requireSelfTradePrevention"`
	PreventSor                 bool            `json:"preventSor"`
	UpdateTime                 int64           `json:"updateTime"`
	AccountType                string          `json:"accountType"`
	Balances                   []Balance       `json:"balances"`
	Permissions                []string        `json:"permissions"`
	UID                        int64           `json:"uid"`
}

// CommissionRates 手续费率
type CommissionRates struct {
	Maker  string `json:"maker"`
	Taker  string `json:"taker"`
	Buyer  string `json:"buyer"`
	Seller string `json:"seller"`
}

// Balance 账户余额
type Balance struct {
	Asset  string `json:"asset"`
	Free   string `json:"free"`
	Locked string `json:"locked"`
}
