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
	transitionToFire   bool
}

func NewBattleState(level int) *BattleState {
	// loading assets, could be in init() and consistent usage of states.
	g := game.Start(level)

	ebiten_monster.LoadBattleImages()
	ebiten_canon.LoadBattleImages()

	ebg := ebiten_background.NewEbitenBattleGround(g.Battleground)
	emt := ebiten_monster.NewEbitenMonsterTeam(&g)

	bs := BattleState{
		game:               &g,
		ebitenBattleGround: ebg,
		ebitenMonsterTeam:  &emt,
		transitionToFire:   false,
	}

	ecd := ebiten_canon.NewEbitenCanonDeck(&g, bs.deployCannon, bs.moveCannon)
	bs.ebitenCanonDeck = &ecd

	return &bs
}

func (s *BattleState) Debug() string {
	return "battleState State"
}

func (s *BattleState) deployCannon(on int) {
	s.game.DeployCannon(on)
	s.transitionToFire = true
}

func (s *BattleState) moveCannon(from, to int) {
	s.game.MoveCannon(from, to)
	s.transitionToFire = true
}

func (s *BattleState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.ebitenCanonDeck.InitDrag()
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		s.ebitenCanonDeck.ReleaseDrag()
	}

	s.ebitenMonsterTeam.Update()
	s.ebitenCanonDeck.Update()

	if s.transitionToFire {
		s.transitionToFire = false
		fireState := NewFireCannonState(s.ebitenCanonDeck, s.ebitenMonsterTeam, s.game)
		stack.Push(fireState)
	}

	return nil
}

func (s *BattleState) Draw(screen *ebiten.Image) {
	s.ebitenBattleGround.Draw(screen)
	s.ebitenCanonDeck.Draw(screen)
	s.ebitenMonsterTeam.Draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
