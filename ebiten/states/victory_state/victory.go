package victory_state

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type VictoryState struct{}

func (s *VictoryState) Debug() string {
	return "VictoryState State"
}

func (s *VictoryState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		// TODO player should be probably global for the whole game.
		st := stack.StateFactory.Create(states.LevelSelectionStateName, stack)
		stack.Pop()
		stack.Pop()
		stack.Switch(st)
	}
	return nil
}

func (s *VictoryState) Draw(screen *ebiten.Image) {
	basicFont := basicfont.Face7x13

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(constants.ScreenWidth/2, constants.ScreenHeight/2)

	text.DrawWithOptions(screen,
		"VICTORY!!!",
		basicFont,
		op,
	)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
