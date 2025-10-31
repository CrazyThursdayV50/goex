package models

import (
	"fmt"
	"time"
)

// Sign 签名基础结构
type Sign struct {
	APIKEY    string `json:"apiKey"`
	Timestamp int64  `json:"timestamp"`
	Signature string `json:"signature"`
}

// implement Mapper
func (p *Sign) Map() map[string]string {
	return map[string]string{
		"apiKey":    p.APIKEY,
		"timestamp": fmt.Sprintf("%d", p.Timestamp),
	}
}

// implement SignerData
func (p *Sign) SetAPIKEY(key string) {
	p.APIKEY = key
}

func (p *Sign) SetTimestamp() {
	p.Timestamp = time.Now().UnixNano() / 1e6
}

func (p *Sign) SetSignature(sig string) {
	p.Signature = sig
}
