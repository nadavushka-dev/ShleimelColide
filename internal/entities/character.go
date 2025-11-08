package entities

import (
	"game/internal/animation"

	"github.com/hajimehoshi/ebiten/v2"
)

type Position struct {
	X int
	Y int
}

type CharacterState struct {
	Sprite      *ebiten.Image
	CurrentAnim animation.Anim
	position    *Position
	actualYPos  int
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
