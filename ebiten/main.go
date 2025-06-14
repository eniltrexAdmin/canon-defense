package main

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/presentation_state"
	"canon-tower-defense/ebiten/states/state_factory"
	"canon-tower-defense/game"
	"canon-tower-defense/game/session"
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
	constants.GlobalContext.Session = session.NewSession(game.TotalLevels)

	factory := state_factory.CanonTowerDefenseStaticStateFactory{}

	stateStack := states.NewStateStack(factory)
	levelSelection := factory.Create(states.LevelSelectionStateName, &stateStack)
	stateStack.Push(levelSelection)
	stateStack.Push(presentation_state.NewPresentationState())

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
	return g.stateStack.Update(g.keys)
}

func (g *EbitenGame) Draw(screen *ebiten.Image) {
	g.stateStack.Draw(screen)
}

func (g *EbitenGame) Layout(outsideWidth, outsideHeight int) (int, int) {
	return outsideWidth, outsideHeight // the contents inside the window, if it doesn't match the "set window size"
	// if its bigger, the window will actually crop, since the ocntents are bigger
	// if its smaller, there will be some black padding.
}

func main() {
	ebiten.SetWindowSize(constants.ScreenWidth, constants.ScreenHeight) // the window that opens when you execute
	ebiten.SetVsyncEnabled(true)
	//ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	ebiten.SetWindowTitle("Canon defense")
	g := NewGame()
	if err := ebiten.RunGame(&g); err != nil {
		log.Fatal(err)
	}
}
