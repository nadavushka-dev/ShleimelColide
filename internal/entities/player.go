package entities

import (
	"bytes"
	"image"
	"shleimel_colide/internal/animation"
	"shleimel_colide/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
)

type Player struct {
	Character
	IsCollide bool
}

const (
	MoveSpeed  = 3
	JumpHeight = 30
)

const (
	AnimRowIdle = iota
	AnimRowRun
	AnimRowJump
)

const (
	AnimFramesIdle = 4
	AnimFramesRun  = 8
)

const (
	AnimationFrameWidth  = 32
	AnimationFrameHeight = 32
)

func NewPlayer() (*Player, error) {

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
			Position: &Position{
				X: 0,
				Y: 0,
			},
		},
	)
	return &Player{Character: *c}, nil
}

func (p *Player) Update(key ebiten.Key) {
	width, height := ebiten.WindowSize()

	switch key {
	case ebiten.KeySpace:
		p.State.CurrentAnim.Row = AnimRowJump

		if p.State.Position.Y == p.State.idleYpos-JumpHeight {
			p.State.descending = true
		}

		if p.State.Position.Y >= p.State.idleYpos-JumpHeight {
			if p.State.descending {
				p.State.Position.Y += MoveSpeed
			} else {
				p.State.Position.Y -= MoveSpeed
			}
		}
		if p.State.Position.Y >= p.State.idleYpos {
			p.State.descending = false
			p.State.CurrentAnim.Row = AnimRowIdle
			p.State.CurrentAnim.FrameCount = AnimFramesIdle
		}

	case ebiten.KeyUp:
		p.State.CurrentAnim.Row = AnimRowRun
		p.State.CurrentAnim.FrameCount = AnimFramesRun
		if p.State.Position.Y >= -height/4 {
			p.State.Position.Y -= MoveSpeed
		}
		p.State.idleYpos = p.State.Position.Y

	case ebiten.KeyDown:
		p.State.CurrentAnim.Row = AnimRowRun
		p.State.CurrentAnim.FrameCount = AnimFramesRun
		if p.State.Position.Y <= height/4 {
			p.State.Position.Y += MoveSpeed
		}
		p.State.idleYpos = p.State.Position.Y

	case ebiten.KeyRight:
		p.State.flipped = false
		p.State.CurrentAnim.Row = AnimRowRun
		p.State.CurrentAnim.FrameCount = AnimFramesRun
		if p.State.Position.X <= width/4 {
			p.State.Position.X += MoveSpeed
		}

	case ebiten.KeyLeft:
		p.State.flipped = true
		p.State.CurrentAnim.Row = AnimRowRun
		p.State.CurrentAnim.FrameCount = AnimFramesRun
		if p.State.Position.X >= -width/4 {
			p.State.Position.X -= MoveSpeed
		}

	default:
		p.State.CurrentAnim.Row = AnimRowIdle
		p.State.CurrentAnim.FrameCount = AnimFramesIdle
		if p.State.Position.Y < p.State.idleYpos {
			p.State.Position.Y += MoveSpeed
		}
	}
}

func (p *Player) Draw(cfg config.Config, screen *ebiten.Image, tick int) {
	op := &ebiten.DrawImageOptions{}

	if p.State.flipped {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(p.State.CurrentAnim.FrameWidth), 0)
	}

	op.GeoM.Translate(
		-float64(p.State.CurrentAnim.FrameWidth)/2,
		-float64(p.State.CurrentAnim.FrameHeight)/2,
	)

	op.GeoM.Translate(
		float64(cfg.ScreenWidth/2+p.State.Position.X),
		float64(cfg.ScreenHeight/2+p.State.Position.Y),
	)

	i := (tick / cfg.TicksPerFrame) % p.State.CurrentAnim.FrameCount
	sx := cfg.FrameOX + i*p.State.CurrentAnim.FrameWidth
	sy := p.State.CurrentAnim.Row * p.State.CurrentAnim.FrameHeight
	screen.DrawImage(p.State.Sprite.SubImage(
		image.Rect(sx, sy, sx+p.State.CurrentAnim.FrameWidth, sy+p.State.CurrentAnim.FrameHeight),
	).(*ebiten.Image), op)
}
