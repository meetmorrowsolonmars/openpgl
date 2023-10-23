package domain

import (
	"errors"
	"image"
	"image/color"
)

type MapToImageConverter struct {
	colors [][]float64
}

func NewMapToImageConverter() *MapToImageConverter {
	return &MapToImageConverter{
		colors: [][]float64{
			//      v,       r,       g,       b
			{0.000000, 0.44314, 0.67059, 0.84706},
			{0.037037, 0.47451, 0.69804, 0.87059},
			{0.074074, 0.51765, 0.72549, 0.89020},
			{0.111111, 0.55294, 0.75686, 0.91765},
			{0.148148, 0.58824, 0.78824, 0.94118},
			{0.185185, 0.63137, 0.82353, 0.96863},
			{0.222222, 0.67451, 0.85882, 0.98431},
			{0.259259, 0.72549, 0.89020, 1.00000},
			{0.296296, 0.77647, 0.92549, 1.00000},
			{0.333333, 0.84706, 0.94902, 0.99608},
			{0.333333, 0.67451, 0.81569, 0.64706},
			{0.370370, 0.58039, 0.74902, 0.54510},
			{0.407407, 0.65882, 0.77647, 0.56078},
			{0.444444, 0.74118, 0.80000, 0.58824},
			{0.481481, 0.81961, 0.84314, 0.67059},
			{0.518519, 0.88235, 0.89412, 0.70980},
			{0.555556, 0.93725, 0.92157, 0.75294},
			{0.592593, 0.90980, 0.88235, 0.71373},
			{0.629630, 0.87059, 0.83922, 0.63922},
			{0.666667, 0.82745, 0.79216, 0.61569},
			{0.703704, 0.79216, 0.72549, 0.50980},
			{0.740741, 0.76471, 0.65490, 0.41961},
			{0.777778, 0.72549, 0.59608, 0.35294},
			{0.814815, 0.66667, 0.52941, 0.32549},
			{0.851852, 0.67451, 0.60392, 0.48627},
			{0.888889, 0.72941, 0.68235, 0.60392},
			{0.925926, 0.79216, 0.76471, 0.72157},
			{0.962963, 0.87843, 0.87059, 0.84706},
			{1.000000, 0.96078, 0.95686, 0.94902},
		},
	}
}

func (c *MapToImageConverter) Convert(width, height int, altitudes [][]float64) (image.Image, error) {
	if width <= 0 || height <= 0 {
		return nil, errors.New("invalid image size")
	}

	img := image.NewRGBA(image.Rect(0, 0, width, height))

	for y, row := range altitudes {
		if y >= height {
			return nil, errors.New("invalid image size")
		}

		for x, altitude := range row {
			if x >= width {
				return nil, errors.New("invalid image size")
			}

			img.SetRGBA(x, y, c.convertHeightToColor(altitude))
		}
	}

	return img, nil
}

func (c *MapToImageConverter) convertHeightToColor(altitude float64) color.RGBA {
	if altitude <= c.colors[0][0] {
		return color.RGBA{
			R: uint8(255 * c.colors[0][1]),
			G: uint8(255 * c.colors[0][2]),
			B: uint8(255 * c.colors[0][3]),
			A: 0xff,
		}
	}

	for i := 0; i < len(c.colors)-1; i++ {
		if c.colors[i][0] < altitude && altitude < c.colors[i+1][0] {
			return color.RGBA{
				R: uint8(255 * c.colors[i][1]),
				G: uint8(255 * c.colors[i][2]),
				B: uint8(255 * c.colors[i][3]),
				A: 0xff,
			}
		}
	}

	return color.RGBA{
		R: uint8(255 * c.colors[len(c.colors)-1][1]),
		G: uint8(255 * c.colors[len(c.colors)-1][2]),
		B: uint8(255 * c.colors[len(c.colors)-1][3]),
		A: 0xff,
	}
}
