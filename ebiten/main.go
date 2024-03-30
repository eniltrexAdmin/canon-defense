package main

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/game"
	"canon-tower-defense/game/player"
	"canon-tower-defense/pkg"
	"errors"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"log"
)

type EbitenGameInterface struct {
	keys       []ebiten.Key
	stateStack pkg.StateStack
}

func NewGame() EbitenGameInterface {
	pl := player.NewPlayer()
	levelSelection := states.NewLevelSelection(pl, game.LevelSelector{})
	// first stack the level selector
	stateStack := pkg.NewStateStack(levelSelection)
	stateStack.Push(states.NewPresentationState())
	return EbitenGameInterface{
		keys:       make([]ebiten.Key, 0),
		stateStack: stateStack,
	}
}

func (g *EbitenGameInterface) Update() error {
	g.keys = inpututil.AppendPressedKeys(g.keys[:0])
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		return errors.New("exiting")
	}
	return g.stateStack.Update(&g.stateStack, g.keys)
}

func (g *EbitenGameInterface) Draw(screen *ebiten.Image) {
	g.stateStack.Draw(screen)
}

func (g *EbitenGameInterface) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return 640, 480
}

func main() {
	ebiten.SetWindowSize(640, 480)
	ebiten.SetVsyncEnabled(true)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Canon defense")
	game := NewGame()
	if err := ebiten.RunGame(&game); err != nil {
		log.Fatal(err)
	}
}
