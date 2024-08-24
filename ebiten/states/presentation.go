package states

import (
	"bytes"
	"canon-tower-defense/assets"
	"canon-tower-defense/ebiten/constants"
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

func (p *PresentationState) Update(stack *StateStack, keys []ebiten.Key) error {
	p.fadeIn += 0.01
	if p.fadeIn > 1 {
		p.fadeIn = 1
		stack.Pop()
	}
	return nil
}

func (p *PresentationState) Draw(screen *ebiten.Image) {
	// TODO this should be loaded once not on every draw tick
	img, _, err := ebitenutil.NewImageFromReader(bytes.NewReader(assets.LogoPng))
	if err != nil {
		log.Fatal(err)
	}
	bgImage := ebiten.NewImageFromImage(img)

	imgWidth := img.Bounds().Dx()
	imgHeight := img.Bounds().Dy()
	scaleX := float64(constants.ScreenWidth) / float64(imgWidth)
	scaleY := float64(constants.ScreenHeight) / float64(imgHeight)

	options := &ebiten.DrawImageOptions{}
	options.GeoM.Scale(scaleX, scaleY)

	options.ColorScale.ScaleAlpha(p.fadeIn)
	screen.Fill(color.Black)
	screen.DrawImage(bgImage, options)
}
