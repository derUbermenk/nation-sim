package sim

type Buyable interface {
}

type Need struct {
	Supply          int
	EffectiveDemand int
	FillRate        int
}
