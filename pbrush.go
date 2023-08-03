package fontpic

import (
	"image"
	"image/color"
)

func fill(img *image.RGBA, col color.Color) {
	for i := 0; i < img.Rect.Dx(); i++ {
		for j := 0; j < img.Rect.Dy(); j++ {
			img.Set(i, j, col)
		}
	}
}
