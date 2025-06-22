package fontpic

import (
	_ "embed"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"os"
	"testing"
)

func TestMicrofont(t *testing.T) {
	for _, fnt := range allImageFonts {
		t.Run(fnt.Name, func(t *testing.T) {
			xsz := fnt.XWidth("Hello, World!")
			dst := image.NewPaletted(image.Rect(0, 0, xsz, 64), color.Palette{color.Black, color.White})
			draw.Draw(dst, image.Rect(0, 0, xsz, 64), image.White, image.Point{}, draw.Src)
			fnt.DrawChar(dst, ' ', image.Pt(0, 0), color.Black, color.White)
			fnt.DrawChar(dst, '!', image.Pt(6, 0), color.Black, color.White)
			fnt.DrawChar(dst, 'A', image.Pt(0, 6), color.Black, color.White)
			fnt.DrawChar(dst, 'B', image.Pt(6, 6), color.Black, color.White)
			fnt.WriteString(dst, "Hello, World!", image.Pt(0, 12), color.Black, color.White)

			f, err := os.Create(fmt.Sprintf("%s.png", fnt.Name))
			if err != nil {
				t.Fatal(err)
			}
			defer f.Close()
			if err := png.Encode(f, dst); err != nil {
				t.Fatal(err)
			}
		})
	}
}

func TestFont_ToBitmap(t *testing.T) {
	for _, fnt := range allImageFonts {
		f, err := os.Create(fmt.Sprintf("%s.fnt", fnt.Name))
		if err != nil {
			t.Fatal(err)
		}
		defer f.Close()
		if err := fnt.WriteBitmap(f, 1); err != nil {
			t.Fatal(err)
		}
	}
}
