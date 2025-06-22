package fontpic

import (
	"bytes"
	_ "embed"
	"image"
	"image/color"
)

//go:embed pngfnt/microfont.png
var microfontpng []byte

// MicroFont GPL2+ font by mibi88
// https://git.planet-casio.com/mibi88/microfont/src/branch/master/microfont.png
//
//	microfont.png:
//	  name: microfont
//	  type: font
//	  charset: print
//	  width: 4
//	  grid.size: 4x4
//	  grid.padding: 1
//	  proportional: false
var MicroFont = ImageFont{
	Name:        "microfont",
	Width:       4,
	GridSize:    image.Pt(4, 4),
	GridPadding: 1,
	CharStart:   32,
	Transparent: color.Transparent,
	CharEnd:     127,
}

func init() {
	MicroFont.load(bytes.NewReader(microfontpng))
}
