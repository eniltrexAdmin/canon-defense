package battle_state

import (
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type ebitenBattleGround struct {
	tiles []ebitenBattleTile
}

func (ecd ebitenBattleGround) draw(screen *ebiten.Image) {
	for _, tile := range ecd.tiles {
		tile.draw(screen)
	}
}

func newEbitenBattleGround(bg game.Battleground) ebitenBattleGround {
	availableWidth := constants.ScreenWidth / int(bg.Columns)
	availableHeight := int(canonYPlacement) / int(bg.VisibleRows)

	var t []ebitenBattleTile

	for i := 0; i < int(bg.VisibleRows); i++ {
		for j := 0; j < int(bg.Columns); j++ {
			ti := newEbitenBattleTile(bg.Monsters[i][j], i, j, availableWidth, availableHeight)
			t = append(t, ti)
		}
	}
	return ebitenBattleGround{tiles: t}
}
