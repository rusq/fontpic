package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	"strings"

	"golang.org/x/image/font"
	"golang.org/x/image/font/basicfont"
	"golang.org/x/image/math/fixed"

	"github.com/rusq/fontpic"
)

var fonts = []*basicfont.Face{
	basicfont.Face7x13,
	fontpic.Keyrus8x16,
	fontpic.Keyrus8x14,
	fontpic.Keyrus8x8,
	fontpic.Microfont,
}

func main() {
	for i := range fonts {
		f, err := os.Create(fmt.Sprintf("%d.png", i))
		if err != nil {
			log.Fatal(err)
		}
		if err := png.Encode(f, fonts[i].Mask); err != nil {
			log.Fatal(err)
		}
		f.Close()
		sample(fmt.Sprintf("sample%d.png", i), fonts[i])
	}
}

func sample(filename string, face *basicfont.Face) {
	f, err := os.Create(filename)
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()

	text := " !\"#$%&'()*+,-./0123456789:;<=>?\n" +
		"@ABCDEFGHIJKLMNOPQRSTUVWXYZ[\\]^_`abcdefghijklmnopqrstuvwxyz{|}~\n" +
		"The quick brown fox jumps over the lazy dog.\n" +
		"Now it is a good time for all fine men to come to an aid of their country.\n"

	lines := strings.Split(text, "\n")
	imgHeight := len(lines) * face.Metrics().Height.Ceil()
	// lineHeight := face.Metrics().Ascent.Ceil() + face.Metrics().Height.Ceil()

	fg, bg := image.Black, image.White
	img := image.NewRGBA(image.Rect(0, 0, 480, imgHeight))
	draw.Draw(img, img.Bounds(), bg, image.Point{}, draw.Src)

	var d = font.Drawer{
		Dst:  img,
		Src:  fg,
		Face: face,
		Dot:  fixed.P(0, face.Metrics().Ascent.Ceil()), // Start at the top
	}
	var replacer = strings.NewReplacer("\t", "        ")
	for _, line := range lines {
		d.DrawString(replacer.Replace(line))
		d.Dot.X = fixed.I(0)             // Reset X position to the start of the line
		d.Dot.Y += face.Metrics().Height // + face.Metrics().Ascent
	}
	if err := png.Encode(f, img); err != nil {
		log.Fatal(err)
	}
}
