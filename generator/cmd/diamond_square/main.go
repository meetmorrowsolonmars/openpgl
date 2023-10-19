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
	diamondSquare := algorithms.NewDiamondSquare(seed)

	size := 513
	altitudes := diamondSquare.Generate(size)

	min, max := math.MaxFloat64, math.SmallestNonzeroFloat64

	mapImage := image.NewRGBA(image.Rect(0, 0, size, size))
	for y, row := range altitudes {
		for x, v := range row {
			var c color.RGBA
			switch {
			case v < 0.2:
				// sea
				c = color.RGBA{
					R: 0x13,
					G: 0x93,
					B: 0xf5,
					A: 0xff,
				}
			case v < 0.4:
				// sand
				c = color.RGBA{
					R: 0xf5,
					G: 0xbc,
					B: 0x2b,
					A: 0xff,
				}
			case v > 0.4:
				// forest
				c = color.RGBA{
					R: 0x0d,
					G: 0xa1,
					B: 0x12,
					A: 0xff,
				}
			}

			mapImage.SetRGBA(x, y, c)

			min = math.Min(min, v)
			max = math.Max(max, v)
		}
	}

	fmt.Printf("min %f, max %f\n", min, max)

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
