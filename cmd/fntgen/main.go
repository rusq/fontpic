// Command fntgen is just a sandbox to fuck around.
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/rusq/fontpic"
	"github.com/rusq/fontpic/charset"
)

var (
	fontfile  = flag.String("f", "../../fnt/08x16.fnt", "font file")
	fontWidth = flag.Int("w", 8, "font width")
	output    = flag.String("o", "fontpic.png", "output file")
)

func main() {
	flag.Parse()

	renderfile()
	RenderText(string(charset.CP866.Translate("Привет из 1989")))
	canvasrender()
}

func canvasrender() {
	font, err := fontpic.LoadFnt(*fontfile, *fontWidth)
	if err != nil {
		panic(err)
	}
	c := fontpic.NewCanvas(font)
	img := c.WithBackground(color.Black).
		WithForeground(color.White).
		RenderText([]byte("Fact of the day:\n\tAt some point in time some things will be\ndifferent from what they are today.")).
		Image()
	if err := writePng("canvas.png", img); err != nil {
		panic(err)
	}
}

func RenderText(text string) {
	img := image.NewRGBA(image.Rect(0, 0, len(text)*fontpic.FntDefault.Width, fontpic.FntDefault.Height))
	fontpic.FntDefault.TextAt(img, 0, 0, []byte(text), color.White, color.Black)
	if err := writePng("text.png", img); err != nil {
		panic(err)
	}
}

func writePng(filename string, img image.Image) error {
	f, err := os.Create(filename)
	if err != nil {
		return err
	}
	defer f.Close()
	if err := png.Encode(f, img); err != nil {
		return err
	}
	return nil
}

func renderfile() {
	font, err := fontpic.LoadFnt(*fontfile, *fontWidth)
	if err != nil {
		panic(err)
	}
	renderFont(font, 16)
}

func renderFont(font *fontpic.FNT, perLine int) {
	if err := writePng(*output, font.Sample(perLine)); err != nil {
		panic(err)
	}
}
