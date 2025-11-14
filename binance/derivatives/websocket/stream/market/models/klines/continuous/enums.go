package continuous

type ContractType string

// perpetual 永续合约
// current_quarter 当季交割合约
// next_quarter 次季交割合约
const (
	Perpetual      ContractType = "perpetual"
	CurrentQuarter ContractType = "current_quarter"
	NextQuarter    ContractType = "next_quarter"
)

func (c ContractType) String() string { return string(c) }
