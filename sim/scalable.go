package sim

// Modifiers contains signals for the direction of a scalables
// change in size.
// Difference between Neutral and Value is the Gap
// Weight is how much the effect of this particular modifier to total growth
type Modifier struct {
	Weight  float64
	Value   float64
	Neutral float64
}

type Scaler interface {
	ComputeTotalGap([]*Modifier) float64
	ApplyChange(gap float64)
}

type Scalable struct {
	Size float64
	Step float64
}

func (s *Scalable) ComputeTotalGap([]*Modifier) float64 {
	return 0.0
}

func (s *Scalable) ApplyChange(total_gap float64) {
	s.Size = max(0, s.Size*(1+s.Step*total_gap))
}
