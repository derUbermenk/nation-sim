package sim

const (
	FirmModProfitability ModKind = "profitability"
)

type Firm struct {
	Scalable
	Modifiers
}

func NewFirm(size float64) *Pop {
	return &Pop{
		Scalable: Scalable{Size: size, Step: 0.05},
		Modifiers: Modifiers{
			FirmModProfitability: {Weight: 1.0, Value: 0.8, Neutral: 0.9},
		},
	}
}
