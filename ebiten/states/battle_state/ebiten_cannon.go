package battle_state

import (
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const canonYPlacement float32 = 500
const canonTileSize float32 = 24

type ebitenCanon struct {
	screenBlock
	placement          int
	canon              *game.Canon
	canonPlacedImage   *ebiten.Image
	canonPedestalImage *ebiten.Image
}

func newEbitenCanon(canon *game.Canon, cImage *ebiten.Image, placement int, availableWidth int) ebitenCanon {

	imgHeight := cImage.Bounds().Dy()
	imgWidth := cImage.Bounds().Dx()

	centerSpace := availableWidth / 2
	imgStartingXPoint := centerSpace - imgWidth/2

	fmt.Println(fmt.Sprintf("placing canon in: %.2f",
		float32((availableWidth*placement)+imgStartingXPoint)),
	)

	return ebitenCanon{
		screenBlock: screenBlock{
			posX:   float32((availableWidth * placement) + imgStartingXPoint),
			posY:   canonYPlacement,
			width:  float32(imgWidth),
			height: float32(imgHeight),
		},
		placement:          placement,
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
	op.GeoM.Translate(float64(ec.posX), float64(ec.posY))
	screen.DrawImage(ec.canonPlacedImage, op)
	//ec.screenBlock.draw(screen)
}
