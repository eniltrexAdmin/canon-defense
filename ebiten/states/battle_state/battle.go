package battle_state

import (
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"strconv"
)

type BattleState struct {
	game        *game.CanonTDGame
	visibleRows int
	columns     int
	// IMPORTANT if I am modifying this objects inside like canon deck and battle ground through the
	// battle state UPDATE function, update must have a POINTER receiver!!!
	// (or else it's just a copy edited and thrown away).
	// or whatever gets modified needs to be referenced.
	ebitenCanonDeck    *ebitenCanonDeck
	ebitenBattleGround *ebitenBattleGround
}

func NewBattleState(level int) BattleState {
	// loading assets, could be in init() and consistent usage of states.
	g := game.Start(level)

	ecd := newEbitenCanonDeck(g.CanonDeck)
	ebg := newEbitenBattleGround(g.Battleground)

	return BattleState{
		game:               &g,
		visibleRows:        5, // TODO that should come from game.
		columns:            int(g.Battleground.Columns),
		ebitenCanonDeck:    &ecd,
		ebitenBattleGround: &ebg,
	}
}

func (s BattleState) Debug() string {
	return "BattleState State"
}

func (s BattleState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ec := s.ebitenCanonDeck.click(ebiten.CursorPosition())
		if ec != nil {
			println("pressed cannon: " + strconv.Itoa(ec.formationPlacement))
		}
	}

	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		s.ebitenCanonDeck.actionButton.click(ebiten.CursorPosition())
	}

	s.ebitenCanonDeck.actionButton.update()

	s.ebitenBattleGround.update()
	return nil
}

func (s BattleState) Draw(screen *ebiten.Image) {
	s.ebitenBattleGround.draw(screen)
	s.ebitenCanonDeck.draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
