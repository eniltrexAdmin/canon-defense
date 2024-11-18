package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_background"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_monster"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

type BattleState struct {
	game               *game.CanonTDGame
	ebitenCanonDeck    *ebiten_canon.EbitenCanonDeck
	ebitenBattleGround ebiten_background.EbitenBattleGround
	ebitenMonsterTeam  *ebiten_monster.EbitenMonsterTeam
}

func NewBattleState(level int) BattleState {
	// loading assets, could be in init() and consistent usage of states.
	g := game.Start(level)

	ebiten_monster.LoadBattleImages()

	ecd := ebiten_canon.NewEbitenCanonDeck(g)
	ebg := ebiten_background.NewEbitenBattleGround(g.Battleground)
	emt := ebiten_monster.NewEbitenMonsterTeam(g)

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
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.ebitenCanonDeck.InitDrag()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		s.ebitenCanonDeck.ReleaseDrag()
		if s.ebitenCanonDeck.Firing {
			fireState := NewFireCannonState(s.ebitenCanonDeck, s.ebitenMonsterTeam, s.game)
			stack.Push(fireState)
		}
	}

	s.ebitenMonsterTeam.Update()
	s.ebitenCanonDeck.Update()

	return nil
}

func (s BattleState) Draw(screen *ebiten.Image) {
	s.ebitenBattleGround.Draw(screen)
	s.ebitenCanonDeck.Draw(screen)
	s.ebitenMonsterTeam.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
