package entities

import (
	"bytes"
	"image"
	"log"
	"math/rand"
	"shleimel_colide/internal/animation"
	"shleimel_colide/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type Enemy struct {
	Character
}

func NewEnemy(conf config.Config) (*Enemy, error) {

	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		return nil, err
	}

	spawnOffset := 70*2 + AnimationFrameHeight
	spawnX := conf.ScreenWidth - spawnOffset
	spawnY := conf.ScreenHeight - spawnOffset + 50

	posX := rand.Intn(spawnX) - spawnX/2
	posY := rand.Intn(spawnY) - spawnY/2

	screenX := conf.ScreenWidth/2 + posX
	screenY := conf.ScreenHeight/2 + posY

	log.Printf("Enemy spawn: Config(%dx%d)       Offset=%d Range(%dx%d) Pos(%d,%d) Screen(%d   ,%d) SpriteY[%d-%d]",
		conf.ScreenWidth, conf.ScreenHeight,
		spawnOffset, spawnX, spawnY,
		posX, posY, screenX, screenY, screenY-16,
		screenY+16)

	c := NewCharacter(
		CharacterState{
			Sprite: ebiten.NewImageFromImage(img),
			CurrentAnim: animation.Anim{
				Row:         AnimRowIdle,
				FrameCount:  AnimFramesIdle,
				FrameWidth:  AnimationFrameWidth,
				FrameHeight: AnimationFrameHeight,
			},
			Position: &Position{
				X: posX,
				Y: posY,
			},
		},
	)
	return &Enemy{Character: *c}, nil
}

func (e *Enemy) Update(pPos Position) {
	ePos := e.State.Position
	if ePos.X != pPos.X {
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if ePos.X < pPos.X {
			ePos.X += 1
			e.State.flipped = false
		} else {
			ePos.X -= 1
			e.State.flipped = true
		}
	}

	if ePos.Y != pPos.Y {
		if ePos.Y < pPos.Y {
			ePos.Y += 1
		} else {
			ePos.Y -= 1
		}
	}

}
