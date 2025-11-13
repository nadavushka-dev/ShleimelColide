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

func (e *Enemy) Update(key ebiten.Key) {
	width, height := ebiten.WindowSize()
	switch key {
	case ebiten.KeySpace:
		e.State.CurrentAnim.Row = AnimRowJump

		if e.State.Position.Y == e.State.idleYpos-JumpHeight {
			e.State.descending = true
		}

		if e.State.Position.Y >= e.State.idleYpos-JumpHeight {
			if e.State.descending {
				e.State.Position.Y += MoveSpeed
			} else {
				e.State.Position.Y -= MoveSpeed
			}
		}
		if e.State.Position.Y >= e.State.idleYpos {
			e.State.descending = false
			e.State.CurrentAnim.Row = AnimRowIdle
			e.State.CurrentAnim.FrameCount = AnimFramesIdle
		}

	case ebiten.KeyUp:
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.Position.Y >= -height/5 {
			e.State.Position.Y -= MoveSpeed
		}
		e.State.idleYpos = e.State.Position.Y

	case ebiten.KeyDown:
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.Position.Y <= height/8 {
			e.State.Position.Y += MoveSpeed
		}
		e.State.idleYpos = e.State.Position.Y

	case ebiten.KeyRight:
		e.State.flipped = false
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.Position.X <= width/5 {
			e.State.Position.X += MoveSpeed
		}

	case ebiten.KeyLeft:
		e.State.flipped = true
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.Position.X >= -width/5 {
			e.State.Position.X -= MoveSpeed
		}

	default:
		e.State.CurrentAnim.Row = AnimRowIdle
		e.State.CurrentAnim.FrameCount = AnimFramesIdle
		if e.State.Position.Y < e.State.idleYpos {
			e.State.Position.Y += MoveSpeed
		}
	}
}

func (e *Enemy) Draw(cfg config.Config, screen *ebiten.Image, tick int) {
	op := &ebiten.DrawImageOptions{}

	if e.State.flipped {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(e.State.CurrentAnim.FrameWidth), 0)
	}

	op.GeoM.Translate(
		-float64(e.State.CurrentAnim.FrameWidth)/2,
		-float64(e.State.CurrentAnim.FrameHeight)/2,
	)

	op.GeoM.Translate(
		float64(cfg.ScreenWidth/2+e.State.Position.X),
		float64(cfg.ScreenHeight/2+e.State.Position.Y),
	)

	i := (tick / cfg.TicksPerFrame) % e.State.CurrentAnim.FrameCount
	sx := cfg.FrameOX + i*e.State.CurrentAnim.FrameWidth
	sy := e.State.CurrentAnim.Row * e.State.CurrentAnim.FrameHeight
	screen.DrawImage(e.State.Sprite.SubImage(
		image.Rect(sx, sy, sx+e.State.CurrentAnim.FrameWidth, sy+e.State.CurrentAnim.FrameHeight),
	).(*ebiten.Image), op)
}
