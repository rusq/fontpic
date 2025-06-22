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

type FNT struct {
	Width   int
	Height  int
	Charset string
	Chars   [CharsetSz][]byte
}

var (
	// Embedded fonts are taken from KeyRus by Dmitry Gurtyak. Charset: 866
	//
	//go:embed fnt/08x08.fnt
	fntKr8x8 []byte
	//go:embed fnt/08x14.fnt
	fntKr8x14 []byte
	//go:embed fnt/08x16.fnt
	fntKr8x16 []byte
	//go:embed fnt/microfont.fnt
	fntMicrofont []byte
	//go:embed fnt/microfont_bold.fnt
	fntMicrofontBold []byte
	//go:embed fnt/microfont_italic.fnt
	fntMicrofontItalic []byte
	//go:embed fnt/milifont.fnt
	fntMilifont []byte
	//go:embed fnt/font.fnt
	fntStupidsimplefont []byte
	//go:embed fnt/font_bold.fnt
	fntStupidsimplefontBold []byte
	//go:embed fnt/font_italic.fnt
	fntStupidsimplefontItalic []byte
)

var (
	// FontDefault is the default font.
	Fnt8x8     = Must(ToFntCharset(fntKr8x8, "866"))
	Fnt8x14    = Must(ToFntCharset(fntKr8x14, "866"))
	Fnt8x16    = Must(ToFntCharset(fntKr8x16, "866"))
	FntDefault = Fnt8x16
)

func Must(fnt *FNT, err error) *FNT {
	if err != nil {
		panic(err)
	}
	return fnt
}

// ToFnt8 is a shortcut for calling ToFont(b, 8).
//
// Usual slice sizes for 8-bit wide fonts:
//  1. 8x8 font: 2048 bytes (1x8bytes x 256)
//  2. 8x14 font: 3584 bytes (1x14bytes x 256)
//  3. 8x16 font: 4096 bytes (1x16bytes x 256)
func ToFnt8(b []byte) (*FNT, error) {
	return ToFnt(b, chrWidth)
}

// ToFnt converts byte data to a Font structure.  It detects the font height
// based on the slice size.
func ToFnt(b []byte, width int) (*FNT, error) {
	height := len(b) / CharsetSz
	return &FNT{
		Width:  width,
		Height: height,
		Chars:  toChars(b, width, height),
	}, nil
}

func LoadFnt(filename string, width int) (*FNT, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return ToFnt(data, width)
}

func (f *FNT) WriteTo(w io.Writer) (n int64, err error) {
	return io.Copy(w, bytes.NewReader(f.Bytes()))
}

func (f *FNT) Bytes() []byte {
	return toBytes(f.Chars)
}

func ToFntCharset(b []byte, charset string) (*FNT, error) {
	font, err := ToFnt8(b)
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
func (f *FNT) Sample(perLine int) image.Image {
	return f.sample(perLine, color.Gray{0xa8}, color.Black, image.Point{1, 1})
}

func (f *FNT) SampleColor(perLine int, fg, bg color.Color) image.Image {
	return f.sample(perLine, fg, bg, image.Point{1, 1})
}

// sample generates a font sample, with perLine characters, fg foreground and
// bg background colors,
func (f *FNT) sample(perLine int, fg, bg color.Color, spacing image.Point) image.Image {
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
