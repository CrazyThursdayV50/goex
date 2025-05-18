package models

import (
	"os"
	"testing"

	"github.com/CrazyThursdayV50/pkgo/json"
)

func TestParseExchangeInfo(t *testing.T) {
	f, err := os.ReadFile("./exchangeInfo.json")
	if err != nil {
		panic(err)
	}

	var result WsAPIResult
	err = json.JSON().Unmarshal(f, &result)
	if err != nil {
		panic(err)
	}

	var data WsExchangeInfoResultData
	err = json.JSON().Unmarshal(result.Result, &data)
	if err != nil {
		panic(err)
	}
}
