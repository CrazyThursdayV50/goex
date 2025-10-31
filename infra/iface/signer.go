package iface

type SignerFunc func(any) (map[string]any, error)

// type SignerData interface {
// 	SetAPIKEY(string)
// 	SetTimestamp()
// 	SetSignature(string)
// 	Mapper
// }

// type Signer interface {
// 	Sign(SignerData)
// }
