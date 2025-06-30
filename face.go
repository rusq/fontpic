package fontpic

import (
	"image"
	"math/bits"

	"golang.org/x/image/font/basicfont"
)

// face.go contains compatibility code for the font package.

const (
	krStride = 8 // stride of the Keyrus font bitmaps
)

var (
	// Face8x8 is the Keyrus 8x8 face.
	Face8x8 = FntToFace(fntKr8x8, 8, 8)
	// Face8x14 is the Keyrus 8x14 font face.
	Face8x14 = FntToFace(fntKr8x14, 8, 14)
	// Face8x16 is the Keyrus 8x16 font face.
	Face8x16 = FntToFace(fntKr8x16, 8, 16)

	// Face4x4 is Microfont 4x4 font face.
	Face4x4 = &basicfont.Face{
		Advance: 5,
		Width:   8,
		Height:  5,
		Ascent:  4,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(fntMicrofont[0x20*5:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	// Face4x4Bold is Microfont Bold 4x4 font face.
	Face4x4Bold = &basicfont.Face{
		Advance: 5,
		Width:   8,
		Height:  5,
		Ascent:  4,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(fntMicrofontBold[0x20*5:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	// Face4x4Italic is Microfont Italic 4x4 font face.
	Face4x4Italic = &basicfont.Face{
		Advance: 5,
		Width:   8,
		Height:  5,
		Ascent:  4,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(fntMicrofontItalic[0x20*5:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	// Face4x5 is the Millifont 5x4 font face.
	Face4x5 = &basicfont.Face{
		Advance: 4,
		Width:   8,
		Height:  6,
		Ascent:  5,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(fntMilifont[0x20*6:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	// Face6x5 is the Stupid Simple font face.
	Face6x5 = &basicfont.Face{
		Advance: 6,
		Width:   8,
		Height:  6,
		Ascent:  5,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(fntStupidsimplefont[0x20*6:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	// Face6x5Bold is the Stupid Simple Bold font face.
	Face6x5Bold = &basicfont.Face{
		Advance: 6,
		Width:   8,
		Height:  6,
		Ascent:  5,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(fntStupidsimplefontBold[0x20*6:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	// Face6x5Italic is the Stupid Simple Italic font face.
	Face6x5Italic = &basicfont.Face{
		Advance: 6,
		Width:   8,
		Height:  6,
		Ascent:  5,
		Descent: 1,
		Left:    -4,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(fntStupidsimplefontItalic[0x20*6:]),
			Stride: 8,
			Rect:   image.Rectangle{Max: image.Point{8, 96 * 8}},
		},
		Ranges: basicfont.Face7x13.Ranges,
	}

	FaceRobotron = &basicfont.Face{
		Advance: 10,
		Width:   9,
		Height:  9,
		Ascent:  7,
		Descent: 2,
		Left:    7,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(uint16ToUint8Rev(robotronFnt)),
			Stride: krStride * 2,
			Rect:   image.Rectangle{Max: image.Point{16, 173 * 9}},
		},
		Ranges: []basicfont.Range{
			{Low: 32, High: 204, Offset: 0},
			{Low: '\ufffd', High: '\ufffe', Offset: 1},
		},
	}
)

const bitsPerByte = 8 // defined as constant to easily update, when this changes

// Bytes2pixels converts a bytes slice where each bit represents a pixel to a
// byte slice where each byte represents a pixel.  If the source bit is 1, the
// corresponding byte will be 0xff (white), and if the source pixel is 0, the
// corresponding byte will be 0x00 (black).
func Bytes2pixels(b []byte) []byte {
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

// FntToFace creates a basicfont.Face from fnt file data.  Width and height are
// the width and height of the font in pixels.  The data must be a valid fnt
// file.  It is designed to work with standard 8x8, 8x14 and 8x16 fonts with
// 256 characters, and might not work correctly with non-standard height and
// width values. For non-standard fonts, you might want to define your own face
// based on the basicfont.Face struct, and convert the font data to pixels
// using [Bytes2pixels] before assigning it to the Mask.Pix field.
func FntToFace(data []byte, width, height int) *basicfont.Face {
	var descent = 1
	if height > 8 {
		descent = 2
	}

	return &basicfont.Face{
		Advance: width,
		Width:   width - 1,
		Height:  height,
		Ascent:  height - descent,
		Descent: descent,
		Left:    0,
		Mask: &image.Alpha{
			Pix:    Bytes2pixels(data),
			Stride: bitsPerByte,
			Rect:   image.Rectangle{Max: image.Point{width, 255 * height}},
		},
		Ranges: []basicfont.Range{
			{Low: '\u0000', High: '\u00ff', Offset: 0},
			{Low: '\ufffd', High: '\ufffe', Offset: 1},
		},
	}
}

// For example, 0xAABB turns into 0xAA, 0xBB (big-endian).
func uint16ToUint8(data []uint16) []byte {
	var ret = make([]byte, len(data)*2)
	for i := range data {
		ret[i<<1] = uint8(data[i] >> 8)
		ret[i<<1+1] = uint8(data[i] & 0xff)
	}
	return ret
}

// Reverses the bits in addition to breaking by bytes.
func uint16ToUint8Rev(data []uint16) []byte {
	var ret = make([]byte, len(data)*2)
	for i := range data {
		x := bits.Reverse16(data[i])
		ret[i<<1] = uint8(x >> 8)
		ret[i<<1+1] = uint8(x & 0xff)
	}
	return ret
}
