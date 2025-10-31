package models

// WsPingParams Ping请求参数
type WsPingParams = WsAPIParams[any]

func NewWsPingParams() *WsPingParams {
	return &WsPingParams{
		Method: "ping",
	}
} 