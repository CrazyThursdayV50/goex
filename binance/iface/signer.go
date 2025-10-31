package iface

type SignerData interface {
	SetAPIKEY(string)
	SetTimestamp()
	SetSignature(string)
	Mapper
}

type Signer interface {
	Sign(SignerData)
}
