package states

import (
	"canon-tower-defense/pkg"
	"canon-tower-defense/ui"
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelSelection struct {
	scrollOffset int
	ui           *ebitenui.UI
}

func NewLevelSelection() *LevelSelection {
	eui := &ebitenui.UI{
		Container: layout(),
	}

	return &LevelSelection{
		scrollOffset: 0,
		ui:           eui,
	}
}

func (p *LevelSelection) Update(stack *pkg.StateStack, keys []ebiten.Key) error {
	p.ui.Update()
	return nil
}

func (p *LevelSelection) Draw(screen *ebiten.Image) {
	p.ui.Draw(screen)
}

func layout() *widget.Container {
	buttonRes, _ := ui.NewButtonResources()
	c := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Spacing(10),
		)))

	bc := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
			Stretch: true,
		})),
		widget.ContainerOpts.Layout(widget.NewGridLayout(
			widget.GridLayoutOpts.Columns(4),
			widget.GridLayoutOpts.Stretch([]bool{true, true, true, true}, nil),
			widget.GridLayoutOpts.Spacing(10, 10))))
	c.AddChild(bc)

	i := 0
	for row := 0; row < 3; row++ {
		for col := 0; col < 4; col++ {
			b := widget.NewButton(
				widget.ButtonOpts.Image(buttonRes.Image),
				widget.ButtonOpts.Text(fmt.Sprintf("%s %d", string(rune('A'+i)), i+1), buttonRes.Face, buttonRes.Text))
			bc.AddChild(b)

			i++
		}
	}

	return c
}
