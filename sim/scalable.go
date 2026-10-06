package sim

type Scaler interface {
	ComputeGap() float64
	ApplyChange(gap float64)
}

type Scalable struct {
	Size float64
	Step float64
}

func (s *Scalable) ApplyChange(gap float64) {
	s.Size = max(0, s.Size*(1+s.Step*gap))
}
