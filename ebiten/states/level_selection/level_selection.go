package level_selection

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LevelSelection struct {
	levels game.Levels
}

func NewLevelSelection(LevelSelector game.LevelSelector) LevelSelection {
	pl := constants.GlobalContext.Player
	state := LevelSelection{
		levels: LevelSelector.LevelSelection(pl),
	}

	return state
}

func (p LevelSelection) Debug() string {
	return "LevelSelection State"
}

func (p LevelSelection) Update(stack *states.StateStack, keys []ebiten.Key) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		bs := stack.StateFactory.Create(states.BattleStateName, 1)
		stack.Switch(bs)
	}
	return nil
}

func (p LevelSelection) Draw(screen *ebiten.Image) {

}
