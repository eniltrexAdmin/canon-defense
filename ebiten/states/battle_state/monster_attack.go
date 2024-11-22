package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_monster"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterAttackState struct {
	ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam
}

func newMonsterAttackState(ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam, g *game.CanonTDGame) MonsterAttackState {
	ebitenMonsterTeam.Advance()
	g.MonstersCharge()
	return MonsterAttackState{
		ebitenMonsterTeam: ebitenMonsterTeam,
	}
}

func (s MonsterAttackState) Debug() string {
	return "MonsterAttackState State"
}

func (s MonsterAttackState) Update(stack *states.StateStack, keys []ebiten.Key) error {

	s.ebitenMonsterTeam.Update()
	if !s.ebitenMonsterTeam.AreAttacking() {
		stack.Pop()
	}

	return nil
}

func (s MonsterAttackState) Draw(screen *ebiten.Image) {

}
