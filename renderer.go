package fontpic

import (
	"bytes"
	"image"
	"image/color"
)

type Canvas struct {
	Width      int
	Height     int
	Background color.Color
	Foreground color.Color // Color to use for the font
	Font       *Font       // Font to use
	Spacing    image.Point // Spacing between chars
	image      *image.RGBA
}

func NewCanvas(defavlt *Font) *Canvas {
	return &Canvas{
		Font:       defavlt,
		Spacing:    image.Point{0, 0},
		Foreground: color.Gray{0xa8},
		Background: color.Black,
	}
}

func (c *Canvas) WithFont(font *Font) *Canvas {
	c.Font = font
	return c
}

func (c *Canvas) WithSpacing(x, y int) *Canvas {
	c.Spacing = image.Point{x, y}
	return c
}

func (c *Canvas) WithSize(w, h int) *Canvas {
	c.Width = w
	c.Height = h
	return c
}

func (c *Canvas) CalcSize(text []byte) *Canvas {
	if len(text) == 0 {
		c.Width = 720
		c.Height = 540 // 4:3
	}
	lines := bytes.Split(text, []byte("\n"))
	maxLineLen := 0
	for _, line := range lines {
		if len(line) > maxLineLen {
			maxLineLen = len(line)
		}
	}
	// account for spacing
	c.Width = (maxLineLen * c.Font.Width) + (c.Spacing.X * maxLineLen)
	c.Height = (len(lines) * c.Font.Height) + (c.Spacing.Y * len(lines))
	return c
}

func (c *Canvas) WithBackground(bg color.Color) *Canvas {
	c.Background = bg
	return c
}

func (c *Canvas) WithForeground(fg color.Color) *Canvas {
	c.Foreground = fg
	return c
}

func (c *Canvas) WithImage(img *image.RGBA) *Canvas {
	c.image = img
	c.Width = img.Bounds().Dx()
	c.Height = img.Bounds().Dy()
	return c
}

func (c *Canvas) Render(text []byte) *Canvas {
	return c.renderAt(text, image.Point{0, 0})
}

func (c *Canvas) RenderAt(text []byte, at image.Point) *Canvas {
	return c.renderAt(text, at)
}

func (c *Canvas) init(text []byte) {
	if c.Width == 0 || c.Height == 0 {
		c.CalcSize(text)
	}
	if c.image == nil {
		c.image = image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
		fill(c.image, c.Background)
	}
}

func (c *Canvas) renderAt(text []byte, at image.Point) *Canvas {
	c.init(text)

	lines := bytes.Split(text, []byte("\n"))
	for y, line := range lines {
		line = bytes.Trim(line, "\r\n")
		for x, ch := range line {
			RenderCharAt(
				c.image,
				at.X+(x*c.Font.Width)+(x*c.Spacing.X),
				at.Y+(y*c.Font.Height)+(y*c.Spacing.Y),
				c.Font.Width,
				c.Font.Height,
				c.Font.Chars[ch],
				c.Foreground,
				c.Background,
			)
		}
	}
	return c
}

func (c *Canvas) Image() *image.RGBA {
	c.init(nil)
	return c.image
}
