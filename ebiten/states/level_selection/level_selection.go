package level_selection

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type LevelSelection struct {
	levels   []*Level
	touchIDs []ebiten.TouchID
}

// TODO that parameter should be a call, or something more related to the generator, probably inverse DI

func NewLevelSelection(LevelSelector game.LevelSelector) LevelSelection {
	pl := constants.GlobalContext.Player
	state := LevelSelection{
		levels: NewLevelSet(LevelSelector.LevelSelection(pl)),
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

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		for _, level := range p.levels {
			if level.InBounds(x, y) {
				GoBattle(stack, level.LevelNumber)
			}
		}
	}
	p.touchIDs = inpututil.AppendJustPressedTouchIDs(p.touchIDs[:0])
	for _, id := range p.touchIDs {
		x, y = ebiten.TouchPosition(id)
		for _, level := range p.levels {
			if level.InBounds(x, y) {
				GoBattle(stack, level.LevelNumber)
			}
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
