package fontpic

import (
	"image/color"
	"image/draw"
)

func fill(img draw.Image, col color.Color) {
	for i := 0; i < img.Bounds().Dx(); i++ {
		for j := 0; j < img.Bounds().Dy(); j++ {
			img.Set(i, j, col)
		}
	}
}
