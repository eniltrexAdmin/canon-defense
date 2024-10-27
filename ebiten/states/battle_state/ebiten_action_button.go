package battle_state

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const actionButtonYPlacement float64 = 600
const actionButtonTileSize float64 = 50

type ebitenActionButton struct {
	canonSprite                          ebiten_sprite.EbitenSprite
	pedestalSprite                       ebiten_sprite.EbitenSprite
	initialPlacementX, initialPlacementY float64
	dragIniX, dragIniY                   float64
	dragged                              bool
}

func newEbitenActionButton(cImage *ebiten.Image, cpImage *ebiten.Image, availableWidth int) ebitenActionButton {
	canonSprite := ebiten_sprite.NewFromCentralPointImageAndScaleToExpected(
		float64(availableWidth/2),
		actionButtonYPlacement,
		cImage,
		actionButtonTileSize,
	)

	pedestalSprite := ebiten_sprite.NewFromCentralPointImageAndScaleToExpected(
		float64(availableWidth/2),
		actionButtonYPlacement+36,
		cpImage,
		actionButtonTileSize,
	)

	return ebitenActionButton{
		canonSprite:       canonSprite,
		pedestalSprite:    pedestalSprite,
		dragged:           false,
		initialPlacementX: canonSprite.PosX,
		initialPlacementY: canonSprite.PosY,
	}
}

func (ec *ebitenActionButton) click(x, y int) {
	if ec.canonSprite.InBounds(x, y) {
		println("canon being dragged")
		ec.dragIniX = float64(x)
		ec.dragIniY = float64(y)
		ec.dragged = true
	}
	fmt.Printf("coordinates %f, %f ", ec.canonSprite.PosX, ec.canonSprite.PosY)
}

func (ec *ebitenActionButton) JustRelease() {
	ec.canonSprite.PosX = ec.initialPlacementX
	ec.canonSprite.PosY = ec.initialPlacementY
	ec.dragged = false
}

func (ec *ebitenActionButton) update(deck *ebitenCanonDeck) {
	if ec.dragged == false {
		return
	}
	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		deck.deploy(ebiten.CursorPosition())
		ec.JustRelease()
		return
	}

	// that means here it's still being dragged.
	x, y := ebiten.CursorPosition()

	dragDeltaX := float64(x) - ec.dragIniX + ec.initialPlacementX
	dragDeltaY := float64(y) - ec.dragIniY + ec.initialPlacementY

	ec.canonSprite.PosX = dragDeltaX
	ec.canonSprite.PosY = dragDeltaY
}

func (ec *ebitenActionButton) draw(screen *ebiten.Image) {
	ec.canonSprite.Draw(screen)
	ec.pedestalSprite.Draw(screen)
}
