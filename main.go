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

// vector math

func (v *Vector) GetDirection() *Point {
	return &Point{
		x: v.end.x - v.start.x,
		y: v.end.y - v.start.y,
		z: v.end.z - v.start.z,
	}
}

func (v *Vector) DirectionVector() *Vector {
	return &Vector{
		start: v.start,
		end:   v.GetDirection(),
	}
}

func (v *Vector) Length() float64 {
	return math.Sqrt(math.Pow(v.end.x-v.start.x, 2) + math.Pow(v.end.y-v.start.y, 2) + math.Pow(v.end.z-v.start.z, 2))
}

func (v *Vector) UnitVector() *Vector {
	direction := v.GetDirection()
	length := v.Length()
	return &Vector{
		start: v.start,
		end: &Point{
			x: v.start.x + direction.x/length,
			y: v.start.y + direction.y/length,
			z: v.start.z + direction.z/length,
		},
	}
}

func (v *Vector) Angle(v2 *Vector) float64 {
	return math.Acos(v.DotProduct(v2) / (v.Length() * v2.Length()))
}

func (v *Vector) DotProduct(v2 *Vector) float64 {
	direction1 := v.GetDirection()
	direction2 := v2.GetDirection()
	return float64(direction1.x*direction2.x + direction1.y*direction2.y + direction1.z*direction2.z)
}

func (v *Vector) CrossProduct(v2 *Vector) *Vector {
	direction1 := v.UnitVector().GetDirection()
	direction2 := v2.UnitVector().GetDirection()

	// get the angle between the two vectors
	angle := v.Angle(v2)

	// get the length of the cross product
	length := v.Length() * v2.Length() * math.Sin(angle)

	// get the direction of the cross product
	direction := &Point{
		x: direction1.y*direction2.z - direction1.z*direction2.y,
		y: direction1.z*direction2.x - direction1.x*direction2.z,
		z: direction1.x*direction2.y - direction1.y*direction2.x,
	}

	// return the cross product
	return &Vector{
		start: v.start,
		end: &Point{
			x: v.start.x + direction.x*length,
			y: v.start.y + direction.y*length,
			z: v.start.z + direction.z*length,
		},
	}
}
