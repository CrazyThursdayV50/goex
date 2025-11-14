package derivatives

import "github.com/CrazyThursdayV50/goex/binance/variables"

const urlBaseRestAPI = "https://fapi.binance.com"
const urlBaseRestAPITest = "https://demo-fapi.binance.com"

type rest struct{}

func Rest() rest {
	return rest{}
}

func (rest) Endpoint() string {
	if variables.IsTest() {
		return urlBaseRestAPITest
	}

	return urlBaseRestAPI
}
