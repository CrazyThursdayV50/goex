package signer

type Mapper interface {
	Map() map[string]string
}

type SignerData interface {
	SetApiKey(string)
	SetTimestamp()
	SetSignature(string)
	Mapper
}

type Signer interface {
	Sign(SignerData)
}
