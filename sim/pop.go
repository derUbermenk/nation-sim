package sim

const (
	PopModQOL ModKind = "qol"
)

type Pop struct {
	Scalable
	Modifiers
}

func NewPop(size float64) *Pop {
	return &Pop{
		Scalable: Scalable{Size: size, Step: 0.05},
		Modifiers: Modifiers{
			PopModQOL: {Weight: 1.0, Value: 0.8, Neutral: 0.9},
		},
	}
}
