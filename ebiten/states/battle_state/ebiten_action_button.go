package battle_state

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

const actionButtonYPlacement float64 = 650
const actionButtonTileSize float64 = 50

type ebitenActionButton struct {
	canonSprite    *ebiten_sprite.EbitenDraggableSprite
	pedestalSprite ebiten_sprite.EbitenSprite
}

func newEbitenActionButton(cImage *ebiten.Image, cpImage *ebiten.Image, availableWidth int) ebitenActionButton {
	canonSprite := ebiten_sprite.NewFromCentralPointScaleImage(
		float64(availableWidth/2),
		actionButtonYPlacement,
		cImage,
		actionButtonTileSize,
	)

	pedestalSprite := ebiten_sprite.NewFromCentralPointScaleImage(
		float64(availableWidth/2),
		actionButtonYPlacement+36,
		cpImage,
		actionButtonTileSize,
	)

	return ebitenActionButton{
		canonSprite:    ebiten_sprite.NewFromSprite(canonSprite),
		pedestalSprite: pedestalSprite,
	}
}

func (ec *ebitenActionButton) update() {
	ec.canonSprite.Update()
}

func (ec *ebitenActionButton) draw(screen *ebiten.Image) {
	ec.canonSprite.Draw(screen)
	ec.pedestalSprite.Draw(screen)
}
