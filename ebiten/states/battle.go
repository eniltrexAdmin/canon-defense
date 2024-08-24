package states

import (
	"bytes"
	"canon-tower-defense/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
	"log"
	"math"
)

//const HighlandSize = 100

type BattleState struct {
	game        game.CanonTDGame
	visibleRows int
	columns     int
	layout      battleLayout // whats that for??
	images      battleStateImages
}

type battleLayout struct {
	layout []*ebiten.Image
}

type battleStateImages struct {
	canonImage *ebiten.Image
}

func NewBattleState(level int) BattleState {
	// loading assets, could be in init() and consistent usage of states.
	img, _, err := image.Decode(bytes.NewReader(assets.RegularCanon))
	if err != nil {
		log.Fatal(err)
	}
	canonImage := ebiten.NewImageFromImage(img)
	imgs := battleStateImages{canonImage: canonImage}
	g := game.Start(level)

	return BattleState{
		game:        g,
		visibleRows: 3,
		columns:     int(g.Battleground.Columns),
		images:      imgs,
	}
}

func (s BattleState) Debug() string {
	return "BattleState State"
}

func (s BattleState) Update(stack *StateStack, keys []ebiten.Key) error {
	x, y := ebiten.CursorPosition()
	// Calculate the size of each Highland square
	HighlandSize := constants.ScreenWidth / s.columns

	// Check if the mouse click is within any of the canon's boundaries
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		for j := range s.game.CanonDeck.Canons {
			// Calculate the position of the top-left corner of the canon image
			canonX := float64(j*HighlandSize + HighlandSize/2 + 10)
			canonY := float64(400)

			// Calculate the boundaries of the square (before rotation)
			left := canonX - float64(s.images.canonImage.Bounds().Dy())/2
			right := canonX + float64(s.images.canonImage.Bounds().Dy())/2
			top := canonY - float64(s.images.canonImage.Bounds().Dx())/2
			bottom := canonY + float64(s.images.canonImage.Bounds().Dx())/2

			// Check if the mouse click is within these boundaries
			if float64(x) >= left && float64(x) <= right && float64(y) >= top && float64(y) <= bottom {
				// The mouse is inside the square - do something
				s.game.PlaceCannon(j)
				fmt.Println("Clicked on canon at position:", j)
			}
		}
	}

	return nil
}

func (s BattleState) Draw(screen *ebiten.Image) {

	padding := 50

	HighlandSize := constants.ScreenWidth / s.columns

	for i := 0; i < s.visibleRows; i++ {
		for j := 0; j < s.columns; j++ {
			vector.DrawFilledRect(screen, float32(j*HighlandSize+padding/2), float32(i*HighlandSize+padding/2),
				float32(HighlandSize-padding), float32(HighlandSize-padding), color.White, false)
		}
	}
	s.DrawCannonDeck(screen)
	ebitenutil.DebugPrint(screen, fmt.Sprintf("TPS: %0.2f", ebiten.ActualTPS()))
}

//func (s BattleState) DrawBattleRow(screen *ebiten.Image, row int) {
//
//}

func (s BattleState) DrawCannonDeck(screen *ebiten.Image) {
	HighlandSize := constants.ScreenWidth / s.columns

	for j, canon := range s.game.CanonDeck.Canons {
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Rotate(0.5 * math.Pi)

		var fill float32 = 0.5
		if canon != nil {
			fill = 1
		}
		op.ColorScale.ScaleAlpha(fill)
		// from the beginning of the grid, but centered, rotation complicates calculation,
		op.GeoM.Translate(float64((j)*HighlandSize+HighlandSize/2+10), float64(400))

		screen.DrawImage(s.images.canonImage, op)
	}
}
