package models

// // Sign 签名基础结构
// type Sign struct {
// 	APIKEY    *string `json:"apiKey,omitempty"`
// 	Timestamp int64   `json:"timestamp"`
// 	Signature *string `json:"signature,omitempty"`
// }

// // implement Mapper
// func (p *Sign) Map() map[string]string {
// 	var m = map[string]string{
// 		"timestamp": fmt.Sprintf("%d", p.Timestamp),
// 	}

// 	if p.APIKEY != nil {
// 		m["apiKey"] = *p.APIKEY
// 	}

// 	return m
// }

// // implement SignerData
// func (p *Sign) SetAPIKEY(key string) {
// 	p.APIKEY = &key
// }

// func (p *Sign) SetTimestamp() {
// 	p.Timestamp = time.Now().UnixNano() / 1e6
// }

// func (p *Sign) SetSignature(sig string) {
// 	p.Signature = &sig
// }
