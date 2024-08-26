package battle_state

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
	"strconv"
)

//const HighlandSize = 100

type BattleState struct {
	game            game.CanonTDGame
	visibleRows     int
	columns         int
	ebitenCanonDeck ebitenCanonDeck
}

func NewBattleState(level int) BattleState {
	// loading assets, could be in init() and consistent usage of states.
	g := game.Start(level)

	return BattleState{
		game:            g,
		visibleRows:     5,
		columns:         int(g.Battleground.Columns),
		ebitenCanonDeck: newEbitenCanonDeck(g.CanonDeck, 50),
	}
}

func (s BattleState) Debug() string {
	return "BattleState State"
}

func (s BattleState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		ec := s.ebitenCanonDeck.click(ebiten.CursorPosition())
		if ec != nil {
			println("pressed cannon: " + strconv.Itoa(ec.placement))
		}
	}

	return nil
}

func (s BattleState) Draw(screen *ebiten.Image) {

	padding := 50

	HighlandSize := constants.ScreenWidth / s.columns

	for i := 0; i < s.visibleRows; i++ {
		for j := 0; j < s.columns; j++ {
			vector.DrawFilledRect(screen, float32(j*HighlandSize+padding/2), float32(i*HighlandSize+padding/2),
				float32(HighlandSize-padding), float32(HighlandSize-padding), color.White, false)
		}
	}
	s.ebitenCanonDeck.draw(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}
