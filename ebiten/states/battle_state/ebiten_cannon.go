package battle_state

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const canonYPlacement float64 = 500

type ebitenCanon struct {
	ebiten_sprite.EbitenSprite
	formationPlacement int
	canon              *game.Canon
	canonPlacedImage   *ebiten.Image
	canonPedestalImage *ebiten.Image
}

func newEbitenCanon(canon *game.Canon, cImage *ebiten.Image, formationPlacement int, availableWidth int) ebitenCanon {

	imgHeight := cImage.Bounds().Dy()
	imgWidth := cImage.Bounds().Dx()

	centerSpace := availableWidth / 2
	imgStartingXPoint := centerSpace - imgWidth/2

	fmt.Println(fmt.Sprintf("placing canon in: %.2f",
		float32((availableWidth*formationPlacement)+imgStartingXPoint)),
	)

	sprite := ebiten_sprite.NewEbitenSprite(
		float64((availableWidth*formationPlacement)+imgStartingXPoint),
		canonYPlacement,
		float64(imgWidth),
		float64(imgHeight),
		nil,
		1,
	)

	return ebitenCanon{
		EbitenSprite:       sprite,
		formationPlacement: formationPlacement,
		canon:              canon,
		canonPlacedImage:   cImage,
		canonPedestalImage: nil,
	}
}

func (ec ebitenCanon) draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}

	var fill float32 = 0.5
	if ec.canon != nil {
		fill = 1
	}
	op.ColorScale.ScaleAlpha(fill)
	op.GeoM.Translate(ec.PosX, ec.PosY)
	screen.DrawImage(ec.canonPlacedImage, op)
	//ec.screenBlock.draw(screen)
}
