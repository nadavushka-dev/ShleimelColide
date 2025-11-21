package entities

import (
	"shleimel_colide/internal/config"

	"github.com/hajimehoshi/ebiten/v2"
)

type Collidable interface {
	GetBoundaries() (xr, xl, yb, yt int)
}

type Drawable interface {
	Draw(cfg config.Config, screen *ebiten.Image, tick int)
}
