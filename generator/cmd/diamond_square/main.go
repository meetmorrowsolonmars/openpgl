package main

import (
	"fmt"
	"github.com/meetmorrowsolonmars/openpgl/generator/internal/domain/algorithms"
)

func main() {
	seed := int64(20)
	max := 2.0
	diamondSquare := algorithms.NewDiamondSquare(seed, max)

	size := 17
	altitudes := diamondSquare.Generate(size)
	PrintMap(size, size, altitudes)
}

func PrintMap(w, h int, altitudes [][]float64) {
	fmt.Print("     ")
	for i := 0; i < w; i++ {
		fmt.Printf("(%3d) ", i+1)
	}
	fmt.Print("\n")

	for y := 0; y < h; y++ {
		if y == 0 {
			fmt.Printf("(%2d) ", 1)
		}

		for x := 0; x < w; x++ {
			fmt.Printf("%5.2f ", altitudes[y][x])
		}

		fmt.Printf("\n(%2d) ", y+1)
	}

	fmt.Print("\n")
}
