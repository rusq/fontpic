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
		&IFMicrofont,
		&IFMicrofontBold,
		&IFMicrofontItalic,
		&IFMiliFont,
		&IFStupidSimple,
		&IFStupidSimpleBold,
		&IFStupidSimpleItalic,
	}
)

var (
	// These fonts are GPL2+ font by mibi88
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
	IFMicrofont = ImageFont{
		Name:        "microfont",
		Width:       4,
		GridSize:    image.Pt(4, 4),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}

	IFMicrofontBold = ImageFont{
		Name:        "microfont_bold",
		Width:       4,
		GridSize:    image.Pt(4, 4),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}
	IFMicrofontItalic = ImageFont{
		Name:        "microfont_italic",
		Width:       4,
		GridSize:    image.Pt(4, 4),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}

	// https://git.planet-casio.com/mibi88/microfont/src/branch/master/milifont.png
	// milifont.png:
	//	name: milifont
	//	type: font
	//	charset: print
	//	width: 3
	//	grid.size: 3x5
	//	grid.padding: 1
	//	proportional: false
	IFMiliFont = ImageFont{
		Name:        "milifont",
		Width:       3,
		GridSize:    image.Pt(3, 5),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}

	IFStupidSimple = ImageFont{
		Name:        "font",
		Width:       5,
		GridSize:    image.Pt(5, 5),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}
	IFStupidSimpleBold = ImageFont{
		Name:        "font_bold",
		Width:       5,
		GridSize:    image.Pt(5, 5),
		GridPadding: 1,
		CharStart:   32,
		Transparent: color.Transparent,
		CharEnd:     127,
	}
	IFStupidSimpleItalic = ImageFont{
		Name:        "font_italic",
		Width:       5,
		GridSize:    image.Pt(5, 5),
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
