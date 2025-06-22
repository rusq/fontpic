package fontpic

import (
	_ "embed"
	"image"
	"image/color"
	"image/png"
	"os"
	"testing"
)

func TestMicrofont(t *testing.T) {
	xsz := MicroFont.XWidth("Hello, World!")
	dst := image.NewPaletted(image.Rect(0, 0, xsz, 64), color.Palette{color.Black, color.White})
	MicroFont.DrawChar(dst, ' ', image.Pt(0, 0), color.Black, color.White)
	MicroFont.DrawChar(dst, '!', image.Pt(6, 0), color.Black, color.White)
	MicroFont.DrawChar(dst, 'A', image.Pt(0, 6), color.Black, color.White)
	MicroFont.DrawChar(dst, 'B', image.Pt(6, 6), color.Black, color.White)
	MicroFont.WriteString(dst, "Hello, World!", image.Pt(0, 12), color.Black, color.White)

	f, err := os.Create("microfont_test.png")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := png.Encode(f, dst); err != nil {
		t.Fatal(err)
	}
}

func TestFont_ToBitmap(t *testing.T) {
	fnt := MicroFont
	f, err := os.Create("microfont.fnt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()
	if err := fnt.WriteBitmap(f, 1); err != nil {
		t.Fatal(err)
	}
}
