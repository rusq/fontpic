package fontpic

import (
	"image"
	"image/color"
	"image/draw"
)

func fill(img draw.Image, col color.Color) {
	draw.Draw(img, img.Bounds(), image.NewUniform(col), image.Point{}, draw.Src)
}
