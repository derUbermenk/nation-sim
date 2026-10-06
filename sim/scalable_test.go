package sim

import "testing"

func TestScalable_ApplyChange_Grow(t *testing.T) {
	// given an initial size, a positive gap should grow the count
	scalable := Scalable{Size: 100, Step: 0.05}
	before := scalable.Size

	scalable.ApplyChange(0.6)

	if scalable.Size <= before {
		t.Errorf("expected Size to grow from %v, got %v", before, scalable.Size)
	}
}

func TestScalable_ApplyChange_Decline(t *testing.T) {
	// given an initial size, a negative gap should shrink the count
	scalable := Scalable{Size: 100, Step: 0.05}
	before := scalable.Size

	scalable.ApplyChange(-0.3)

	if scalable.Size >= before {
		t.Errorf("expected Size to decline from %v, got %v", before, scalable.Size)
	}
}

func TestScalable_ApplyChange_Neutral(t *testing.T) {
	// given an initial size, a zero gap should leave the count unchanged
	scalable := Scalable{Size: 100, Step: 0.05}
	before := scalable.Size

	scalable.ApplyChange(0.0)

	if scalable.Size != before {
		t.Errorf("expected Size to stay at %v, got %v", before, scalable.Size)
	}
}
