package models

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func hmacHashing(apiSecret string, data string) string {
	mac := hmac.New(sha256.New, []byte(apiSecret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

// WsAPIAuthParamsData WebSocket API 身份验证请求参数
type WsAPIAuthParamsData struct {
	RecvWindow int64 `json:"recvWindow,omitempty"`
	Sign
}

// Map 返回用于签名的参数map
func (p *WsAPIAuthParamsData) Map() map[string]string {
	params := p.Sign.Map()
	if p.RecvWindow > 0 {
		params["recvWindow"] = fmt.Sprintf("%d", p.RecvWindow)
	}
	return params
}

type WsAPIAuthParams = WsAPIParams[*WsAPIAuthParamsData]

func NewWsAPIAuthParams() *WsAPIAuthParams {
	return &WsAPIAuthParams{
		Method: "session.logon",
	}
}

// WsAPIAuthResultData WebSocket API 身份验证响应数据
type WsAPIAuthResultData struct {
	APIKey           string `json:"apiKey"`
	AuthorizedSince  int64  `json:"authorizedSince"`
	ConnectedSince   int64  `json:"connectedSince"`
	ReturnRateLimits bool   `json:"returnRateLimits"`
	ServerTime       int64  `json:"serverTime"`
	UserDataStream   bool   `json:"userDataStream"`
} 