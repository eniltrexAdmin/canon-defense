package ebiten_sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"log"
	"math"
)

type EbitenAnimatedSprite struct {
	PosX           float64
	PosY           float64
	width          float64
	height         float64
	Image          *ebiten.Image
	imageScale     float64
	subImageWidth  int
	subImageHeight int
	currentFrame   float64
	totalFrames    int
	animationSpeed float64
	debugColor     color.Color
}

func NewFromCentralPoint(
	centralX, centralY float64,
	image *ebiten.Image,
	subImageWidth int,
	subImageHeight int,
	scaleFactor float64,
	animationSpeed float64,
) EbitenAnimatedSprite {
	intendedWidth := float64(subImageWidth) * scaleFactor
	intendedHeight := float64(subImageHeight) * scaleFactor

	posX := centralX - intendedWidth/2
	posY := centralY - intendedHeight/2

	return EbitenAnimatedSprite{
		PosX:           posX,
		PosY:           posY,
		width:          intendedWidth,
		height:         intendedHeight,
		Image:          image,
		imageScale:     scaleFactor,
		subImageWidth:  subImageWidth,
		subImageHeight: subImageHeight,
		currentFrame:   0,
		totalFrames:    getTotalFrames(image, subImageWidth),
		animationSpeed: animationSpeed,
		debugColor:     RandomColor(),
	}
}

func getTotalFrames(image *ebiten.Image, subImageWidth int) int {
	imgWidth := image.Bounds().Dx()

	if subImageWidth <= 0 {
		panic("subImageWidth must be greater than 0")
	}

	if imgWidth%subImageWidth != 0 {
		log.Printf("Warning: Image width %d is not an exact multiple of subImageWidth %d. "+
			"Frames may not cover the entire image.", imgWidth, subImageWidth)
	}
	return imgWidth / subImageWidth
}

func (e *EbitenAnimatedSprite) getRectangle() image.Rectangle {
	x := int(e.PosX)
	y := int(e.PosY)
	width := int(e.width)
	height := int(e.height)
	return image.Rect(x, y, x+width, y+height)
}

func (e *EbitenAnimatedSprite) Update() {
	e.currentFrame += e.animationSpeed
	if e.currentFrame >= float64(e.totalFrames) {
		e.currentFrame = 0
	}
}

func (e *EbitenAnimatedSprite) Draw(screen *ebiten.Image) {
	op := ebiten.DrawImageOptions{}
	op.GeoM.Scale(e.imageScale, e.imageScale)
	op.GeoM.Translate(e.PosX, e.PosY)
	screen.DrawImage(e.Image.SubImage(e.getCurrentSubImage()).(*ebiten.Image), &op)
}

func (e *EbitenAnimatedSprite) getCurrentSubImage() image.Rectangle {
	i := int(math.Floor(e.currentFrame))
	initialX := i * e.subImageWidth
	return image.Rect(initialX, 0, initialX+e.subImageWidth, e.subImageHeight)
}

func (e *EbitenAnimatedSprite) DrawDebug(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		float32(e.PosX),
		float32(e.PosY),
		float32(e.width),
		float32(e.height),
		e.debugColor,
		false)
}
