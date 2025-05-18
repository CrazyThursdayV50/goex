package signer

type SignerData interface {
	SetApiKey(string)
	SetTimestamp()
	SetSignature(string)
	Map() map[string]string
}

type Signer interface {
	Sign(SignerData)
}
