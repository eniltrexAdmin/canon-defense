package level_selection

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LevelSelection struct {
	levels   []*Level
	touchIDs []ebiten.TouchID
}

func NewLevelSelection() LevelSelection {
	pl := constants.GlobalContext.Session
	state := LevelSelection{
		levels: NewLevelSet(pl.CompletedLevels),
	}

	return state
}

func (p LevelSelection) Debug() string {
	return "LevelSelection State"
}

func (p LevelSelection) Update(stack *states.StateStack, keys []ebiten.Key) error {
	x, y := ebiten.CursorPosition()
	for _, level := range p.levels {
		level.Update(x, y)
	}

	var selectedLevel *Level

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		selectedLevel = p.GetSelectedLevel(x, y)
	}
	p.touchIDs = inpututil.AppendJustPressedTouchIDs(p.touchIDs[:0])
	for _, id := range p.touchIDs {
		x, y = ebiten.TouchPosition(id)
		selectedLevel = p.GetSelectedLevel(ebiten.TouchPosition(id))
	}
	if selectedLevel != nil {
		GoBattle(stack, selectedLevel.LevelNumber)
	}

	return nil
}

func (p LevelSelection) GetSelectedLevel(xInt, yInt int) *Level {
	for _, level := range p.levels {
		if level.enabled && level.InBounds(xInt, yInt) {
			return level
		}
	}
	return nil
}

func GoBattle(stack *states.StateStack, level int) {
	bs := stack.StateFactory.Create(states.BattleStateName, level)
	stack.Switch(bs)
}

func (p LevelSelection) Draw(screen *ebiten.Image) {
	for _, level := range p.levels {
		level.Draw(screen)
	}
}
