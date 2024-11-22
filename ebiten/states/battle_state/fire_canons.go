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

func NewFireCannonState(
	ebitenCanonDeck *ebiten_canon.EbitenCanonDeck,
	ebitenMonsterTeam *ebiten_monster.EbitenMonsterTeam,
	game *game.CanonTDGame,
) *FireCannonsState {
	ebitenMonsterTeam.DeckFiring(ebitenCanonDeck.CurrentBullets())
	fcs := FireCannonsState{
		ebitenCanonDeck:   ebitenCanonDeck,
		ebitenMonsterTeam: ebitenMonsterTeam,
		game:              game,
	}
	return &fcs
}

func (s *FireCannonsState) Debug() string {
	return "FireCannons State"
}

func (s *FireCannonsState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	s.ebitenMonsterTeam.Update()
	s.ebitenCanonDeck.Update()

	bullets := s.ebitenCanonDeck.CurrentBullets()
	if len(bullets) == 0 {
		s.ebitenCanonDeck.Firing = false
		monsterAdvanceState := newMonsterAttackState(s.ebitenMonsterTeam, s.game)
		stack.Switch(monsterAdvanceState)
	}

	return nil
}

func (s *FireCannonsState) Draw(screen *ebiten.Image) {

}
