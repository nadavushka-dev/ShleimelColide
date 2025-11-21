package scenes

import (
	"fmt"
	"shleimel_colide/internal/entities"
	"shleimel_colide/internal/utils"

	"github.com/hajimehoshi/ebiten/v2"
)

type PlaygroundScene struct {
}

func (sc *PlaygroundScene) Update(g GameContext) error {
	keysWatched := g.GetKeysWatched()
	player := g.GetPlayer()
	config := g.GetConfig()
	count := g.GetCount()
	enemies := g.GetEnemies()

	isJumpKeyPressed := false

	for _, key := range keysWatched {
		if ebiten.IsKeyPressed(key) {
			if key == ebiten.KeySpace {
				isJumpKeyPressed = true
			}

			g.AddPressedKey(key)
		}
	}

	player.Update(entities.UpdateInput{Keys: g.GetPressedKeys(), IsJumpKeyPressed: isJumpKeyPressed, Cf: *config})

	if count%50 == 0 {
		e, err := entities.NewEnemy(*config)
		if err != nil {
			return fmt.Errorf("Failed to create enemy: %w", err)
		}
		g.AddEnemy(e)
		g.IncrementScore()
	}

	for _, e := range enemies {
		e.Update(*player.State.Position)
	}

	return nil
}

func (sc *PlaygroundScene) Draw(g GameContext, screen *ebiten.Image) {
	config := g.GetConfig()
	count := g.GetCount()
	enemies := g.GetEnemies()
	player := g.GetPlayer()

	for _, e := range enemies {
		e.Draw(*config, screen, count)
	}
	t := fmt.Sprintf("score: %d", g.GetScore())

	for _, e := range enemies {
		if utils.CollisionDetection(player, e) {
			g.SetScene(GameOver)
		}
	}
	player.Draw(*config, screen, count)

	utils.LogOnSceen(screen, t, nil, nil, nil, false, false)
}
