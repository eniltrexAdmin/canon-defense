package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_monster"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterAttackState struct {
	ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam
}

func newMonsterAttackState(numPositions int, ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam) MonsterAttackState {
	//ebitenMonsterTeam.monsterAdvancePositions(numPositions)
	return MonsterAttackState{
		ebitenMonsterTeam: ebitenMonsterTeam,
	}
}

func (s MonsterAttackState) Debug() string {
	return "MonsterAttackState State"
}

func (s MonsterAttackState) Update(stack *states.StateStack, keys []ebiten.Key) error {

	s.ebitenMonsterTeam.UpdateAttack()
	//if !s.ebitenMonsterTeam.monsterAttacking {
	//	stack.Pop()
	//}

	return nil
}

func (s MonsterAttackState) Draw(screen *ebiten.Image) {

}
