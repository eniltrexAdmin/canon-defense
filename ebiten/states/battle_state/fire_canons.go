package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_monster"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type FireCannonsState struct {
	ebitenCanonDeck   *ebiten_canon.EbitenCanonDeck
	ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam
	game              *game.CanonTDGame
}

func (s FireCannonsState) Debug() string {
	return "BattleState State"
}

func (s FireCannonsState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	s.ebitenMonsterTeam.UpdateDeckFiring(s.ebitenCanonDeck.CurrentBullets())
	s.ebitenCanonDeck.FiringUpdate()

	//if !s.ebitenCanonDeck.Firing {
	//	monsterAdvanceState := newMonsterAttackState(1, s.ebitenMonsterTeam)
	//	stack.Switch(monsterAdvanceState)
	//}

	return nil
}

func (s FireCannonsState) Draw(screen *ebiten.Image) {

}
