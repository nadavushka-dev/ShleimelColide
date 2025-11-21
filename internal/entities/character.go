package entities

import (
	"bytes"
	"image"
	"shleimel_colide/internal/animation"
	"shleimel_colide/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/examples/resources/images"
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

func (c *Character) GetBoundaries() (xr, xl, yb, yt int) {
	w := int(float64(c.State.CurrentAnim.FrameWidth) * 0.6)
	h := int(float64(c.State.CurrentAnim.FrameHeight) * 0.85)
	xr = c.State.Position.X + (w / 2)
	xl = c.State.Position.X - (w / 2)
	yb = c.State.Position.Y + (h / 2)
	yt = c.State.Position.Y - (h / 2)
	return
}

func (c *Character) Draw(cfg config.Config, screen *ebiten.Image, tick int) {
	op := &ebiten.DrawImageOptions{}

	if c.State.flipped {
		op.GeoM.Scale(-1, 1)
		op.GeoM.Translate(float64(c.State.CurrentAnim.FrameWidth), 0)
	}

	op.GeoM.Translate(
		-float64(c.State.CurrentAnim.FrameWidth)/2, -float64(c.State.CurrentAnim.FrameHeight)/2,
	)

	op.GeoM.Translate(
		float64(cfg.ScreenWidth/2+c.State.Position.X),
		float64(cfg.ScreenHeight/2+c.State.Position.Y),
	)

	i := (tick / cfg.TicksPerFrame) % c.State.CurrentAnim.FrameCount
	sx := cfg.FrameOX + i*c.State.CurrentAnim.FrameWidth
	sy := c.State.CurrentAnim.Row * c.State.CurrentAnim.FrameHeight
	screen.DrawImage(c.State.Sprite.SubImage(
		image.Rect(sx, sy, sx+c.State.CurrentAnim.FrameWidth, sy+c.State.CurrentAnim.FrameHeight),
	).(*ebiten.Image), op)
}

func loadSprite() (image.Image, error) {
	img, _, err := image.Decode(bytes.NewReader(images.Runner_png))
	if err != nil {
		return nil, err
	}
	return img, nil
}
