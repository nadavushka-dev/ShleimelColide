# Shliemel Colide

A (very) simple 2D game built with Go and Ebiten.

---

*Game is still under development*

---

## About

All Shliemel wanted was to buy milk for his ol' bird at home. Simple task, right? WRONG. 

Somehow, somewhere between the couch and the grocery store, things went absolutely sideways. Now he's stuck in the most bizarre
situation imaginable - monsters are spawning left and right, and they just keep coming! No one knows why. No one knows how. But
it is not Shliemel's fault!! Now he needs to dodge, weave, and survive this absolute nonsense while clutching his shopping list.

Will he get the milk? Will he make it home? Will his woman ever forgive him for being late? Find out in this bizzare game where a
simple errand turns into the adventure **nobody asked for!**

## Features

- **Player Movement**: Arrow keys to move up, down, left, and right
<!-- - **Jump Mechanics**: Space bar to jump with smooth ascending/descending motion -->
- **Sprite Animations**: 
  - Idle animation (4 frames)
  - Running animation (8 frames)
  <!-- - Jump animation -->
- **Sprite Flipping**: Character faces the direction of movement

## Tech Stack

- **Language**: Go 1.25.3
- **Game Engine**: [Ebiten v2](https://ebiten.org/) - A simple 2D game engine for Go

## Controls

- `Arrow Keys`: Move the character
<!-- - `Space`: Jump -->
- `Up/Down`: Move vertically
- `Left/Right`: Move horizontally

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

## Project Structure

```
.
├── main.go                    # Entry point
├── internal/
│   ├── animation/            # Animation system
│   ├── config/               # Game configuration
│   ├── entities/             # Game entities (Player, Character)
│   └── game/                 # Core game loop
└── go.mod                    # Dependencies
```

## Future Ideas

- Add collision detection
- Implement enemies
- Create levels
- Add sound effects
- Scoring system
