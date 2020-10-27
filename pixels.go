package goutils

import "image"

const darkness = 60

func checkDarkness(p Pixel) bool {

	var r, g, b int
	if p.R <= darkness {
		r = 1
	}
	if p.G <= darkness {
		g = 1
	}
	if p.B <= darkness {
		b = 1
	}
	if r+g+b == 3 {
		return true
	}
	return false
}

// GetPixels gets the bi-dimensional pixel array
func GetPixels(img image.Image) ([][]Pixel, error) {
	bounds := img.Bounds()
	width, height := bounds.Max.X, bounds.Max.Y

	var pixels [][]Pixel
	for y := 0; y < height; y++ {
		var row []Pixel
		for x := 0; x < width; x++ {
			row = append(row, rgbaToPixel(img.At(x, y).RGBA()))
		}
		pixels = append(pixels, row)
	}

	return pixels, nil
}

// img.At(x, y).RGBA() returns four uint32 values; we want a Pixel
func rgbaToPixel(r uint32, g uint32, b uint32, a uint32) Pixel {
	return Pixel{int(r / 257), int(g / 257), int(b / 257), int(a / 257)}
}

// Pixel struct of RGBA values
type Pixel struct {
	R int
	G int
	B int
	A int
}
