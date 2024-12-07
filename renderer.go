package fontpic

import (
	"bytes"
	"image"
	"image/color"
	"image/draw"
)

// Canvas is a canvas that can be rendered to.
//
// The member function naming convention is the following:
//
//   - Render* - renders text or lines of text to the canvas
//   - With* - sets a property
//
// Render functions, that have Text in the name do the minimum amount of
// transformations to the text to make it renderable.  For example, tabs
// are replaced with spaces, and newlines are treated as line separators.
//
// Render functions without Text in the name, render the provided lines of
// text verbatim, so any \n or \t characters will be rendered as a characters
// from the supplied font.
//
// Render functions, that have At in the name, render the text at the
// specified location.
//
// Zero canvas value is usable.  It will use the default font, and will
// render the text in Grey (0xa8) on Black background, just like the good
// old days.
type Canvas struct {
	Width      int
	Height     int
	Background color.Color
	Foreground color.Color // Color to use for the font
	Font       *Font       // Font to use
	Spacing    image.Point // Spacing between characters.
	Scale      image.Point // scaling factor (not used yet)
	image      draw.Image
}

// NewCanvas creates the new canvas with the default font.
func NewCanvas(defavlt *Font) *Canvas {
	return &Canvas{
		Font:       defavlt,
		Spacing:    image.Point{0, 0},
		Foreground: color.Gray{0xa8},
		Background: color.Black,
		Scale:      image.Point{1, 1},
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

// ensure ensure that all dodgy variables are set to safe values.
func (c *Canvas) ensure() {
	if c.Scale.X < 1 {
		c.Scale.X = 1
	}
	if c.Scale.Y < 1 {
		c.Scale.Y = 1
	}
	if c.Spacing.X < 0 {
		c.Spacing.X = 0
	}
	if c.Spacing.Y < 0 {
		c.Spacing.Y = 0
	}
	if c.Font == nil {
		c.Font = FontDefault
	}
	if c.Foreground == nil {
		c.Foreground = color.Gray{0xa8}
	}
	if c.Background == nil {
		c.Background = color.Black
	}
}

const (
	DefaultWidth  = 720
	DefaultHeight = 540
)

// CalcSize calculates the size of the canvas based on the provided lines
// of text.
func (c *Canvas) CalcSize(lines [][]byte) *Canvas {
	c.ensure()
	if len(lines) == 0 {
		c.Width = DefaultWidth * c.Scale.X
		c.Height = DefaultHeight * c.Scale.Y // 4:3
		return c
	}
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

func (c *Canvas) WithImage(img draw.Image) *Canvas {
	c.image = img
	c.Width = img.Bounds().Dx()
	c.Height = img.Bounds().Dy()
	return c
}

// Render renders the lines of text to the canvas.
func (c *Canvas) Render(lines [][]byte) *Canvas {
	return c.renderAt(lines, image.Point{0, 0})
}

// RenderAt renders the lines of text at the specified location.
func (c *Canvas) RenderAt(lines [][]byte, at image.Point) *Canvas {
	return c.renderAt(lines, at)
}

// RenderText renders the text to the canvas.
func (c *Canvas) RenderText(text []byte) *Canvas {
	return c.renderTextAt(text, image.Point{0, 0})
}

// RenderTextAt renders the text at the specified location.
func (c *Canvas) RenderTextAt(text []byte, at image.Point) *Canvas {
	return c.renderTextAt(text, at)
}

func (c *Canvas) init(lines [][]byte) {
	c.ensure()
	if c.Width == 0 || c.Height == 0 {
		c.CalcSize(lines)
	}
	if c.image == nil {
		c.image = image.NewRGBA(image.Rect(0, 0, c.Width, c.Height))
		fill(c.image, c.Background)
	}
}

// renderTextAt renders the text at the specified location.  It assumes
// that lines are separated by \n, and wraps at the end of the line.
// It also replaces tabs with 8 spaces.
func (c *Canvas) renderTextAt(text []byte, at image.Point) *Canvas {
	lines := bytes.Split(text, []byte("\n"))
	for i := range lines {
		lines[i] = bytes.ReplaceAll(bytes.TrimRight(lines[i], "\r\n"), []byte("\t"), []byte("        "))
	}
	return c.renderAt(lines, at)
}

// renderAt renders the lines at the specified location.
func (c *Canvas) renderAt(lines [][]byte, at image.Point) *Canvas {
	c.init(lines)
	for y, line := range lines {
		for x, ch := range line {
			RenderCharAt(
				c.image,
				image.Point{
					X: at.X + (x * c.Font.Width) + (x * c.Spacing.X),
					Y: at.Y + (y * c.Font.Height) + (y * c.Spacing.Y),
				},
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

func (c *Canvas) Image() draw.Image {
	if c.image == nil {
		c.init(nil)
	}
	return c.image
}
