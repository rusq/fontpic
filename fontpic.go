// Package fontpic is a package for rendering fonts to images.
package fontpic

import (
	"image"
	"image/color"
	"image/draw"
)

// RenderCharAt is a low level function that renders a character, defined in
// bits, at the given position on the image.  It uses width and height to know
// how to render the character in bits.
func RenderCharAt(img draw.Image, at image.Point, width, height int, bits []byte, hi color.Color, lo color.Color) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if bits[y]&(1<<uint(x)) != 0 {
				img.Set(((width-1)-x)+at.X, y+at.Y, hi)
			} else {
				img.Set(((width-1)-x)+at.X, y+at.Y, lo)
			}
		}
	}
}

func (f *FNT) TextAt(img draw.Image, x, y int, text []byte, fg, bg color.Color) {
	NewCanvas(f).WithBackground(bg).WithForeground(fg).WithImage(img).RenderTextAt(text, image.Point{x, y})
	// for i := 0; i < len(text); i++ {
	// 	RenderCharAt(img, x+(i*f.Width), y, f.Width, f.Height, f.Chars[text[i]], fg, bg)
	// }
}
