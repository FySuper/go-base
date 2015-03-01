// defer
package main

import (
	"fmt"
)

func main() {
	defer fmt.Println("world")
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
	fmt.Println("Hello World!")
}
