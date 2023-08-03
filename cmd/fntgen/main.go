// Command fntgen is just a sandbox to fuck around.
package main

import (
	"flag"
	"image"
	"image/color"
	"image/png"
	"io"
	"os"

	"github.com/rusq/fontpic"
)

var (
	fontfile = flag.String("f", "../../fnt/08x16.fnt", "font file")
	output   = flag.String("o", "fontpic.png", "output file")
)

func main() {
	flag.Parse()

	renderfile()
	RenderText("Hello from 1989")
	canvasrender()
}

func canvasrender() {
	c := fontpic.NewCanvas(fontpic.FontDefault)
	img := c.WithBackground(color.Black).
		WithForeground(color.White).
		Render([]byte("I believe everything that has been taught to me\ncan be taught to others if they are willing to learn.\n       -- Rev. Mychael Shane")).
		Image()
	if err := writePng("canvas.png", img); err != nil {
		panic(err)
	}
}

func RenderText(text string) {
	img := image.NewRGBA(image.Rect(0, 0, len(text)*fontpic.FontDefault.Width, fontpic.FontDefault.Height))
	fontpic.FontDefault.TextAt(img, 0, 0, []byte(text), color.White, color.Black)
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
	font, err := LoadFont(*fontfile)
	if err != nil {
		panic(err)
	}
	renderFont(font, 16)
}

func renderFont(font *fontpic.Font, perLine int) {
	if err := writePng(*output, font.Sample(perLine)); err != nil {
		panic(err)
	}
}

func LoadFont(filename string) (*fontpic.Font, error) {
	f, err := os.Open(filename)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	data, err := io.ReadAll(f)
	if err != nil {
		return nil, err
	}
	return fontpic.ToFont8(data)
}
