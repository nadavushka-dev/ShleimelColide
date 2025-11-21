package scenes

import (
	"fmt"
	"shleimel_colide/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameOverScene struct {
}

func (sc *GameOverScene) Update(g GameContext) error {
	return nil
}

func (sc *GameOverScene) Draw(g GameContext, screen *ebiten.Image) {
	utils.LogOnSceen(screen, fmt.Sprintf("Game Over! Your score: %d", g.GetScore()), nil, nil, nil, true, true)
}
