package fontpic

import (
	"image"

	"golang.org/x/image/font/basicfont"
)

// face.go contains compatibility code for the font package.

const (
	krStride = 8 // stride of the Keyrus font bitmaps
)

var (
	Keyrus8x8 = &basicfont.Face{
		Advance: 8,
		Width:   7,
		Height:  8,
		Ascent:  7,
		Descent: 1,
		Left:    0,
		Mask: &image.Alpha{
			Pix:    bytes2pixels(kr8x8),
			Stride: krStride,
			Rect:   image.Rectangle{Max: image.Point{8, 255 * 8}},
		},
		Ranges: []basicfont.Range{
			{Low: '\u0000', High: '\u00ff', Offset: 0},
			{Low: '\ufffd', High: '\ufffe', Offset: 1},
		},
	}

	Keyrus8x14 = &basicfont.Face{
		Advance: 8,
		Width:   7,
		Height:  14,
		Ascent:  12,
		Descent: 2,
		Left:    0,
		Mask: &image.Alpha{
			Pix:    bytes2pixels(kr8x14),
			Stride: krStride,
			Rect:   image.Rectangle{Max: image.Point{8, 255 * 14}},
		},
		Ranges: []basicfont.Range{
			{Low: '\u0000', High: '\u00ff', Offset: 0},
			{Low: '\ufffd', High: '\ufffe', Offset: 1},
		},
	}

	Keyrus8x16 = &basicfont.Face{
		Advance: 8,
		Width:   7,
		Height:  16,
		Ascent:  14,
		Descent: 2,
		Left:    0,
		Mask: &image.Alpha{
			Pix:    bytes2pixels(kr8x16),
			Stride: krStride,
			Rect:   image.Rectangle{Max: image.Point{8, 255 * 16}},
		},
		Ranges: []basicfont.Range{
			{Low: '\u0000', High: '\u00ff', Offset: 0},
			{Low: '\ufffd', High: '\ufffe', Offset: 1},
		},
	}

	Microfont = &basicfont.Face{
		Advance: 5,
		Width:   8,
		Height:  5,
		Ascent:  4,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    bytes2pixels(microfont[0x20*5:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	MicrofontBold = &basicfont.Face{
		Advance: 5,
		Width:   8,
		Height:  5,
		Ascent:  4,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    bytes2pixels(microfontBold[0x20*5:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	MicrofontItalic = &basicfont.Face{
		Advance: 5,
		Width:   8,
		Height:  5,
		Ascent:  4,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    bytes2pixels(microfontItalic[0x20*5:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	Milifont = &basicfont.Face{
		Advance: 4,
		Width:   8,
		Height:  6,
		Ascent:  5,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    bytes2pixels(milifont[0x20*6:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}
)

const bitsPerByte = 8 // defined as constant to easily update, when this changes

// bitToBytes converts a bytes slice where each bit represents a pixel
// to a byte slice where each byte represents a pixel.  If the source
// bit is 1, the corresponding byte will be 0xff (white), and if
// the source pixel is 0, the corresponding byte will be 0x00 (black).
func bytes2pixels(b []byte) []byte {
	var pixels = make([]byte, 0, len(b)*bitsPerByte)
	for i := range b {
		p := bits2bytes(b[i])
		pixels = append(pixels, p[:]...)
	}
	return pixels
}

func bits2bytes(b byte) (p [8]byte) {
	p[0] = 0xff * (b >> 7 & 1)
	p[1] = 0xff * (b >> 6 & 1)
	p[2] = 0xff * (b >> 5 & 1)
	p[3] = 0xff * (b >> 4 & 1)
	p[4] = 0xff * (b >> 3 & 1)
	p[5] = 0xff * (b >> 2 & 1)
	p[6] = 0xff * (b >> 1 & 1)
	p[7] = 0xff * (b >> 0 & 1)
	return
}
