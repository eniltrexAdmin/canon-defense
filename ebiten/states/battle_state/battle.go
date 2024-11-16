package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

type BattleState struct {
	game *game.CanonTDGame
	// IMPORTANT if I am modifying this objects inside like canon deck and battle ground through the
	// battle state UPDATE function, update must have a POINTER receiver!!!
	// (or else it's just a copy edited and thrown away).
	// or whatever gets modified needs to be referenced.
	ebitenCanonDeck    *ebitenCanonDeck
	ebitenBattleGround ebitenBattleGround
	ebitenMonsterTeam  *ebitenMonsterTeam
}

func NewBattleState(level int) BattleState {
	// loading assets, could be in init() and consistent usage of states.
	g := game.Start(level)

	ecd := newEbitenCanonDeck(g)
	ebg := newEbitenBattleGround(g.Battleground)
	emt := NewEbitenMonsterTeam(g.MonsterTeam)

	return BattleState{
		game:               &g,
		ebitenCanonDeck:    &ecd,
		ebitenBattleGround: ebg,
		ebitenMonsterTeam:  &emt,
	}
}

func (s BattleState) Debug() string {
	return "BattleState State"
}

func (s BattleState) Update(stack *states.StateStack, keys []ebiten.Key) error {

	s.ebitenMonsterTeam.update(s.ebitenCanonDeck.currentBullets())
	s.ebitenCanonDeck.update()
	if s.ebitenCanonDeck.Firing {
		stack.Push(FireCannonsState{
			ebitenCanonDeck:   s.ebitenCanonDeck,
			ebitenMonsterTeam: s.ebitenMonsterTeam,
		})
	}

	return nil
}

func (s BattleState) Draw(screen *ebiten.Image) {
	s.ebitenBattleGround.draw(screen)
	s.ebitenCanonDeck.draw(screen)
	s.ebitenMonsterTeam.draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
