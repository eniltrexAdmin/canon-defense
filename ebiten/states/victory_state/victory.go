package victory_state

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

type VictoryState struct {
	ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam
}

func NewVictoryState(ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam, completedLevel int) *VictoryState {
	constants.GlobalContext.Session.CompleteLevel(completedLevel - 1)
	return &VictoryState{
		ebitenMonsterTeam: ebitenMonsterTeam,
	}
}

func (s *VictoryState) Debug() string {
	return "VictoryState State"
}

func (s *VictoryState) Update(stack *states.StateStack, keys []ebiten.Key) error {
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

func (s *VictoryState) Draw(screen *ebiten.Image) {
	basicFont := basicfont.Face7x13

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(4, 4)
	op.GeoM.Translate(constants.ScreenWidth/2-100, constants.ScreenHeight/2-10)

	text.DrawWithOptions(screen,
		"VICTORY!!!",
		basicFont,
		op,
	)

	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
