package main

import (
    "fmt"
    "math"
)

// Define the Circle struct
type Circle struct {
    Radius float64
}

// Define the Area method for Circle
func (c Circle) Area() float64 {
    return math.Pi * c.Radius * c.Radius
}

func main() {
    c := Circle{Radius: 5}
    fmt.Printf("Area of the circle: %.2f\n", c.Area())
}
