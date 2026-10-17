package sim

import (
	"math"
	"sort"
	"testing"
)

const floatTolerance = 1e-9

func FloatsEqualWithTolerance(expected float64, got float64) bool {
	return math.Abs(expected-got) <= floatTolerance
}

func TestModifier_weightedGap(t *testing.T) {
	tests := []struct {
		name     string
		weight   float64
		value    float64
		neutral  float64
		expected float64
	}{
		{
			name:     "positive gap",
			weight:   0.5,
			value:    1.0,
			neutral:  0.9,
			expected: 0.05,
		},
		{
			name:     "negative gap",
			weight:   0.3,
			value:    0.5,
			neutral:  0.9,
			expected: -0.12,
		},
		{
			name:     "zero gap",
			weight:   0.2,
			value:    0.0,
			neutral:  0.0,
			expected: 0.0,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			m := &Modifier{
				Weight:  tt.weight,
				Value:   tt.value,
				Neutral: tt.neutral,
			}

			calculated := m.WeightedGap()

			if !FloatsEqualWithTolerance(tt.expected, calculated) {
				t.Errorf("weightedGap() = %v, want %v", calculated, tt.expected)
			}
		})
	}
}

func TestScalable_ApplyChange_Grow(t *testing.T) {
	// given an initial size, a positive gap should grow the count
	scalable := &Scalable{Size: 100, Step: 0.05}
	before := scalable.Size

	scalable.ApplyChange(0.6)

	if scalable.Size <= before {
		t.Errorf("expected Size to grow from %v, got %v", before, scalable.Size)
	}
}

func TestScalable_ApplyChange_Decline(t *testing.T) {
	// given an initial size, a negative gap should shrink the count
	scalable := &Scalable{Size: 100, Step: 0.05}
	before := scalable.Size

	scalable.ApplyChange(-0.3)

	if scalable.Size >= before {
		t.Errorf("expected Size to decline from %v, got %v", before, scalable.Size)
	}
}

func TestScalable_ApplyChange_Neutral(t *testing.T) {
	// given an initial size, a zero gap should leave the count unchanged
	scalable := &Scalable{Size: 100, Step: 0.05}
	before := scalable.Size

	scalable.ApplyChange(0.0)

	if scalable.Size != before {
		t.Errorf("expected Size to stay at %v, got %v", before, scalable.Size)
	}
}

func TestScalable_ComputeTotalGap(t *testing.T) {
	// given a modifier list it should create the modifiers
	scalable := &Scalable{Size: 100, Step: 0.05}
	weightedGaps := []float64{0.05, -0.12, 0.0}

	calculatedTotalGap := scalable.ComputeTotalGap(weightedGaps)
	expectedTotalGap := -0.07

	if !FloatsEqualWithTolerance(expectedTotalGap, calculatedTotalGap) {
		t.Errorf("Calculated total gap %v, not equal to expected total gap %v", calculatedTotalGap, expectedTotalGap)
	}
}

const (
	testFood    ModKind = "food"
	testHousing ModKind = "housing"
	testHealth  ModKind = "health"
)

func TestModifiers_WeightedGaps(t *testing.T) {
	tests := []struct {
		name      string
		modifiers Modifiers
		expected  []float64
	}{
		{
			name: "positive, negative and zero gaps",
			modifiers: Modifiers{
				testFood:    &Modifier{Weight: 0.5, Value: 1.0, Neutral: 0.9},
				testHousing: &Modifier{Weight: 0.3, Value: 0.5, Neutral: 0.9},
				testHealth:  &Modifier{Weight: 0.2, Value: 0.0, Neutral: 0.0},
			},
			expected: []float64{-0.12, 0.0, 0.05}, // sorted ascending
		},
		{
			name: "single modifier",
			modifiers: Modifiers{
				testFood: &Modifier{Weight: 0.5, Value: 1.0, Neutral: 0.9},
			},
			expected: []float64{0.05},
		},
		{
			name:      "empty modifiers",
			modifiers: Modifiers{},
			expected:  []float64{},
		},
		{
			name:      "nil modifiers",
			modifiers: nil,
			expected:  []float64{},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got := tt.modifiers.WeightedGaps()

			if len(got) != len(tt.expected) {
				t.Fatalf("WeightedGaps() returned %d gaps, want %d", len(got), len(tt.expected))
			}

			// map iteration order is random, so compare in sorted order
			sort.Float64s(got)

			for i := range tt.expected {
				if !FloatsEqualWithTolerance(tt.expected[i], got[i]) {
					t.Errorf("gap[%d] = %v, want %v", i, got[i], tt.expected[i])
				}
			}
		})
	}
}
