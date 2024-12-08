package ebiten_sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"log"
	"math"
)

type ScreenCoordinate struct {
	X float64
	Y float64
}

type AnimatedSprite struct {
	Image                         *ebiten.Image
	SubImageWidth, SubImageHeight int
}

func NewAnimatedSprite(i *ebiten.Image, subImageWidth, subImageHeight int) AnimatedSprite {
	return AnimatedSprite{
		Image:          i,
		SubImageWidth:  subImageWidth,
		SubImageHeight: subImageHeight,
	}
}

func (sc ScreenCoordinate) Equals(sc2 ScreenCoordinate) bool {
	return sc.X == sc2.X && sc.Y == sc2.Y
}

type EbitenAnimatedSprite struct {
	position       *ScreenCoordinate
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
	sprite AnimatedSprite,
	scaleFactor float64,
	animationSpeed float64,
) EbitenAnimatedSprite {
	cp := ScreenCoordinate{
		X: centralX,
		Y: centralY,
	}

	intendedWidth := float64(sprite.SubImageWidth) * scaleFactor
	intendedHeight := float64(sprite.SubImageHeight) * scaleFactor

	return EbitenAnimatedSprite{
		position:       &cp,
		width:          intendedWidth,
		height:         intendedHeight,
		Image:          sprite.Image,
		imageScale:     scaleFactor,
		subImageWidth:  sprite.SubImageWidth,
		subImageHeight: sprite.SubImageHeight,
		currentFrame:   0,
		totalFrames:    getTotalFrames(sprite.Image, sprite.SubImageWidth),
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

func (e *EbitenAnimatedSprite) Position() ScreenCoordinate {
	return *e.position
}

func (e *EbitenAnimatedSprite) topLeftCorner() (float64, float64) {
	posX := e.position.X - e.width/2
	posY := e.position.Y - e.height/2
	return posX, posY
}

func (e *EbitenAnimatedSprite) LinkSprite(e2 *EbitenAnimatedSprite) {
	e.position = e2.position
}

func (e *EbitenAnimatedSprite) GetRectangle() image.Rectangle {
	floatX, floatY := e.topLeftCorner()
	x := int(floatX)
	y := int(floatY)
	width := int(e.width)
	height := int(e.height)
	return image.Rect(x, y, x+width, y+height)
}

func (e *EbitenAnimatedSprite) Move(destination ScreenCoordinate, movementSpeed float64) {
	dx := destination.X - e.position.X
	dy := destination.Y - e.position.Y

	// Limit the movement to the remaining distance
	xMovement := math.Min(math.Abs(dx), movementSpeed) * math.Copysign(1, dx)
	yMovement := math.Min(math.Abs(dy), movementSpeed) * math.Copysign(1, dy)

	e.position.X += xMovement
	e.position.Y += yMovement
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
	op.GeoM.Translate(e.topLeftCorner())
	screen.DrawImage(e.Image.SubImage(e.getCurrentSubImage()).(*ebiten.Image), &op)
}

func (e *EbitenAnimatedSprite) DrawWithFade(screen *ebiten.Image, alfa float32) {
	op := ebiten.DrawImageOptions{}
	op.ColorScale.ScaleAlpha(alfa)
	op.GeoM.Scale(e.imageScale, e.imageScale)
	op.GeoM.Translate(e.topLeftCorner())
	screen.DrawImage(e.Image.SubImage(e.getCurrentSubImage()).(*ebiten.Image), &op)
}

func (e *EbitenAnimatedSprite) getCurrentSubImage() image.Rectangle {
	i := int(math.Floor(e.currentFrame))
	initialX := i * e.subImageWidth
	return image.Rect(initialX, 0, initialX+e.subImageWidth, e.subImageHeight)
}

func (e *EbitenAnimatedSprite) DrawDebug(screen *ebiten.Image) {
	x, y := e.topLeftCorner()
	vector.DrawFilledRect(screen,
		float32(x),
		float32(y),
		float32(e.width),
		float32(e.height),
		e.debugColor,
		false)
}
