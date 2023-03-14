package govecs

import (
	"testing"

	"github.com/cantoramann/govecs"
)

// constructor
func TestNewVector(t *testing.T) {

	// positive ints
	v := govecs.NewVector(govecs.NewPoint(1, 2, 3), govecs.NewPoint(4, 5, 6))
	if v == nil {
		t.Errorf("NewVector(NewPoint(1, 2, 3), NewPoint(4, 5, 6)) == nil")
	}

	// zeros
	v = govecs.NewVector(govecs.NewPoint(0, 0, 0), govecs.NewPoint(0, 0, 0))
	if v == nil {
		t.Errorf("NewVector(NewPoint(0, 0, 0), NewPoint(0, 0, 0)) == nil")
	}

	// negative ints
	v = govecs.NewVector(govecs.NewPoint(-1, -2, -3), govecs.NewPoint(-4, -5, -6))
	if v == nil {
		t.Errorf("NewVector(NewPoint(-1, -2, -3), NewPoint(-4, -5, -6)) == nil")
	}

	// positive floats
	v = govecs.NewVector(govecs.NewPoint(1.1, 2.2, 3.3), govecs.NewPoint(4.4, 5.5, 6.6))
	if v == nil {
		t.Errorf("NewVector(NewPoint(1.1, 2.2, 3.3), NewPoint(4.4, 5.5, 6.6)) == nil")
	}

	v = govecs.NewVector(govecs.NewPoint(1.11, 2.22, 3.33), govecs.NewPoint(4.44, 5.55, 6.66))
	if v == nil {
		t.Errorf("NewVector(NewPoint(1.11, 2.22, 3.33), NewPoint(4.44, 5.55, 6.66)) == nil")
	}

	v = govecs.NewVector(govecs.NewPoint(1.111, 2.222, 3.333), govecs.NewPoint(4.444, 5.555, 6.666))
	if v == nil {
		t.Errorf("NewVector(NewPoint(1.111, 2.222, 3.333), NewPoint(4.444, 5.555, 6.666)) == nil")
	}

	// negative floats
	v = govecs.NewVector(govecs.NewPoint(-1.1, -2.2, -3.3), govecs.NewPoint(-4.4, -5.5, -6.6))
	if v == nil {
		t.Errorf("NewVector(NewPoint(-1.1, -2.2, -3.3), NewPoint(-4.4, -5.5, -6.6)) == nil")
	}

	v = govecs.NewVector(govecs.NewPoint(-1.11, -2.22, -3.33), govecs.NewPoint(-4.44, -5.55, -6.66))
	if v == nil {
		t.Errorf("NewVector(NewPoint(-1.11, -2.22, -3.33), NewPoint(-4.44, -5.55, -6.66)) == nil")
	}

	v = govecs.NewVector(govecs.NewPoint(-1.111, -2.222, -3.333), govecs.NewPoint(-4.444, -5.555, -6.666))
	if v == nil {
		t.Errorf("NewVector(NewPoint(-1.111, -2.222, -3.333), NewPoint(-4.444, -5.555, -6.666)) == nil")
	}
}
