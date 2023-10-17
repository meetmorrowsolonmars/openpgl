package domain

import (
	"golang.org/x/exp/constraints"
)

type HeightMap[T constraints.Integer | constraints.Float] struct {
	altitudes []T
	width     int
	height    int
}

func NewHeightMap[T constraints.Integer | constraints.Float](width, height int) *HeightMap[T] {
	return &HeightMap[T]{
		altitudes: make([]T, width*height),
		width:     width,
		height:    height,
	}
}

func (m *HeightMap[T]) Get(x, y int) T {
	if x < 0 || y < 0 {
		panic("x or y can not be less than 0")
	}

	if m.height >= y || m.width >= x {
		panic("x can not be greater or equal to width and y can not be greater or equal to height")
	}

	return m.altitudes[y*m.width+x]
}

func (m *HeightMap[T]) Set(x, y int, altitude T) {
	if x < 0 || y < 0 {
		panic("x or y can not be less than 0")
	}

	if m.height >= y || m.width >= x {
		panic("x can not be greater or equal to width and y can not be greater or equal to height")
	}

	m.altitudes[y*m.width+x] = altitude
}
