package utils

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/fonts"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type color struct {
	r float32
	g float32
	b float32
	a float32
}

func LogOnSceen(screen *ebiten.Image, s string, c *color) {
	t, err := text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))

	if err != nil {
		log.Fatal(err)
	}

	// Use default color if none provided
	defaultColor := color{r: 0, g: 255, b: 0, a: 255}
	if c == nil {
		c = &defaultColor
	}

	op := &text.DrawOptions{}
	op.ColorScale.Scale(c.r, c.g, c.b, c.a)
	op.GeoM.Translate(10, 10)
	face := &text.GoTextFace{Source: t, Size: 12}

	text.Draw(screen, s, face, op)
}
