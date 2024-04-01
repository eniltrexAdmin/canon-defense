package states

import (
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

const HighlandSize = 50

type BattleState struct {
	battleGround game.Battleground
	visibleRows  int
	columns      int
}

func NewBattleState(level int) BattleState {
	return BattleState{
		battleGround: game.HardcodedLevelGenerator{}.Generate(level),
		visibleRows:  5,
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
	for i := 0; i < s.visibleRows; i++ {
		for j := 0; j < s.columns; j++ {
			drawSquare(screen, j*HighlandSize, i*HighlandSize)
		}
	}
}

func drawSquare(screen *ebiten.Image, x, y int) {
	highland := ebiten.NewImage(HighlandSize, HighlandSize)
	//t(dst *ebiten.Image, x, y, width, height float32, clr color.Color, antialias bool)
	vector.DrawFilledRect(highland, float32(x), float32(y), HighlandSize, HighlandSize, color.White, true)
}
