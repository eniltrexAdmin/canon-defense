package state_factory

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state"
	"canon-tower-defense/ebiten/states/level_selection_state"
	"canon-tower-defense/game"
	"fmt"
)

type CanonTowerDefenseStaticStateFactory struct{}

func (sgs CanonTowerDefenseStaticStateFactory) Create(stateName string, params ...any) states.State {
	switch stateName {
	case states.BattleStateName:
		return sgs.battleState(getLevel(params...))
	case states.LevelSelectionStateName:
		st := getPointerStack(params...)
		return sgs.levelSelection(st)
	default:
		panic(fmt.Sprintf("Unknown state name: %s", stateName))
	}
}

func getLevel(params ...any) int {
	if len(params) < 1 {
		panic(fmt.Errorf("missing parameter: level"))
	}
	level, ok := params[0].(int)
	if !ok {
		panic(fmt.Errorf("invalid parameter type: level"))
	}
	return level
}

func getPointerStack(params ...any) *states.StateStack {
	if len(params) < 1 {
		panic(fmt.Errorf("missing parameters: player and stack"))
	}
	stack, ok := params[0].(*states.StateStack)
	if !ok {
		panic(fmt.Errorf("invalid parameter type: stack, real type"))
	}
	return stack
}

func (sgs CanonTowerDefenseStaticStateFactory) battleState(level int) states.State {
	return battle_state.NewBattleState(level)
}

func (sgs CanonTowerDefenseStaticStateFactory) levelSelection(stack *states.StateStack) states.State {
	return level_selection_state.NewLevelSelection(game.LevelSelector{}, stack)
}
