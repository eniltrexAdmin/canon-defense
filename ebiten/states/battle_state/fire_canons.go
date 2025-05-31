package battle_state

import (
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/ebiten_sound"
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
	ebitenCanonDeck.FireCanons(game)
	ebitenMonsterTeam.DeckFiring(ebitenCanonDeck.CurrentBullets())
	fireSound := ebiten_sound.MustNewPlayer(assets.FireCanon)
	fireSound.Play()
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
		monsterAdvanceState := newMonsterAttackState(s.ebitenMonsterTeam, s.game)
		stack.Switch(monsterAdvanceState)
	}

	return nil
}

func (s *FireCannonsState) Draw(screen *ebiten.Image) {

}
