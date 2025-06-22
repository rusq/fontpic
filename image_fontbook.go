package fontpic

import (
	"embed"
	"image"
	"image/color"
	"path"
)

var (
	//go:embed imgfonts/*.png
	imgfontsFS embed.FS

	allImageFonts = []*ImageFont{
		&MicroFont,
		&MicroFontBold,
		&MicroFontItalic,
		&MiliFont,
	}
)

var (
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
	MicroFont = ImageFont{
		Name:        "microfont",
		Width:       4,
		GridSize:    image.Pt(4, 4),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}

	MicroFontBold = ImageFont{
		Name:        "microfont_bold",
		Width:       4,
		GridSize:    image.Pt(4, 4),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}
	MicroFontItalic = ImageFont{
		Name:        "microfont_italic",
		Width:       4,
		GridSize:    image.Pt(4, 4),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}
	// MiliFont GPL2+ font by mibi88
	// https://git.planet-casio.com/mibi88/microfont/src/branch/master/milifont.png
	// milifont.png:
	//	name: milifont
	//	type: font
	//	charset: print
	//	width: 3
	//	grid.size: 3x5
	//	grid.padding: 1
	//	proportional: false
	MiliFont = ImageFont{
		Name:        "milifont",
		Width:       3,
		GridSize:    image.Pt(3, 5),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}
)

func init() {
	for _, fnt := range allImageFonts {
		f, err := imgfontsFS.Open(path.Join("imgfonts", fnt.Name+".png"))
		if err != nil {
			panic(err)
		}
		if err := fnt.Load(f); err != nil {
			panic(err)
		}
		f.Close()
	}
}
