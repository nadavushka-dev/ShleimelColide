package game

import (
	"log"
	"shleimel_colide/internal/config"
	"shleimel_colide/internal/entities"
	"shleimel_colide/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Count  int
	Config config.Config
	Player *entities.Player
	Enemy  *entities.Enemy
}

func CreateGame(cfg config.Config) *Game {
	g := &Game{
		Config: cfg,
	}
	var err error
	g.Player, err = entities.NewPlayer()
	if err != nil {
		log.Fatal("Failed to create player:", err)
	}

	g.Enemy, err = entities.NewEnemy()
	if err != nil {
		log.Fatal("Failed to create enemy:", err)
	}

	return g
}

func (g *Game) Update() error {
	g.Count++
	keys := []ebiten.Key{
		ebiten.KeyUp,
		ebiten.KeyRight,
		ebiten.KeyLeft,
		ebiten.KeyDown,
		ebiten.KeySpace,
	}
	pressedKey := ebiten.Key(0)

	for _, key := range keys {
		if ebiten.IsKeyPressed(key) {
			pressedKey = key
			break
		}
	}

	g.Player.Update(pressedKey)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.Player.Draw(g.Config, screen, g.Count)
	g.Enemy.Draw(g.Config, screen, g.Count)

	utils.LogOnSceen(screen, "Hello Shleimel", nil)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Config.ScreenWidth, g.Config.ScreenHeight
}
