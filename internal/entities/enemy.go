package entities

import (
	"bytes"
	"image"
	"shleimel_colide/internal/animation"
	"shleimel_colide/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type Enemy struct {
	Character
}

func NewEnemy() (*Enemy, error) {

	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		return nil, err
	}

	c := NewCharacter(
		CharacterState{
			Sprite: ebiten.NewImageFromImage(img),
			CurrentAnim: animation.Anim{
				Row:         AnimRowIdle,
				FrameCount:  AnimFramesIdle,
				FrameWidth:  AnimationFrameWidth,
				FrameHeight: AnimationFrameHeight,
			},
			position: &Position{
				X: 100,
				Y: 100,
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

		if e.State.position.Y == e.State.actualYPos-JumpHeight {
			e.State.descending = true
		}

		if e.State.position.Y >= e.State.actualYPos-JumpHeight {
			if e.State.descending {
				e.State.position.Y += MoveSpeed
			} else {
				e.State.position.Y -= MoveSpeed
			}
		}
		if e.State.position.Y >= e.State.actualYPos {
			e.State.descending = false
			e.State.CurrentAnim.Row = AnimRowIdle
			e.State.CurrentAnim.FrameCount = AnimFramesIdle
		}

	case ebiten.KeyUp:
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.position.Y >= -height/5 {
			e.State.position.Y -= MoveSpeed
		}
		e.State.actualYPos = e.State.position.Y

	case ebiten.KeyDown:
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.position.Y <= height/8 {
			e.State.position.Y += MoveSpeed
		}
		e.State.actualYPos = e.State.position.Y

	case ebiten.KeyRight:
		e.State.flipped = false
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.position.X <= width/5 {
			e.State.position.X += MoveSpeed
		}

	case ebiten.KeyLeft:
		e.State.flipped = true
		e.State.CurrentAnim.Row = AnimRowRun
		e.State.CurrentAnim.FrameCount = AnimFramesRun
		if e.State.position.X >= -width/5 {
			e.State.position.X -= MoveSpeed
		}

	default:
		e.State.CurrentAnim.Row = AnimRowIdle
		e.State.CurrentAnim.FrameCount = AnimFramesIdle
		if e.State.position.Y < e.State.actualYPos {
			e.State.position.Y += MoveSpeed
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
		float64(cfg.ScreenWidth/2+e.State.position.X),
		float64(cfg.ScreenHeight/2+e.State.position.Y),
	)

	i := (tick / cfg.TicksPerFrame) % e.State.CurrentAnim.FrameCount
	sx := cfg.FrameOX + i*e.State.CurrentAnim.FrameWidth
	sy := e.State.CurrentAnim.Row * e.State.CurrentAnim.FrameHeight
	screen.DrawImage(e.State.Sprite.SubImage(
		image.Rect(sx, sy, sx+e.State.CurrentAnim.FrameWidth, sy+e.State.CurrentAnim.FrameHeight),
	).(*ebiten.Image), op)
}
