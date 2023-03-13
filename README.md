# Govecs

A lightweight, basic vector library for Go. It is designed to be simple and easy to use, and to be as fast as possible.

## Installation

```bash
go get github.com/cantoramann/govecs
```

## Usage

```go
package main

import (
    "fmt"

    "github.com/cantoramann/govecs"
)


func main() {
    // Create a new vector (NewVector takes two points, start and end points)
    v := govecs.NewVector(govecs.NewPoint(1, 2, 3), govecs.NewPoint(4, 5, 6))

    // Get the direction of the vector
    point := v.Direction()

    // Get the length of the vector
    fmt.Println(v.Length())

    // Get the unit vector
    vector := v.UnitVector()

    // Get the angle between two vectors
    angle := v.Angle(anotherVector)

    // Get the dot product of two vectors
    dotProduct := v.DotProduct(anotherVector)

    // Get the cross product of two vectors
    crossProduct := v.CrossProduct(anotherVector)

    // more coming soon...

}
```
