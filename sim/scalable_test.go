package sim

import "testing"

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
	modifiers := []*Modifier{
		{Weight: 0.5, Value: 1.0, Neutral: 0.9}, // gap +0.1
		{Weight: 0.3, Value: 0.5, Neutral: 0.9}, // gap -0.4
		{Weight: 0.2, Value: 0.0, Neutral: 0.0}, // gap  0.0
	}

	scalable := &Scalable{Size: 100, Step: 0.05}

	calculated_total_gap := scalable.ComputeTotalGap(modifiers)
	expected_total_gap := -0.3

	if calculated_total_gap != expected_total_gap {
		t.Errorf("Calculated total gap %v, not equal to expected total gap %v", calculated_total_gap, expected_total_gap)
	}
}
