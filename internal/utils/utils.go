package utils

import (
	"bytes"
	"errors"
	"log"
	"shleimel_colide/internal/entities"
	"sync"

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

var (
	textFaceSource *text.GoTextFaceSource
	initOnce       sync.Once
)

func initFont() error {
	var err error
	textFaceSource, err = text.NewGoTextFaceSource(bytes.NewReader(fonts.PressStart2P_ttf))
	return err
}

func LogOnSceen(screen *ebiten.Image, s string, c *color, posX *float64, posY *float64, centerX bool, centerY bool) error {
	initOnce.Do(
		func() {
			if err := initFont(); err != nil {
				log.Printf("Failed to create text face source: %v", err)
			}
		},
	)

	if textFaceSource == nil {
		return errors.New("font not initialized")
	}

	defaultColor := color{r: 0, g: 255, b: 0, a: 255}
	if c == nil {
		c = &defaultColor
	}

	// Dynamic scaling based on screen size
	screenWidth := screen.Bounds().Dx()
	scale := float64(screenWidth) / 1920.0
	fontSize := 36.0 * scale

	face := &text.GoTextFace{
		Source: textFaceSource,
		Size:   fontSize,
	}

	boundsWidth, boundsHeight := text.Measure(s, face, 0)
	x := 10.00
	y := 10.00

	if centerX {
		x = float64(screenWidth)/2 - boundsWidth/2
	}
	if centerY {
		screenHeight := screen.Bounds().Dy()
		y = float64(screenHeight)/2 - boundsHeight/2
	}

	if posX != nil {
		x = *posX
	}

	if posY != nil {
		y = *posY
	}

	op := &text.DrawOptions{}
	op.ColorScale.Scale(c.r, c.g, c.b, c.a)
	op.Filter = ebiten.FilterLinear // Smooth scaling
	op.GeoM.Translate(x, y)

	text.Draw(screen, s, face, op)
	return nil
}

func CollisionDetection(ent1, ent2 entities.Collidable) bool {
	ent1xr, ent1xl, ent1yb, ent1yt := ent1.GetBoundaries()
	ent2xr, ent2xl, ent2yb, ent2yt := ent2.GetBoundaries()

	if ent1xr > ent2xl && ent1xl < ent2xr && ent1yt < ent2yb && ent1yb > ent2yt {
		return true
	}

	return false
}
