package ebiten_sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type EbitenSprite struct {
	PosX       float64
	PosY       float64
	width      float64
	height     float64
	Image      *ebiten.Image
	imageScale float64
}

func NewEbitenSprite(
	posX float64,
	posY float64,
	width float64,
	height float64,
	image *ebiten.Image,
	imageScale float64,
) EbitenSprite {
	return EbitenSprite{
		PosX:       posX,
		PosY:       posY,
		width:      width,
		height:     height,
		Image:      image,
		imageScale: imageScale,
	}
}

func NewFromCentralPointImageAndScaleToExpected(
	centralX, centralY float64,
	image *ebiten.Image,
	expectedWidth float64,
) EbitenSprite {
	imgWidth := image.Bounds().Dx()
	imgHeight := image.Bounds().Dy()

	imageScale := expectedWidth / float64(imgWidth)

	newWidth := float64(imgWidth) * imageScale
	newHeight := float64(imgHeight) * imageScale

	posX := centralX - newWidth/2
	posY := centralY - newHeight/2

	return EbitenSprite{
		PosX:       posX,
		PosY:       posY,
		width:      newWidth,
		height:     newHeight,
		Image:      image,
		imageScale: imageScale,
	}
}

func (b EbitenSprite) InBounds(xInt, yInt int) bool {
	x := float64(xInt)
	y := float64(yInt)
	return x >= b.PosX && x <= b.PosX+b.width && y >= b.PosY && y <= b.PosY+b.height
}

func (b EbitenSprite) Draw(screen *ebiten.Image) {
	if b.Image == nil {
		vector.DrawFilledRect(screen,
			float32(b.PosX),
			float32(b.PosY),
			float32(b.width),
			float32(b.height),
			color.White,
			false)
	} else {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(b.imageScale, b.imageScale)
		op.GeoM.Translate(b.PosX, b.PosY)
		screen.DrawImage(b.Image, &op)
	}
}

func (b EbitenSprite) DrawDebug(screen *ebiten.Image) {
	if b.Image != nil {
		op := ebiten.DrawImageOptions{}
		op.GeoM.Scale(b.imageScale, b.imageScale)
		op.GeoM.Translate(b.PosX, b.PosY)
		screen.DrawImage(b.Image, &op)
	}

	vector.DrawFilledRect(screen,
		float32(b.PosX),
		float32(b.PosY),
		float32(b.width),
		float32(b.height),
		color.White,
		false)
}

type highland struct {
	EbitenSprite
}
