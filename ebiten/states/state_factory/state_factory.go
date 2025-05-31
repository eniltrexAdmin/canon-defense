package state_factory

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state"
	"canon-tower-defense/ebiten/states/level_selection"
	"fmt"
)

type CanonTowerDefenseStaticStateFactory struct{}

func (sgs CanonTowerDefenseStaticStateFactory) Create(stateName string, params ...any) states.State {
	switch stateName {
	case states.BattleStateName:
		return sgs.battleState(getLevel(params...))
	case states.LevelSelectionStateName:
		return sgs.levelSelection()
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

func (sgs CanonTowerDefenseStaticStateFactory) battleState(level int) states.State {
	return battle_state.NewBattleState(level)
}

func (sgs CanonTowerDefenseStaticStateFactory) levelSelection() states.State {
	return level_selection.NewLevelSelection()
}
