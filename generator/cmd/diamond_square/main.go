package main

import (
	"flag"
	"image/jpeg"
	"log"
	"os"
	"time"

	"github.com/meetmorrowsolonmars/openpgl/generator/internal/domain"
	"github.com/meetmorrowsolonmars/openpgl/generator/internal/domain/algorithms"
)

func main() {
	seed := flag.Int64("seed", time.Now().UnixNano(), "Seed of the generation algorithm.")
	size := flag.Int("size", 513, "Height map size. Should equal (2**n + 1).")
	path := flag.String("path", "map.jpeg", "Path to save height map.")
	flag.Parse()

	// create the diamond-square algorithm
	algorithm := algorithms.NewDiamondSquare(*seed)

	// generate a height map
	altitudes := algorithm.Generate(*size)

	// save the height map to the file
	converter := domain.NewMapToImageConverter()
	image, err := converter.Convert(*size, *size, altitudes)
	if err != nil {
		log.Fatal("can't convert map to image")
	}

	file, err := os.Create(*path)
	if err != nil {
		log.Fatal("can't create file")
	}

	defer func() {
		_ = file.Close()
	}()

	err = jpeg.Encode(file, image, nil)
	if err != nil {
		log.Fatal("can't encode image")
	}
}
