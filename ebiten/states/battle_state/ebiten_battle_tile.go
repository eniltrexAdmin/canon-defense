package battle_state

import (
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type ebitenBattleTile struct {
	screenBlock
	monster *game.Monster
}

const tileSize float32 = 50

func (ebt ebitenBattleTile) draw(screen *ebiten.Image) {
	ebt.screenBlock.draw(screen)
}

func newEbitenBattleTile(m *game.Monster, row, column int, availableWidth, availableHeight int) ebitenBattleTile {
	centerXSpace := availableWidth / 2
	tileStartingPointX := float32(centerXSpace) - tileSize/2

	centerYSpace := availableHeight / 2
	tileStartingPointY := float32(centerYSpace) - tileSize/2

	return ebitenBattleTile{
		screenBlock: screenBlock{
			posX:   float32(availableWidth*column) + tileStartingPointX,
			posY:   float32(availableHeight*row) + tileStartingPointY,
			width:  tileSize,
			height: tileSize,
		},
		monster: m,
	}
}
