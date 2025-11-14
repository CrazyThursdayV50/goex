package models

import (
	"fmt"
	"time"
)

// Sign 签名基础结构
type Sign struct {
	ApiKey    string `json:"apiKey"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

func (p *Sign) SetTimestamp() {
	p.Timestamp = time.Now().UnixMilli()
}

func (p *Sign) SetAPIKEY(apikey string) {
	p.ApiKey = apikey
}

func (p *Sign) SetSignature(signature string) {
	p.Signature = signature
}

func (p *Sign) Map() map[string]string {
	m := map[string]string{}

	if p.Timestamp != 0 {
		m["timestamp"] = fmt.Sprintf("%d", p.Timestamp)
	}

	if p.ApiKey != "" {
		m["apiKey"] = p.ApiKey
	}

	if p.Signature != "" {
		m["signature"] = p.Signature
	}

	return m
}
