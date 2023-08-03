// Package fontpic is a package for rendering fonts to images.
package fontpic

import (
	"image"
	"image/color"
)

// RenderCharAt is a low level primitive that renders a character at the given
// position.
func RenderCharAt(img *image.RGBA, x0, y0, width, height int, bits []byte, hi color.Color, lo color.Color) {
	for y := 0; y < height; y++ {
		for x := 0; x < width; x++ {
			if bits[y]&(1<<uint(x)) != 0 {
				img.Set(((width-1)-x)+x0, y+y0, hi)
			} else {
				img.Set(((width-1)-x)+x0, y+y0, lo)
			}
		}
	}
}

func (f *Font) TextAt(img *image.RGBA, x, y int, text []byte, fg, bg color.Color) {
	NewCanvas(f).WithBackground(bg).WithForeground(fg).WithImage(img).RenderAt(text, image.Point{x, y})
	// for i := 0; i < len(text); i++ {
	// 	RenderCharAt(img, x+(i*f.Width), y, f.Width, f.Height, f.Chars[text[i]], fg, bg)
	// }
}
