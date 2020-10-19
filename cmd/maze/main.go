package main

import (
	"os"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/necrophonic/gopher-maze/internal/game"
)

func main() {
	run()
}

func run() {
	g := game.New()

	if err := g.Run(); err != nil {
		debug.Println("Error running game:", err)
		os.Exit(1)
	}

	// fmt.Println(game.Swatch())
}

// ╔════════════════════════╗
// ║ ▓▓                  ▓▓ ║
// ║ ▓▓▓▓              ▓▓▓▓ ║
// ║ ▓▓▓▓▓▓            ▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓▓      ░░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓▓▓▓  ▓▓░░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓▓▓      ░░░░▓▓▓▓ ║
// ║ ▓▓▓▓▓▓            ▓▓▓▓ ║
// ║ ▓▓▓▓              ▓▓▓▓ ║
// ║ ▓▓                  ▓▓ ║
// ╚════════════════════════╝
