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

type UpdateInput struct {
	Keys             []ebiten.Key
	IsJumpKeyPressed bool
	Cf               config.Config
}

func (p *Player) Update(input UpdateInput) {
	isJumping := false
	isMoving := false

	if len(input.Keys) == 0 {
		isJumping = false
		isMoving = false
		if p.State.Position.Y < p.State.idleYpos {
			p.State.Position.Y += MoveSpeed
		}
		p.State.CurrentAnim.Row = AnimRowIdle
		p.State.CurrentAnim.FrameCount = AnimFramesIdle
		return
	}

	for _, key := range input.Keys {
		switch key {
		case ebiten.KeySpace:
			p.handleGravity(&isJumping)

		case ebiten.KeyUp:
			if input.IsJumpKeyPressed {
				break
			}
			isMoving = true
			if p.State.Position.Y >= -input.Cf.ScreenHeight/2 {
				p.State.Position.Y -= MoveSpeed
			}
			p.State.idleYpos = p.State.Position.Y

		case ebiten.KeyDown:
			if input.IsJumpKeyPressed {
				break
			}
			isMoving = true
			if p.State.Position.Y <= input.Cf.ScreenHeight/2 {
				p.State.Position.Y += MoveSpeed
			}
			p.State.idleYpos = p.State.Position.Y

		case ebiten.KeyRight:
			p.State.flipped = false
			isMoving = true
			if p.State.Position.X <= input.Cf.ScreenWidth/2 {
				p.State.Position.X += MoveSpeed
			}
			if (p.State.descending || !input.IsJumpKeyPressed) && p.State.Position.Y < p.State.idleYpos {
				p.State.Position.Y += MoveSpeed
			}

		case ebiten.KeyLeft:
			p.State.flipped = true

			isMoving = true
			if p.State.Position.X >= -input.Cf.ScreenWidth/2 {
				p.State.Position.X -= MoveSpeed
			}
			if (p.State.descending || !input.IsJumpKeyPressed) && p.State.Position.Y < p.State.idleYpos {
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
