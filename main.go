package govecs

import (
	"fmt"
	"math"
)

// point struct

type Point struct {
	x, y, z float64
}

func NewPoint(x, y, z float64) *Point {
	return &Point{
		x: x,
		y: y,
		z: z,
	}
}

func (p *Point) String() string {
	return fmt.Sprintf("Point(%f, %f, %f)", p.x, p.y, p.z)
}

func (p *Point) GetX() float64 {
	return p.x
}

func (p *Point) GetY() float64 {
	return p.y
}

func (p *Point) GetZ() float64 {
	return p.z
}

// vector struct

type Vector struct {
	start *Point
	end   *Point
}

func NewVector(start, end *Point) *Vector {
	return &Vector{
		start: start,
		end:   end,
	}
}

func (v *Vector) String() string {
	return fmt.Sprintf("Vector(%d, %s, %s, %s)", v.start, v.end)
}

func (v *Vector) GetStart() *Point {
	return v.start
}

func (v *Vector) GetEnd() *Point {
	return v.end
}

// Returns the magnitude of the vector

func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(v.end.x-v.start.x, 2) + math.Pow(v.end.y-v.start.y, 2) + math.Pow(v.end.z-v.start.z, 2))
}

func (v *Vector) Magnitude() float64 {
	return v.Length()
}

// Returns the unit vector in the direction of the vector

func (v *Vector) UnitVector() *Point {
	length := v.Length()
	return &Point{
		x: (v.end.x - v.start.x) / length,
		y: (v.end.y - v.start.y) / length,
		z: (v.end.z - v.start.z) / length,
	}
}

// Returns the direction vector of the vector from the origin

func (v *Vector) DirectionVectorFromOrigin() *Point {
	// origin
	o := NewPoint(0, 0, 0)

	// direction point
	d := NewPoint(v.end.x-v.start.x, v.end.y-v.start.y, v.end.z-v.start.z)

	// return the direction vector
	return NewVector(o, d).UnitVector()
}

// Returns the dot product of the vector and another vector named other

func (v *Vector) DotProduct(other *Vector) float64 {
	// get the unit vectors
	unit1 := v.UnitVector()
	unit2 := other.UnitVector()

	// return the dot product
	return unit1.x*unit2.x + unit1.y*unit2.y + unit1.z*unit2.z
}

// Returns the angle between the given vector and another vector named other, in radians

func (v *Vector) Angle(other *Vector) float64 {
	return v.AngleRadian(other)
}

// Returns the angle between the given vector and another vector named other, in radians

func (c *Vector) AngleRadian(other *Vector) float64 {
	return math.Acos(c.DotProduct(other) / (c.Length() * other.Length()))
}

// Returns the angle between the given vector and another vector named other, in degrees

func (v *Vector) AngleDegree(other *Vector) float64 {
	return v.AngleRadian(other) * (180 / math.Pi)
}

// Returns the unit of the cross product of the vector and another vector named other

func (v *Vector) CrossProduct(other *Vector) *Point {
	// get the unit vectors
	unit1 := v.UnitVector()
	unit2 := other.UnitVector()

	// rthe direction point
	dp := &Point{
		x: unit1.y*unit2.z - unit1.z*unit2.y,
		y: unit1.z*unit2.x - unit1.x*unit2.z,
		z: unit1.x*unit2.y - unit1.y*unit2.x,
	}

	// return the direction vector
	return NewVector(v.start, dp).UnitVector()
}

// Returns the projection of the vector onto another vector named other

func (v *Vector) Project(other *Vector) *Vector {
	// get the unit vectors
	unit1 := v.UnitVector()
	unit2 := other.UnitVector()

	// get the dot product
	dot := unit1.x*unit2.x + unit1.y*unit2.y + unit1.z*unit2.z

	// get the length of the vector
	length := v.Length()

	// get the projection point
	p := &Point{
		x: v.start.x + dot*length*unit2.x,
		y: v.start.y + dot*length*unit2.y,
		z: v.start.z + dot*length*unit2.z,
	}

	// return the projection vector
	return NewVector(v.start, p)
}
