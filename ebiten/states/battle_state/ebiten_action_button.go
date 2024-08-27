package battle_state

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

const actionButtonYPlacement float32 = 600
const actionButtonTileSize float32 = 50

type ebitenActionButton struct {
	screenBlock
	canonPlacedImage            *ebiten.Image
	canonPlacedImageDrawOptions ebiten.DrawImageOptions
	canonPedestalImage          *ebiten.Image
	canonPedestalDrawOptions    ebiten.DrawImageOptions
}

func newEbitenActionButton(cImage *ebiten.Image, cpImage *ebiten.Image, availableWidth int) ebitenActionButton {

	centerSpace := availableWidth / 2

	canonImgWidth := cImage.Bounds().Dx()
	canonImgHeight := cImage.Bounds().Dy()
	fmt.Printf("Width: %d, Height: %d\n", canonImgWidth, canonImgHeight)

	op := ebiten.DrawImageOptions{}
	scaleX := float64(actionButtonTileSize) / float64(canonImgWidth)
	op.GeoM.Scale(scaleX, scaleX)

	newWidth := float64(canonImgWidth) * scaleX
	newHeight := float64(canonImgHeight) * scaleX
	fmt.Printf("New Width: %f, New Height: %f\n", newWidth, newHeight)
	imgStartingXPoint := centerSpace - int(newWidth)/2

	op.GeoM.Translate(float64(imgStartingXPoint), float64(actionButtonYPlacement))

	placementWidth := float64(cpImage.Bounds().Dx()) * scaleX
	placeImgStartingXPoint := centerSpace - int(placementWidth)/2

	opPedestal := ebiten.DrawImageOptions{}
	opPedestal.GeoM.Scale(scaleX, scaleX)
	opPedestal.GeoM.Translate(float64(placeImgStartingXPoint), float64(actionButtonYPlacement)+newHeight-4)

	return ebitenActionButton{
		screenBlock: screenBlock{
			posX:   float32(imgStartingXPoint),
			posY:   actionButtonYPlacement,
			width:  float32(newWidth),
			height: float32(newHeight),
			image:  nil,
		},
		canonPlacedImage:            cImage,
		canonPlacedImageDrawOptions: op,
		canonPedestalImage:          cpImage,
		canonPedestalDrawOptions:    opPedestal,
	}
}

func (ec ebitenActionButton) draw(screen *ebiten.Image) {
	//op := &ebiten.DrawImageOptions{}
	//
	//op.GeoM.Translate(float64(ec.posX), float64(ec.posY))
	screen.DrawImage(ec.canonPlacedImage, &ec.canonPlacedImageDrawOptions)

	screen.DrawImage(ec.canonPedestalImage, &ec.canonPedestalDrawOptions)
	//ec.screenBlock.draw(screen)
}
