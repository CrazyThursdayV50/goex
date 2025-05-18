package models

import (
	"fmt"
	"time"
)

type Sign struct {
	ApiKey    string `json:"apiKey"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

func (p *Sign) SetTimestamp() {
	p.Timestamp = time.Now().UnixMilli()
}

func (p *Sign) SetApiKey(apikey string) {
	p.ApiKey = apikey
}

func (p *Sign) SetSignature(signature string) {
	p.Signature = signature
}

func (p *Sign) Map() map[string]string {
	return map[string]string{
		"apiKey":    p.ApiKey,
		"timestamp": fmt.Sprintf("%d", p.Timestamp),
	}
}
