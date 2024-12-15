package game_over_state

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_monster"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text"
	"golang.org/x/image/font/basicfont"
)

type GameOverState struct {
	ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam
}

func NewGameOverState(ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam) *GameOverState {
	return &GameOverState{
		ebitenMonsterTeam: ebitenMonsterTeam,
	}
}

func (s *GameOverState) Debug() string {
	return "Game Over State"
}

func (s *GameOverState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	s.ebitenMonsterTeam.Update()

	var justPressedTouchIDs []ebiten.TouchID
	justPressedTouchIDs = inpututil.AppendJustPressedTouchIDs(justPressedTouchIDs)
	touchJustPressed := len(justPressedTouchIDs) > 0

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) || touchJustPressed {
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
