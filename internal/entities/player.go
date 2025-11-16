package entities

import (
	"shleimel_colide/internal/animation"
	"shleimel_colide/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
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
	AnimFramesJump = 4
)

const (
	AnimationFrameWidth  = 32
	AnimationFrameHeight = 32
)

func NewPlayer() (*Player, error) {

	img, err := loadSprite()
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

func (p *Player) Update(keys []ebiten.Key, isJumpKeyPressed bool, cf config.Config) {
	isJumping := false
	isMoving := false

	if len(keys) == 0 {
		isJumping = false
		isMoving = false
		if p.State.Position.Y < p.State.idleYpos {
			p.State.Position.Y += MoveSpeed
		}
		p.State.CurrentAnim.Row = AnimRowIdle
		p.State.CurrentAnim.FrameCount = AnimFramesIdle
		return
	}

	for _, key := range keys {
		switch key {
		case ebiten.KeySpace:
			p.handleGravity(&isJumping)

		case ebiten.KeyUp, ebiten.KeyK:
			if isJumpKeyPressed {
				break
			}
			isMoving = true
			if p.State.Position.Y >= -cf.ScreenHeight/2 {
				p.State.Position.Y -= MoveSpeed
			}
			p.State.idleYpos = p.State.Position.Y

		case ebiten.KeyDown, ebiten.KeyJ:
			if isJumpKeyPressed {
				break
			}
			isMoving = true
			if p.State.Position.Y <= cf.ScreenHeight/2 {
				p.State.Position.Y += MoveSpeed
			}
			p.State.idleYpos = p.State.Position.Y

		case ebiten.KeyRight, ebiten.KeyL:
			p.State.flipped = false
			isMoving = true
			if p.State.Position.X <= cf.ScreenWidth/2 {
				p.State.Position.X += MoveSpeed
			}
			if (p.State.descending || !isJumpKeyPressed) && p.State.Position.Y < p.State.idleYpos {
				p.State.Position.Y += MoveSpeed
			}

		case ebiten.KeyLeft, ebiten.KeyR:
			p.State.flipped = true
			isMoving = true
			if p.State.Position.X >= -cf.ScreenWidth/2 {
				p.State.Position.X -= MoveSpeed
			}
			if (p.State.descending || !isJumpKeyPressed) && p.State.Position.Y < p.State.idleYpos {
				p.State.Position.Y += MoveSpeed
			}

		}

	}
	p.handleAnimationFrameCount(isJumping, isMoving)
}

func (p *Player) handleGravity(isJumping *bool) {
	*isJumping = true

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
	}
}

func (p *Player) handleAnimationFrameCount(isJumping, isMoving bool) {
	if isJumping {
		p.State.CurrentAnim.Row = AnimRowJump
		p.State.CurrentAnim.FrameCount = AnimFramesJump
	} else if isMoving {
		p.State.CurrentAnim.Row = AnimRowRun
		p.State.CurrentAnim.FrameCount = AnimFramesRun
	} else {
		p.State.CurrentAnim.Row = AnimRowIdle
		p.State.CurrentAnim.FrameCount = AnimFramesIdle
	}
}
