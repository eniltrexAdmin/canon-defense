package ebiten_sprite

import (
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

// About colors

func RandomColor() color.Color {
	// Generate random RGB values
	r := uint8(rand.Intn(256))
	g := uint8(rand.Intn(256))
	b := uint8(rand.Intn(256))

	// Create a random color
	return color.RGBA{R: r, G: g, B: b, A: 0}
}
