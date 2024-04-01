package main

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/game"
	"canon-tower-defense/game/player"
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type EbitenGame struct {
	keys       []ebiten.Key
	stateStack *states.StateStack
}

func NewGame() EbitenGame {
	pl := player.NewPlayer()

	// first stack the level selector
	stateStack := states.NewStateStack()
	levelSelection := states.NewLevelSelection(pl, game.LevelSelector{}, &stateStack)
	stateStack.Push(levelSelection)
	stateStack.Push(states.NewPresentationState())
	return EbitenGame{
		keys:       make([]ebiten.Key, 0),
		stateStack: &stateStack,
	}
}

func (g *EbitenGame) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("exiting")
	}
	return g.stateStack.Update(g.stateStack, g.keys)
}

func (g *EbitenGame) Draw(screen *ebiten.Image) {
	g.stateStack.Draw(screen)
}

func (g *EbitenGame) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Canon defense")
	g := NewGame()
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
