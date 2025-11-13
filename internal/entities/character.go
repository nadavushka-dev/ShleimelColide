package entities

import (
	"shleimel_colide/internal/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X int
	Y int
}

type CharacterState struct {
	Sprite      *ebiten.Image
	CurrentAnim animation.Anim
	Position    *Position
	idleYpos    int
	flipped     bool
	descending  bool
}

type Character struct {
	State CharacterState
}

func NewCharacter(State CharacterState) *Character {
	c := &Character{
		State: State,
	}

	return c
}

func (c *Character) GetBounderies() (xr, xl, yb, yt int) {
	w := int(float64(c.State.CurrentAnim.FrameWidth) * 0.6)
	h := int(float64(c.State.CurrentAnim.FrameHeight) * 0.85)
	xr = c.State.Position.X + (w / 2)
	xl = c.State.Position.X - (w / 2)
	yb = c.State.Position.Y + (h / 2)
	yt = c.State.Position.Y - (h / 2)
	return
}
