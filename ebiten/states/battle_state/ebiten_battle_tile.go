package battle_state

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
)

type ebitenBattleTile struct {
	ebiten_sprite.EbitenSprite
	monster *game.Monster
}

const tileSize float64 = 50

func (ebt ebitenBattleTile) draw(screen *ebiten.Image) {
	ebt.EbitenSprite.Draw(screen)
}

func newEbitenBattleTile(m *game.Monster, row, column int, availableWidth, availableHeight int) ebitenBattleTile {
	centerXSpace := availableWidth / 2
	tileStartingPointX := float64(centerXSpace) - tileSize/2

	centerYSpace := availableHeight / 2
	tileStartingPointY := float64(centerYSpace) - tileSize/2

	sprite := ebiten_sprite.NewEbitenSprite(
		float64(availableWidth*column)+tileStartingPointX,
		float64(availableHeight*row)+tileStartingPointY,
		tileSize,
		tileSize,
		nil,
		1,
	)

	return ebitenBattleTile{
		EbitenSprite: sprite,
		monster:      m,
	}
}
