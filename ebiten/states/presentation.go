package states

import (
	"bytes"
	"canon-tower-defense/assets"
	_ "embed"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"image/color"
	_ "image/png"
	"log"
)

type PresentationState struct {
	fadeIn float32
}

func NewPresentationState() *PresentationState {
	return &PresentationState{fadeIn: 0}
}

func (p *PresentationState) Debug() string {
	return "Presentation State"
}

func (p *PresentationState) Update(stack *StateStack, keys []ebiten.Key) error {
	p.fadeIn += 0.01
	if p.fadeIn > 1 {
		p.fadeIn = 1
		stack.Pop()
	}
	return nil
}

func (p *PresentationState) Draw(screen *ebiten.Image) {
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.LogoPng))
	if err != nil {
		log.Fatal(err)
	}
	bgImage := ebiten.NewImageFromImage(img)

	// Calculate scaling factors to fit the window
	screenWidth := screen.Bounds().Dx()
	screenHeight := screen.Bounds().Dy()
	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	scaleX := float64(screenWidth) / float64(imgWidth)
	scaleY := float64(screenHeight) / float64(imgHeight)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(scaleX, scaleY)

	options.ColorScale.ScaleAlpha(p.fadeIn)
	screen.Fill(color.Black)
	screen.DrawImage(bgImage, options)
}
