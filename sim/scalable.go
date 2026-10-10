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

func (m *Modifier) WeightedGap() float64 {
	return m.Weight * (m.Value - m.Neutral)
}

type Scaler interface {
	ComputeTotalGap([]*Modifier) float64
	ApplyChange(gap float64)
}

type Scalable struct {
	Size float64
	Step float64
}

func (s *Scalable) ComputeTotalGap(weightedGaps []float64) float64 {
	totalGap := 0.0
	for _, wG := range weightedGaps {
		totalGap += wG
	}
	return totalGap
}

func (s *Scalable) ApplyChange(totalGap float64) {
	s.Size = max(0, s.Size*(1+s.Step*totalGap))
}
