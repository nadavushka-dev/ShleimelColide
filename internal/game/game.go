package game

import (
	"fmt"
	"shleimel_colide/internal/config"
	"shleimel_colide/internal/entities"
	"shleimel_colide/internal/scenes"

	"github.com/hajimehoshi/ebiten/v2"
)

type Game struct {
	Count       int
	Config      config.Config
	Player      *entities.Player
	Enemy       []*entities.Enemy
	PressedKeys []ebiten.Key
	keysWatched []ebiten.Key
	Scene       scenes.GameScene
}

func CreateGame(cfg config.Config) (*Game, error) {
	g := &Game{
		Config:      cfg,
		PressedKeys: make([]ebiten.Key, 0, 5),
		keysWatched: []ebiten.Key{
			ebiten.KeyUp,
			ebiten.KeyRight,
			ebiten.KeyLeft,
			ebiten.KeyDown,
			ebiten.KeySpace,
			ebiten.KeyH,
			ebiten.KeyJ,
			ebiten.KeyK,
			ebiten.KeyL,
		},
	}

	var err error
	g.Player, err = entities.NewPlayer()
	if err != nil {
		return nil, fmt.Errorf("Failed to create player: %w", err)
	}

	// for range 5 {
	// 	e, err := entities.NewEnemy(g.Config)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("Failed to create enemy: %w", err)
	// 	}
	//
	// 	g.Enemy = append(g.Enemy, e)
	// }

	return g, nil
}

func (g *Game) Update() error {
	g.Count++
	g.PressedKeys = g.PressedKeys[:0]
	var sc scenes.GameSceneConfig
	switch g.Scene {
	case scenes.Playgroung:
		sc = &scenes.PlaygroundScene{}
		err := sc.Update(g)
		if err != nil {
			return fmt.Errorf("Failed to update playground scene: %w", err)
		}
	}

	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	var sc scenes.GameSceneConfig
	switch g.Scene {
	case scenes.Playgroung:
		sc = &scenes.PlaygroundScene{}
		sc.Draw(g, screen)
	}

}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Config.ScreenWidth, g.Config.ScreenHeight
}

func (g *Game) GetKeysWatched() []ebiten.Key {
	return g.keysWatched
}

func (g *Game) GetPressedKeys() []ebiten.Key {
	return g.PressedKeys
}

func (g *Game) AddPressedKey(key ebiten.Key) {
	g.PressedKeys = append(g.PressedKeys, key)
}

func (g *Game) GetPlayer() *entities.Player {
	return g.Player
}

func (g *Game) GetConfig() *config.Config {
	return &g.Config
}

func (g *Game) GetCount() int {
	return g.Count
}

func (g *Game) GetEnemies() []*entities.Enemy {
	return g.Enemy
}

func (g *Game) AddEnemy(enemy *entities.Enemy) {
	g.Enemy = append(g.Enemy, enemy)
}
