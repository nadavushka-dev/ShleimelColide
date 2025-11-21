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
	defer func() {
		if closeErr := logfile.Close(); closeErr != nil {
			log.Printf("Failed to close log fole :%v", closeErr)
		}
	}()
	log.SetOutput(logfile)

	g, err := game.CreateGame(
		config.GetDefaultConfig(),
	)
	if err != nil {
		log.Fatal(err)
	}

	monitorWidth, monitorHeight := ebiten.Monitor().Size()
	windowWidth := int(float64(monitorWidth) * 0.95)
	windowHeight := int(float64(monitorHeight) * 0.95)

	ebiten.SetWindowSize(windowWidth, windowHeight)
	ebiten.SetWindowTitle("My first game")

	// Call ebiten.RunGame to start your game loop.
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
