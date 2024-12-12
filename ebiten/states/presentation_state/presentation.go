package presentation_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/states"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	_ "image/png"
	"log"
)

type PresentationState struct {
	fadeIn    float32
	logoImage *ebiten.Image
}

func NewPresentationState() *PresentationState {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.LogoPng))
	if err != nil {
		log.Fatal(err)
	}
	bgImage := ebiten.NewImageFromImage(img)
	return &PresentationState{fadeIn: 0, logoImage: bgImage}
}

func (p *PresentationState) Debug() string {
	return "Presentation State"
}

func (p *PresentationState) Update(stack *states.StateStack, keys []ebiten.Key) error {
	p.fadeIn += 0.01
	if p.fadeIn > 1 {
		p.fadeIn = 1
		stack.Pop()
	}
	return nil
}

func (p *PresentationState) Draw(screen *ebiten.Image) {

	imgWidth := p.logoImage.Bounds().Dx()
	// Calculate the scale factor based on width only
	scaleX := float64(constants.ScreenWidth) / float64(imgWidth)

	// Scale the Y-axis to maintain the aspect ratio
	scaleY := scaleX

	imgHeight := p.logoImage.Bounds().Dy()
	newHeight := scaleY * float64(imgHeight)

	// Calculate the offset to center the image vertically
	offsetY := (float64(constants.ScreenHeight) - newHeight) / 2

	// Set up the transformation options
	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(scaleX, scaleY)
	options.GeoM.Translate(0, offsetY)

	options.ColorScale.ScaleAlpha(p.fadeIn)
	screen.Fill(color.Black)
	screen.DrawImage(p.logoImage, options)
}
