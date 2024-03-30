package states

import (
	"canon-tower-defense/game"
	"canon-tower-defense/game/player"
	"canon-tower-defense/pkg"
	"canon-tower-defense/ui"
	"fmt"
	"github.com/ebitenui/ebitenui"
	"github.com/ebitenui/ebitenui/widget"
	"github.com/hajimehoshi/ebiten/v2"
)

type LevelSelection struct {
	scrollOffset  int
	ui            *ebitenui.UI
	LevelSelector game.LevelSelector
}

func NewLevelSelection(pl player.Player, LevelSelector game.LevelSelector) *LevelSelection {
	l := LevelSelector.LevelSelection(pl)

	eui := &ebitenui.UI{
		Container: layout(l),
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

func layout(levels game.Levels) *widget.Container {

	c := widget.NewContainer(
		widget.ContainerOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.AnchorLayoutData{
			StretchHorizontal: true,
		})),
		widget.ContainerOpts.Layout(widget.NewRowLayout(
			widget.RowLayoutOpts.Direction(widget.DirectionVertical),
			widget.RowLayoutOpts.Padding(widget.Insets{
				Top:    10,
				Left:   10,
				Right:  10,
				Bottom: 0,
			}),
			widget.RowLayoutOpts.Spacing(10),
		)))

	buttonRes, _ := ui.NewButtonResources()
	for row := 0; row < len(levels); row++ {

		b := widget.NewButton(
			widget.ButtonOpts.WidgetOpts(widget.WidgetOpts.LayoutData(widget.RowLayoutData{
				Stretch: true,
			})),
			widget.ButtonOpts.Image(buttonRes.Image),
			widget.ButtonOpts.Text(fmt.Sprintf("%s %d", "Level", row+1), buttonRes.Face, buttonRes.Text),
			// specify that the button's text needs some padding for correct display
			widget.ButtonOpts.TextPadding(widget.Insets{
				Left:   30,
				Right:  30,
				Top:    5,
				Bottom: 5,
			}),

			// add a handler that reacts to clicking the button
			widget.ButtonOpts.ClickedHandler(func(args *widget.ButtonClickedEventArgs) {
				println(fmt.Sprintf("Button %d clicked", row+1))
			}),
		)
		b.GetWidget().Disabled = !levels[row]

		c.AddChild(b)
	}

	return c
}
