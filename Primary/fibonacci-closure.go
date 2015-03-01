package main

import "fmt"

// fibonacci 函数会返回一个返回 int 的函数。
func fibonacci() func() int {
	a := 0
	b := 1
	i := 0
	return func() int {
		if i == 1 || i == 2 {
			return 1
		} else {
			c := a + b
			a = b
			b = c
			return c
		}
	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
