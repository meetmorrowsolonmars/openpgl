package algorithms

import (
	"math"
	"math/rand"
)

type DiamondSquare struct {
	generator *rand.Rand

	size  int
	side  int
	half  int
	area  float64
	ratio float64
}

func NewDiamondSquare(seed int64) *DiamondSquare {
	return &DiamondSquare{
		generator: rand.New(rand.NewSource(seed)),
	}
}

func (a *DiamondSquare) Generate(size int) [][]float64 {
	// calculate map area
	a.size = size
	a.area = float64(a.size * a.size)
	a.ratio = a.area
	a.half = 0
	a.side = 0

	// create map
	altitudes := make([][]float64, a.size)
	for i := range altitudes {
		altitudes[i] = make([]float64, a.size)
	}

	// set the altitude of the corners
	altitudes[0][0] = math.Abs(2 * a.rand())
	altitudes[0][a.size-1] = math.Abs(2 * a.rand())
	altitudes[a.size-1][0] = math.Abs(2 * a.rand())
	altitudes[a.size-1][a.size-1] = math.Abs(2 * a.rand())

	for a.side = a.size; a.side > 2; a.side = a.side/2 + 1 {
		// calculate area ratio
		a.ratio = float64(a.side*a.side) / a.area
		a.half = a.side / 2

		// diamond step
		for y := 0; y < a.size-1; y += a.side - 1 {
			for x := 0; x < a.size-1; x += a.side - 1 {
				altitudes[y+a.half][x+a.half] = a.diamond(y, x, altitudes)
			}
		}

		// square step
		for y := 0; y < a.size-1; y += a.side - 1 {
			for x := 0; x < a.size-1; x += a.side - 1 {
				altitudes[y][x+a.half] = a.square(y, x+a.half, altitudes)                   // top
				altitudes[y+a.half][x] = a.square(y+a.half, x, altitudes)                   // left
				altitudes[y+a.half][x+a.side-1] = a.square(y+a.half, x+a.side-1, altitudes) // right
				altitudes[y+a.side-1][x+a.half] = a.square(y+a.side-1, x+a.half, altitudes) // bottom
			}
		}
	}

	return altitudes
}

func (a *DiamondSquare) diamond(y, x int, altitudes [][]float64) float64 {
	avg := (altitudes[y][x] + altitudes[y][x+a.side-1] + altitudes[y+a.side-1][x] + altitudes[y+a.side-1][x+a.side-1]) / 4
	return math.Max(0, math.Min(1, a.rand()*a.ratio+avg))
}

func (a *DiamondSquare) square(y, x int, altitudes [][]float64) float64 {
	sum := 0.0
	c := 0.0

	// top
	if y-a.half >= 0 {
		sum += altitudes[y-a.half][x]
		c += 1
	}

	// left
	if x-a.half >= 0 {
		sum += altitudes[y][x-a.half]
		c += 1
	}

	// right
	if x+a.half < a.size {
		sum += altitudes[y][x+a.half]
		c += 1
	}

	// bottom
	if y+a.half < a.size {
		sum += altitudes[y+a.half][x]
		c += 1
	}

	return math.Max(0, math.Min(1, a.rand()*a.ratio+sum/c))
}

func (a *DiamondSquare) rand() float64 {
	return a.generator.Float64() - 0.5
}
