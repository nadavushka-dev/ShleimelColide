package scenes

import (
	"shleimel_colide/internal/config"
	"shleimel_colide/internal/entities"

	"github.com/hajimehoshi/ebiten/v2"
)

type GameContext interface {
	GetKeysWatched() []ebiten.Key
	GetPressedKeys() []ebiten.Key
	AddPressedKey(key ebiten.Key)
	GetPlayer() *entities.Player
	GetConfig() *config.Config
	GetCount() int
	GetEnemies() []*entities.Enemy
	AddEnemy(enemy *entities.Enemy)
}

type GameSceneConfig interface {
	Update(g GameContext) error
	Draw(g GameContext, screen *ebiten.Image)
}

type GameScene int

const (
	Playgroung = iota
	GameOver
)
