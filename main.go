package main

import (
	_ "image/png"
	"log"
	"os"
	"shleimel_colide/internal/config"
	"shleimel_colide/internal/game"

	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	logfile, err := os.OpenFile("debug.log", os.O_RDWR|os.O_CREATE|os.O_SYNC|os.O_APPEND, 0644)
	if err != nil {
		log.Fatal(err)
	}
	log.SetOutput(logfile)
	defer logfile.Close()

	g := game.CreateGame(
		config.GetDefaultConfig(),
	)

	// Specify the window size as you like. Here, a doubled size is specified.
	ebiten.SetWindowSize(g.Config.ScreenWidth*2, g.Config.ScreenHeight*2)
	ebiten.SetWindowTitle("My first game")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
