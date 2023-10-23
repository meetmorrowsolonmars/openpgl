package main

import (
	"fmt"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/meetmorrowsolonmars/openpgl/generator/internal/domain"
	"github.com/meetmorrowsolonmars/openpgl/generator/internal/domain/algorithms"
)

func main() {
	seed := time.Now().UnixNano()
	diamondSquare := algorithms.NewDiamondSquare(seed)

	size := 1025
	altitudes := diamondSquare.Generate(size)
	converter := domain.NewMapToImageConverter()
	mapImage, err := converter.Convert(size, size, altitudes)
	if err != nil {
		log.Fatal("can't convert map to image")
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
