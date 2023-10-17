package algorithms

import (
	"math/rand"
)

type DiamondSquare struct {
	max       float64
	generator *rand.Rand
}

func NewDiamondSquare(seed int64, max float64) *DiamondSquare {
	return &DiamondSquare{
		max:       max,
		generator: rand.New(rand.NewSource(seed)),
	}
}

func (a *DiamondSquare) Generate(size int) [][]float64 {
	// create map
	altitudes := make([][]float64, size)
	for i := range altitudes {
		altitudes[i] = make([]float64, size)
	}

	// set the altitude of the corners
	altitudes[0][0] = 0.5 * a.rand()
	altitudes[0][size-1] = 0.5 * a.rand()
	altitudes[size-1][0] = 0.5 * a.rand()
	altitudes[size-1][size-1] = 0.5 * a.rand()

	// calculate map area
	sMap := float64(size * size)

	for side := size; side > 2; side = side/2 + 1 {
		// calculate area ratio
		ratio := float64(side*side) / sMap
		half := side / 2

		// diamond step
		for y := 0; y < size-1; y += side - 1 {
			for x := 0; x < size-1; x += side - 1 {
				altitudes[y+half][x+half] = a.diamond(y, x, side, ratio, altitudes)
			}
		}

		// square step
		for y := 0; y < size-1; y += side - 1 {
			for x := 0; x < size-1; x += side - 1 {
				altitudes[y][x+half] = a.square(y, x+half, half, size, ratio, altitudes)               // top
				altitudes[y+half][x] = a.square(y+half, x, half, size, ratio, altitudes)               // left
				altitudes[y+half][x+side-1] = a.square(y+half, x+side-1, half, size, ratio, altitudes) // right
				altitudes[y+side-1][x+half] = a.square(y+side-1, x+half, half, size, ratio, altitudes) // bottom
			}
		}
	}

	return altitudes
}

func (a *DiamondSquare) diamond(y, x, side int, ratio float64, altitudes [][]float64) float64 {
	avg := (altitudes[y][x] + altitudes[y][x+side-1] + altitudes[y+side-1][x] + altitudes[y+side-1][x+side-1]) / 4
	return a.rand()*ratio + avg
}

func (a *DiamondSquare) square(y, x, half, size int, ratio float64, altitudes [][]float64) float64 {
	sum := 0.0
	c := 0.0

	// top
	if y-half >= 0 {
		sum += altitudes[y-half][x]
		c += 1
	}

	// left
	if x-half >= 0 {
		sum += altitudes[y][x-half]
		c += 1
	}

	// right
	if x+half < size {
		sum += altitudes[y][x+half]
		c += 1
	}

	// bottom
	if y+half < size {
		sum += altitudes[y+half][x]
		c += 1
	}

	return a.rand()*ratio + sum/c
}

func (a *DiamondSquare) rand() float64 {
	return a.max * a.generator.Float64()
}
