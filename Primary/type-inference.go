// type-inference.go
package main

import (
	"fmt"
	"math/cmplx"
)

func main() {
	v := cmplx.Sqrt(4 + 16i)
	fmt.Printf("v is of type %T\n", v)
}
