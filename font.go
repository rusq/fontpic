package fontpic

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
	"io"
	"os"
)

const (
	// CharsetSz is the number of characters in a font file.  It is unlikely to
	// ever change.
	CharsetSz = 256
	chrWidth  = 8 // character width in bits
)

type Font struct {
	Width   int
	Height  int
	Charset string
	Chars   [CharsetSz][]byte
}

var (
	// Embedded fonts are taken from KeyRus by Dmitry Gurtyak. Charset: 866
	//
	//go:embed fnt/08x08.fnt
	kr8x8 []byte
	//go:embed fnt/08x14.fnt
	kr8x14 []byte
	//go:embed fnt/08x16.fnt
	kr8x16 []byte
)

var (
	// FontDefault is the default font.
	Font8x8     = Must(ToFontCharset(kr8x8, "866"))
	Font8x14    = Must(ToFontCharset(kr8x14, "866"))
	Font8x16    = Must(ToFontCharset(kr8x16, "866"))
	FontDefault = Font8x16
)

func Must(font *Font, err error) *Font {
	if err != nil {
		panic(err)
	}
	return font
}

// ToFont8 is a shortcut for calling ToFont(b, 8).
//
// Usual slice sizes for 8-bit wide fonts:
//  1. 8x8 font: 2048 bytes (1x8bytes x 256)
//  2. 8x14 font: 3584 bytes (1x14bytes x 256)
//  3. 8x16 font: 4096 bytes (1x16bytes x 256)
func ToFont8(b []byte) (*Font, error) {
	return ToFont(b, chrWidth)
}

// ToFont8 converts byte data to a Font structure.  It detects the font height
// based on the slice size.
func ToFont(b []byte, width int) (*Font, error) {
	height := len(b) / CharsetSz
	return &Font{
		Width:  width,
		Height: height,
		Chars:  toChars(b, width, height),
	}, nil
}

func LoadFont(filename string, width int) (*Font, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return ToFont(data, width)
}

func (f *Font) WriteTo(w io.Writer) (n int64, err error) {
	return io.Copy(w, bytes.NewReader(f.Bytes()))
}

func (f *Font) Bytes() []byte {
	return toBytes(f.Chars)
}

func ToFontCharset(b []byte, charset string) (*Font, error) {
	font, err := ToFont8(b)
	if err != nil {
		return nil, err
	}
	font.Charset = charset
	return font, nil
}

// charStride returns the number of bytes required to store a character of the
// given width.  The width is in bits.  If width is 0, it is assumed to be 8.
func charStride(width int) int {
	if width == 0 {
		width = chrWidth
	}
	return (width + 7) / 8
}

func toChars(fnt []byte, width int, height int) [CharsetSz][]byte {
	var chars [CharsetSz][]byte
	wb := charStride(width)
	for i := 0; i < CharsetSz; i++ {
		chars[i] = fnt[i*wb*height : (i+1)*wb*height]
	}
	return chars
}

func toBytes(chars [CharsetSz][]byte) []byte {
	var data = make([]byte, 0, CharsetSz*16)
	for i := 0; i < CharsetSz; i++ {
		data = append(data, chars[i]...)
	}
	return data
}

// Sample renders a sample of the font.  The font is rendered in a grid of
// perLine characters.  The sx and sy parameters are space between characters in
// pixels
func (f *Font) Sample(perLine int) image.Image {
	return f.sample(perLine, color.Gray{0xa8}, color.Black, image.Point{1, 1})
}

func (f *Font) SampleColor(perLine int, fg, bg color.Color) image.Image {
	return f.sample(perLine, fg, bg, image.Point{1, 1})
}

// sample generates a font sample, with perLine characters, fg foreground and bg background colors,
func (f *Font) sample(perLine int, fg, bg color.Color, spacing image.Point) image.Image {
	perY := CharsetSz / perLine
	img := image.NewRGBA(image.Rect(0, 0, f.Width*perLine+(spacing.X*perLine), f.Height*perY+(spacing.Y*perY)))
	fill(img, bg)
	for i := 0; i < len(f.Chars); i++ {
		x := (i%perLine)*f.Width + spacing.X*(i%perLine)
		y := (i/perLine)*f.Height + spacing.Y*(i/perLine)
		RenderCharAt(img, image.Point{x, y}, f.Width, f.Height, f.Chars[i], fg, bg)
	}
	return img
}
