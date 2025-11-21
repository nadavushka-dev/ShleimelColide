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

func LogOnSceen(screen *ebiten.Image, s string, c *color) error {
	initOnce.Do(
		func() {
			if err := initFont(); err != nil {
				log.Printf("Faild to create text face source: %v", err)
			}
		},
	)

	if textFaceSource == nil {
		return errors.New("font not initialized")
	}

	// Use default color if none provided
	defaultColor := color{r: 0, g: 255, b: 0, a: 255}
	if c == nil {
		c = &defaultColor
	}

	op := &text.DrawOptions{}
	op.ColorScale.Scale(c.r, c.g, c.b, c.a)
	op.GeoM.Translate(10, 10)
	face := &text.GoTextFace{Source: textFaceSource, Size: 12}

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
