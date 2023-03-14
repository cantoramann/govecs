package govecs

import (
	"testing"

	"github.com/cantoramann/govecs"
)

// constructor
func TestNewPoint(t *testing.T) {
	p := govecs.NewPoint(1, 2, 3)
	p2 := govecs.NewPoint(1, 2, 3)
	if p == nil || p2 == nil {
		t.Errorf("NewPoint(1, 2, 3) != NewPoint(1, 2, 3)")
	}

	p = govecs.NewPoint(0, 0, 0)
	p2 = govecs.NewPoint(0, 0, 0)
	if p == nil || p2 == nil {
		t.Errorf("NewPoint(1, 2, 3) != NewPoint(1, 2, 3)")
	}

	p = govecs.NewPoint(-1, -2, -3)
	p2 = govecs.NewPoint(-1, -2, -3)
	if p == nil || p2 == nil {
		t.Errorf("NewPoint(1, 2, 3) != NewPoint(1, 2, 3)")
	}

	p = govecs.NewPoint(1.1, 2.2, 3.3)
	p2 = govecs.NewPoint(1.1, 2.2, 3.3)
	if p == nil || p2 == nil {
		t.Errorf("NewPoint(1, 2, 3) != NewPoint(1, 2, 3)")
	}
}

func TestGetPoints(t *testing.T) {

	// test positive ints
	p := govecs.NewPoint(1, 2, 3)
	if p.GetX() != 1 {
		t.Errorf("p.GetX() != 1")
	}
	if p.GetY() != 2 {
		t.Errorf("p.GetY() != 2")
	}
	if p.GetZ() != 3 {
		t.Errorf("p.GetZ() != 3")
	}

	// test positive floats
	p = govecs.NewPoint(1.1, 2.2, 3.3)
	if p.GetX() != 1.1 {
		t.Errorf("p.GetX() != 1.1")
	}
	if p.GetY() != 2.2 {
		t.Errorf("p.GetY() != 2.2")
	}
	if p.GetZ() != 3.3 {
		t.Errorf("p.GetZ() != 3.3")
	}

	// test negative ints
	p = govecs.NewPoint(-1, -2, -3)
	if p.GetX() != -1 {
		t.Errorf("p.GetX() != -1")
	}
	if p.GetY() != -2 {
		t.Errorf("p.GetY() != -2")
	}
	if p.GetZ() != -3 {
		t.Errorf("p.GetZ() != -3")
	}

	// test negative floats
	p = govecs.NewPoint(-1.1, -2.2, -3.3)
	if p.GetX() != -1.1 {
		t.Errorf("p.GetX() != -1.1")
	}
	if p.GetY() != -2.2 {
		t.Errorf("p.GetY() != -2.2")
	}
	if p.GetZ() != -3.3 {
		t.Errorf("p.GetZ() != -3.3")
	}

	// test zeros
	p = govecs.NewPoint(0, 0, 0)
	if p.GetX() != 0 {
		t.Errorf("p.GetX() != 0")
	}
	if p.GetY() != 0 {
		t.Errorf("p.GetY() != 0")
	}
	if p.GetZ() != 0 {
		t.Errorf("p.GetZ() != 0")
	}
}
