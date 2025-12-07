# Shleimel Colide

A (very) simple 2D game built with Go and Ebiten.

---

*Game is still under development*

---

## About

All Shleimel wanted was to buy milk for his ol' bird at home. Simple task, right? WRONG. 

Somehow, somewhere between the couch and the grocery store, things went absolutely sideways. Now he's stuck in the most bizarre
situation imaginable - monsters are spawning left and right, and they just keep coming! No one knows why. No one knows how. But
it is not Shleimel's fault!! Now he needs to dodge, weave, and survive this absolute nonsense while clutching his shopping list.

Will he get the milk? Will he make it home? Will his woman ever forgive him for being late? Find out in this bizarre game where a
simple errand turns into the adventure **nobody asked for!**

## Features

- **Player Movement**: Smooth 8-directional movement
- **Jump Mechanics**: Space bar to jump with gravity-based physics
- **Sprite Animations**: 
  - Idle animation (4 frames)
  - Running animation (8 frames)
  - Jump animation (4 frames)
- **Sprite Flipping**: Character faces the direction of movement
- **Enemy System**: Monsters spawn dynamically and chase the player
- **Collision Detection**: AABB-based collision system between player and enemies
- **Scene System**: Modular scene management (currently: Playground scene)
- **Debug Logging**: On-screen collision status and file-based debug logs

## Tech Stack

- **Language**: Go 1.25.3
- **Game Engine**: [Ebiten v2](https://ebiten.org/) - A simple 2D game engine for Go

## Controls

- `Arrow Keys`: Move in all directions (Up/Down/Left/Right)
- `Space`: Jump

## Getting Started

### Prerequisites

- Go 1.25.3 or higher

### Installation

1. Clone the repository
2. Install dependencies:
   ```bash
   go mod download
   ```

### Running the Game

```bash
go run main.go
```

The game window will automatically size to 95% of your monitor dimensions.

## Project Structure

```
.
├── main.go                    # Entry point
├── internal/
│   ├── animation/            # Animation system
│   ├── config/               # Game configuration (screen size, frame timing)
│   ├── entities/             # Game entities (Player, Enemy, Character base)
│   ├── game/                 # Core game loop and state management
│   ├── scenes/               # Scene system (Playground scene)
│   └── utils/                # Utilities (collision detection, on-screen text)
├── go.mod                    # Dependencies
└── debug.log                 # Runtime debug logs
```

## Game Mechanics

### Player
- Moves at 3 pixels per frame
- Can jump up to 30 pixels high
- Uses gravity-based physics for realistic jumping
- Collision hitbox is approximately 60% width × 85% height of sprite

### Enemies
- Spawn every 50 game ticks at random positions
- Chase the player by moving 1 pixel per frame toward player position
- Use the same sprite sheet as the player
- Detect collisions with the player

### Technical Details
- Default screen resolution: 640×480 (scales to 95% of monitor)
- Animation timing: 10 ticks per frame
- Sprite dimensions: 32×32 pixels per frame

## Future Ideas

- Game over/restart system when collision occurs
- Health system and damage mechanics
- Multiple enemy types with different behaviors
- Power-ups and collectibles
- Level progression system
- Sound effects and background music
- Scoring system based on survival time
- Particle effects for collisions
- Menu system (start screen, pause, game over)
