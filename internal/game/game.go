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
	Enemy  []*entities.Enemy
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

	for range 5 {
		e, err := entities.NewEnemy(g.Config)
		if err != nil {
			log.Fatal("Failed to create enemy:", err)
		}

		g.Enemy = append(g.Enemy, e)
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
	for _, e := range g.Enemy {
		e.Draw(g.Config, screen, g.Count)
	}
	t := "No Collision"

	for _, e := range g.Enemy {
		if collisionDetection(g.Player, e) {
			t = "Collision!!!"
		}
	}

	utils.LogOnSceen(screen, t, nil)
}

func collisionDetection(ent1 *entities.Player, ent2 *entities.Enemy) bool {
	ent1xr, ent1xl, ent1yb, ent1yt := ent1.GetBounderies()
	ent2xr, ent2xl, ent2yb, ent2yt := ent2.GetBounderies()

	if ent1xr > ent2xl && ent1xl < ent2xr && ent1yt < ent2yb && ent1yb > ent2yt {
		return true
	}

	return false
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.Config.ScreenWidth, g.Config.ScreenHeight
}
