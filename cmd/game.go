package main

import (
	"canon-tower-defense/pkg"
	"canon-tower-defense/states"
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type Game struct {
	keys       []ebiten.Key
	stateStack pkg.StateStack
}

func NewGame() Game {
	return Game{
		keys:       make([]ebiten.Key, 0),
		stateStack: pkg.NewStateStack(states.NewPresentationState()),
	}
}

func (g *Game) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("exiting")
	}
	return g.stateStack.Update()
}

func (g *Game) Draw(screen *ebiten.Image) {
	g.stateStack.Draw(screen)
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 320, 240
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetWindowTitle("Canon defense")
	game := NewGame()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
