package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"github.com/hajimehoshi/ebiten/v2"
)

type FireCannonsState struct {
	ebitenCanonDeck    *ebitenCanonDeck
	ebitenBattleGround *ebitenBattleGround
}

func (s FireCannonsState) Debug() string {
	return "BattleState State"
}

func (s FireCannonsState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	s.ebitenBattleGround.update(s.ebitenCanonDeck.currentBullets())
	s.ebitenCanonDeck.firingUpdate()

	if !s.ebitenCanonDeck.Firing {
		monsterAdvanceState := newMonsterAttackState(1, s.ebitenBattleGround)
		stack.Switch(monsterAdvanceState)
	}

	return nil
}

func (s FireCannonsState) Draw(screen *ebiten.Image) {

}