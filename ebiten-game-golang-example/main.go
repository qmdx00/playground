package main

import (
	"fmt"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

var _ ebiten.Game = (*Game)(nil)

const (
	screenWidth  = 1000
	screenHeight = 750
)

type Game struct {
	player *Player

	width, height int
}

func NewGame() *Game {
	return &Game{
		player: NewPlayer(100, 100),
		width:  screenWidth,
		height: screenHeight,
	}
}

// Update proceeds the game state.
// Update is called every tick (1/60 [s] by default).
func (g *Game) Update() (err error) {
	g.player.Update()

	var event PlayerEvent = StopEvent
	for key, _event := range KeyboardEventMap {
		if ebiten.IsKeyPressed(key) {
			event = _event
		}
	}

	g.player.Transition(g, event)
	return
}

// Draw draws the game screen.
// Draw is called every frame (typically 1/60[s] for 60Hz display).
func (g *Game) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	op.GeoM.Translate(g.player.X, g.player.Y)
	// screen.Fill(color.RGBA{0xA0, 0xA0, 0xA0, 255})
	screen.DrawImage(backgroundImage, nil)
	screen.DrawImage(g.player.RenderImage(), op)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("Use arrow keys to move the player.\nPosition: (%.2f, %.2f)\nState: %s", g.player.X, g.player.Y, g.player.fsm.currentState))
}

// Layout takes the outside size (e.g., the window size) and returns the (logical) screen size.
// If you don't have to adjust the screen size with the outside size, just return a fixed size.
func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return g.width, g.height
}

func main() {
	game := NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("My First Ebiten Game")

	if err := ebiten.RunGame(game); err != nil {
		panic(err)
	}
}
