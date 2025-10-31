package auth

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"

	"github.com/CrazyThursdayV50/goex/binance/spot/websocket/api/models"
)

func hmacHashing(apiSecret string, data string) string {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// ParamsData WebSocket API 身份验证请求参数
type ParamsData struct {
	RecvWindow int64 `json:"recvWindow,omitempty"`
}

func NewRequest() *models.WsAPIRequest {
	return &models.WsAPIRequest{
		Method: "session.logon",
	}
}

// ResultData WebSocket API 身份验证响应数据
type ResultData struct {
	APIKey           string `json:"apiKey"`
	AuthorizedSince  int64  `json:"authorizedSince"`
	ConnectedSince   int64  `json:"connectedSince"`
	ReturnRateLimits bool   `json:"returnRateLimits"`
	ServerTime       int64  `json:"serverTime"`
	UserDataStream   bool   `json:"userDataStream"`
}
