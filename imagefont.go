package fontpic

import (
	"bytes"
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
)

// ImageFont represents a bitmap font loaded from an image file. A great
// example of such a font is the MicroFont, which is a 4x4 pixel font with a
// grid size of 5x4 pixels and a padding of 1 pixel
type ImageFont struct {
	Name        string
	Width       int
	GridSize    image.Point // Single Character dimensions
	GridPadding int         // amount of whitespace in the image between chars
	CharStart   byte        // ASCII code of the first character
	CharEnd     byte        // ASCII code of the last character
	Image       image.Image
	Transparent color.Color
	Chars       []image.Image
}

// Load the font.
func (f *ImageFont) load(r io.Reader) error {
	img, err := png.Decode(r)
	if err != nil {
		return err
	}
	mf := img.(*image.NRGBA)
	f.Image = mf
	f.Chars = make([]image.Image, MicroFont.CharEnd-MicroFont.CharStart+1)
	i := 0
	for y := 0; y < mf.Bounds().Dy(); y += MicroFont.GridSize.Y + MicroFont.GridPadding*2 {
		for x := 0; x < mf.Bounds().Dx(); x += MicroFont.GridSize.X + MicroFont.GridPadding*2 {
			c := mf.SubImage(image.Rect(
				x,
				y,
				x+MicroFont.GridSize.X+MicroFont.GridPadding*2,
				y+MicroFont.GridSize.Y+MicroFont.GridPadding*2,
			))
			MicroFont.Chars[i] = c
			i++
		}
	}
	return nil
}

func charOffset(c byte) int {
	return int(c - MicroFont.CharStart)
}

func (f *ImageFont) Char(c byte) image.Image {
	return f.Chars[charOffset(c)]
}

func (f *ImageFont) DrawChar(dst draw.Image, c byte, at image.Point, fg, bg color.Color) error {
	if c < f.CharStart || c > f.CharEnd {
		return fmt.Errorf("character out of range: %c", c)
	}
	src := f.Char(c)
	sp := src.Bounds().Min

	dstfg := dst.ColorModel().Convert(fg)
	dstbg := dst.ColorModel().Convert(bg)

	for dy := 0; dy < f.GridSize.Y+f.GridPadding*2; dy++ {
		for dx := 0; dx < f.GridSize.X+f.GridPadding*2; dx++ {
			srcColor := src.At(sp.X+dx, sp.Y+dy)
			if !colEq(srcColor, f.Transparent) {
				dst.Set(at.X+dx, at.Y+dy, dstfg)
			} else {
				dst.Set(at.X+dx, at.Y+dy, dstbg)
			}
		}
	}
	return nil
}

func colEq(a, b color.Color) bool {
	ar, ag, ab, aa := a.RGBA()
	br, bg, bb, ba := b.RGBA()
	return ar == br && ag == bg && ab == bb && aa == ba
}

func (f *ImageFont) WriteString(dst draw.Image, s string, at image.Point, fg, bg color.Color) error {
	for _, c := range s {
		if err := f.DrawChar(dst, byte(c), at, fg, bg); err != nil {
			return err
		}
		at.X += f.GridSize.X + f.GridPadding*2
	}
	return nil
}

func (f *ImageFont) XWidth(s string) int {
	return len(s) * (f.GridSize.X + f.GridPadding*2)
}

// ToBitmap converts the font a byte array. Each byte represents a horizontal
// line of pixels.  The first byte is the top row of the first character, the
// second byte is the second row of the first character, and so on.  Each bit
// in the byte represents a pixel, with the least significant bit being the
// leftmost pixel.  If the bit is set, the pixel is on, otherwise it is off.
// The array is indexed by the ASCII value of the character.
//
// ypad is how many lines of padding to add to the bottom of each character.
func (f *ImageFont) ToBitmap(ypad uint8) [256][]byte {
	const charsetSz = 256
	ypad = ypad & 0x07
	var bitmap [charsetSz][]byte
	for ch := range charsetSz {
		bitmap[ch] = make([]byte, f.GridSize.Y+int(ypad))
		if ch < int(f.CharStart) || ch > int(f.CharEnd) {
			continue
		}
		src := f.Char(byte(ch))
		sp := src.Bounds().Min
		pad := f.GridPadding
		for dy := pad; dy < f.GridSize.Y+pad; dy++ {
			for dx := pad; dx < f.GridSize.X+pad; dx++ {
				srcColor := src.At(sp.X+dx, sp.Y+dy)
				if !colEq(srcColor, f.Transparent) {
					bitmap[ch][dy-pad] |= 1 << uint(f.GridSize.X-dx)
				}
			}
		}
	}
	return bitmap
}

func (f *ImageFont) Bytes(ypad uint8) []byte {
	var buf bytes.Buffer
	f.WriteBitmap(&buf, ypad)
	return buf.Bytes()
}

func (f *ImageFont) WriteBitmap(w io.Writer, ypad uint8) error {
	bm := f.ToBitmap(ypad)
	for _, b := range bm {
		w.Write(b)
	}
	return nil
}
