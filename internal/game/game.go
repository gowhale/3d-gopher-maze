package game

import (
	"bufio"
	"fmt"
	"os"

	"github.com/necrophonic/gopher-maze/internal/debug"
	"github.com/necrophonic/gopher-maze/internal/game/element"
	"github.com/necrophonic/gopher-maze/internal/game/terminal"
	"github.com/pkg/errors"
)

type (
	moveVector struct {
		x int8
		y int8
	}

	// Game represents the current game state
	Game struct {
		player *Player
		m      *Maze
		v      *view
		move   moveVector
		gopher *gopher
		items  []item
		state  State

		Msg string
	}
)

// Constants for types of maze space
const (
	SpaceEmpty       spaceType = ' '
	SpaceWall        spaceType = 'X'
	SpacePlayerStart spaceType = 'p'
	SpaceGopherStart spaceType = 'g'
)

// State represents the current game state
type State uint8

// Game states
const (
	sWin State = iota
	sRunning
	sReady
)

type (
	spaceType uint8

	space struct {
		t spaceType
	}
)

type item interface {
	GetPoint() Point
	GetMatrix(distance int) (element.PixelMatrix, error)
}

// New creates a new game state
func New() *Game {
	return &Game{
		player: &Player{
			o: 'n',
		},
		m: &Maze{
			grid:   grid{},
			panels: nil,
		},
		v: &view{
			screen:  make([]element.PixelMatrix, numPanels),
			overlay: element.PixelMatrix{},
		},
		move:   moveVector{0, -1},
		gopher: &gopher{},
		items:  []item{},
		state:  sReady,
		Msg:    "",
	}
}

// Swatch returns a string with the defined game elements
func Swatch() string {
	return fmt.Sprintf(
		"%-10s: %c\n%-10s: %c\n",
		"Wall #1", walls[0],
		"Wall #2", walls[1],
	)
}

// Run performs the main game loop
func (g *Game) Run() error {
	g.state = sRunning

	// TODO randomly (totally or from set of criteria) select a maze
	// TODO would be nice to able to dynamically create one!
	if err := g.importMaze(mazes[2]); err != nil {
		return errors.WithMessage(err, "failed to import maze")
	}

	for {
		if !debug.Debug {
			terminal.Clear()
		}

		if g.state == sWin {
			g.Msg = "You won!"
		}

		if err := g.updateView(); err != nil {
			return errors.WithMessage(err, "failed to update view")
		}

		viewport, err := g.render()
		if err != nil {
			return errors.WithMessage(err, "failed to render scene")
		}
		fmt.Print(viewport)

		if g.state == sWin {
			debug.Println("Game was won")
			return nil
		}

		reader := bufio.NewReader(os.Stdin)

		char, _, err := reader.ReadRune()
		if err != nil {
			return errors.WithMessage(err, "error reading rune from terminal")
		}

		// TODO remove need to hit return!

		switch char {
		case 'w':
			debug.Println("Move forward")
			g.moveForward()
		case 's':
			debug.Println("Move backward")
			g.moveBackwards()
		case 'd':
			debug.Println("Turn right")
			g.rotateRight()
		case 'a':
			debug.Println("Turn left")
			g.rotateLeft()
		case 'h':
			g.Msg = "Move with w,a,s,d (w=forward, a=turn left, d=turn right, s=backwards)"
		case 'q':
			debug.Println("Exiting game")
			fmt.Println("Goodbye!")
			return nil
		default:
			g.Msg = "Sorry, I didn't understand that one! (use 'h' for help)"
		}
		debug.Printf("Player is now at (%v). Facing (%c)\n", g.player.p, g.player.o)
	}
}
