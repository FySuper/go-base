// methods
package main

import (
	"fmt"
	"math"
)

type Car struct {
	X, Y float64
}

func (v *Car) Go() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func main() {
	v := &Car{3, 4}
	fmt.Println(v.Go())
}
