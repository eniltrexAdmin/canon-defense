package game_over_state

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

type GameOverState struct{}

func (s *GameOverState) Debug() string {
	return "Game Over State"
}

func (s *GameOverState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		st := stack.StateFactory.Create(states.LevelSelectionStateName, stack)
		stack.Pop()
		stack.Pop()
		stack.Switch(st)
	}
	return nil
}

func (s *GameOverState) Draw(screen *ebiten.Image) {
	basicFont := basicfont.Face7x13

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(constants.ScreenWidth/2-120, constants.ScreenHeight/2-100)

	text.DrawWithOptions(screen,
		"Game Over!",
		basicFont,
		op,
	)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
