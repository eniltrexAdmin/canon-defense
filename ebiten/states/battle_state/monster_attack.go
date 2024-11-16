package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterAttackState struct {
	ebitenMonsterTeam *ebitenMonsterTeam
}

func newMonsterAttackState(numPositions int, ebitenMonsterTeam *ebitenMonsterTeam) MonsterAttackState {
	ebitenMonsterTeam.monsterAdvancePositions(numPositions)
	return MonsterAttackState{
		ebitenMonsterTeam: ebitenMonsterTeam,
	}
}

func (s MonsterAttackState) Debug() string {
	return "MonsterAttackState State"
}

func (s MonsterAttackState) Update(stack *states.StateStack, keys []ebiten.Key) error {

	s.ebitenMonsterTeam.updateAttack()
	if !s.ebitenMonsterTeam.monsterAttacking {
		stack.Pop()
	}

	return nil
}

func (s MonsterAttackState) Draw(screen *ebiten.Image) {

}
