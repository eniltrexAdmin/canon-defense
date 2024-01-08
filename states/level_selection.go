package states

import (
	"canon-tower-defense/pkg"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type LevelSelection struct {
	scrollOffset int
}

func NewLevelSelection() *LevelSelection {
	return &LevelSelection{scrollOffset: 0}
}

func (p *LevelSelection) Update(stack *pkg.StateStack, keys []ebiten.Key) error {
	return nil
}

func (p *LevelSelection) Draw(screen *ebiten.Image) {
	ebitenutil.DebugPrint(screen, "LEVEL SELECTION SCREEN")
	levelSquareSize := 30
	screenHeight := screen.Bounds().Dy()

	screen.Fill(color.RGBA{R: 0xff, G: 0xff, B: 0xff, A: 0xff})

	// Draw level indicators
	for i := 0; i < 10; i++ {
		y := (i * levelSquareSize) - p.scrollOffset
		if y < screenHeight && y > -levelSquareSize {
			// Draw square for each level
			vector.DrawFilledRect(
				screen,
				float32(10),
				float32(y),
				float32(levelSquareSize),
				float32(levelSquareSize),
				color.RGBA{R: 0xcc, G: 0xcc, B: 0xcc, A: 0xff},
				false,
			)

			// Draw level number in the center of the square
			textX := 10 + levelSquareSize/2 - 10
			textY := y + levelSquareSize/2 + 5
			ebitenutil.DebugPrintAt(screen, fmt.Sprintf("%d", i+1), textX, textY)
		}
	}

	// Draw dots or circles for each level (you can customize the style)
	for i := 0; i < 10; i++ {
		y := (i * levelSquareSize) - p.scrollOffset
		if y < screenHeight && y > -levelSquareSize {
			// Draw dots/circles below each level square
			dotX := 30
			dotY := y + levelSquareSize + 10

			// You can customize the style of the dots/circles
			vector.DrawFilledRect(screen, float32(dotX), float32(dotY), 10, 10, color.RGBA{A: 0xff}, false)
			// Or draw circles using ebitenutil.DrawArc
		}
	}
}
