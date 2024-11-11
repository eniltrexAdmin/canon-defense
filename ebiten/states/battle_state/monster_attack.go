package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterAttackState struct {
	ebitenBattleGround *ebitenBattleGround
}

func newMonsterAttackState(numPositions int, ebitenBattleGround *ebitenBattleGround) MonsterAttackState {
	ebitenBattleGround.monsterAdvancePositions(numPositions)
	return MonsterAttackState{
		ebitenBattleGround: ebitenBattleGround,
	}
}

func (s MonsterAttackState) Debug() string {
	return "MonsterAttackState State"
}

func (s MonsterAttackState) Update(stack *states.StateStack, keys []ebiten.Key) error {

	s.ebitenBattleGround.updateAttack()
	if !s.ebitenBattleGround.monsterAttacking {
		stack.Pop()
	}

	return nil
}

func (s MonsterAttackState) Draw(screen *ebiten.Image) {

}
