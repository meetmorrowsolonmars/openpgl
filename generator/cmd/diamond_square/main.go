package main

import (
	"fmt"
	"image"
	"image/color"
	"image/jpeg"
	"log"
	"math"
	"os"
	"time"

	"github.com/meetmorrowsolonmars/openpgl/generator/internal/domain/algorithms"
)

func main() {
	seed := time.Now().UnixNano()
	max := 1.0
	diamondSquare := algorithms.NewDiamondSquare(seed, max)

	size := 513
	altitudes := diamondSquare.Generate(size)

	mapImage := image.NewRGBA(image.Rect(0, 0, size, size))
	for y, row := range altitudes {
		for x, v := range row {
			altitude := uint8(math.Max(math.Min(255*v, 255), 0))

			mapImage.SetRGBA(x, y, color.RGBA{
				R: altitude,
				G: altitude,
				B: altitude,
				A: 255,
			})
		}
	}

	file, err := os.Create("test.jpeg")
	if err != nil {
		log.Fatal("can't create file")
	}

	defer func() {
		_ = file.Close()
	}()

	err = jpeg.Encode(file, mapImage, nil)
	if err != nil {
		log.Fatal("can't encode image")
	}
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
