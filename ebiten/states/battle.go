package states

import (
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const HighlandSize = 100

type BattleState struct {
	battleGround game.Battleground
	visibleRows  int
	columns      int
	layout       battleLayout
}

type battleLayout struct {
	layout []*ebiten.Image
}

func NewBattleState(level int) BattleState {
	return BattleState{
		battleGround: game.HardcodedLevelGenerator{}.Generate(level),
		visibleRows:  3,
		columns:      5,
	}
}

func (s BattleState) Debug() string {
	return "BattleState State"
}

func (s BattleState) Update(stack *StateStack, keys []ebiten.Key) error {
	return nil
}

func (s BattleState) Draw(screen *ebiten.Image) {

	padding := 10

	for i := 0; i < s.visibleRows; i++ {
		for j := 0; j < s.columns; j++ {
			vector.DrawFilledRect(screen,
				float32(j*HighlandSize+padding+s.centerBattleFieldPadding(screen)),
				float32(i*HighlandSize+padding),
				HighlandSize/2, HighlandSize/2, color.White, false)
		}
	}
}

func (s BattleState) centerBattleFieldPadding(screen *ebiten.Image) int {
	// 10 is padding, it all will go away when drawing each sqaure goes to the object.
	battlefieldLengthMiddle := (HighlandSize)*s.columns + (10 * (s.columns - 1))
	middlePointWidth := screen.Bounds().Size().X
	println((middlePointWidth - battlefieldLengthMiddle) / 2)
	return (middlePointWidth - battlefieldLengthMiddle) / 2
}

func drawSquare(screen *ebiten.Image, x, y int) {
	col := color.White
	//if x%2 == 0 {
	//	col = color.Black
	//}
	//t(dst *ebiten.Image, x, y, width, height float32, clr color.Color, antialias bool)
	println(x, y)
	vector.DrawFilledRect(screen, float32(x), float32(y), HighlandSize, HighlandSize, col, false)
}
