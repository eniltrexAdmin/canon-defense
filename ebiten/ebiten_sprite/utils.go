package ebiten_sprite

import (
	"canon-tower-defense/ebiten/constants"
	"image"
	"image/color"
	"math/rand"
)

type Rectangle interface {
	GetRectangle() image.Rectangle
}

func Collision(sprite1, sprite2 Rectangle) bool {
	return sprite1.GetRectangle().Overlaps(sprite2.GetRectangle())
}

func SpriteInScreen(sprite1 Rectangle) bool {
	screenRect := image.Rect(0, 0, constants.ScreenWidth, constants.ScreenHeight)
	return sprite1.GetRectangle().Overlaps(screenRect)
}

// About colors

func RandomColor() color.Color {
	// Generate random RGB values
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	// Create a random color
	return color.RGBA{R: r, G: g, B: b, A: 0}
}
