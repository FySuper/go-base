package main

import (
	"code.google.com/p/go-tour/pic"
)
func Pic(dx, dy int) [][]uint8 {
	L := [][]uint8{}
	
	for i := 0; i < dy ; i++ {
		L[i] = []uint8{}
		for j := 0; j < dx; j++ {
			L[i][j] = uint8((i + j) / 2))
		}
	}
	return L
}

func main() {
			
	pic.Show(Pic)
}
