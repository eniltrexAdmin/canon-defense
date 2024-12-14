package ebiten_sprite

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func EbitenStrokeStarted() (Stroke, bool) {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		return newMouseStroke(), true
	}
	var touchIDs []ebiten.TouchID

	touchIDs = inpututil.AppendJustPressedTouchIDs(touchIDs[:0])
	for _, id := range touchIDs {
		return newTouchStroke(id), true
	}

	return nil, false
}

type MouseStroke struct {
	IniX, IniY int
}

func newMouseStroke() MouseStroke {
	x, y := ebiten.CursorPosition()
	return MouseStroke{
		IniX: x,
		IniY: y,
	}
}

func (m MouseStroke) CurrentPosition() (x, y int) {
	return ebiten.CursorPosition()
}

func (m MouseStroke) DragDelta() (x, y float64) {
	currentX, currentY := ebiten.CursorPosition()
	deltaX := currentX - m.IniX
	deltaY := currentY - m.IniY
	return float64(deltaX), float64(deltaY)
}

func (m MouseStroke) IsJustReleased() bool {
	return inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft)
}

type TouchStroke struct {
	ID         ebiten.TouchID
	IniX, IniY int
}

func newTouchStroke(ID ebiten.TouchID) TouchStroke {
	x, y := ebiten.TouchPosition(ID)
	return TouchStroke{
		ID:   ID,
		IniX: x,
		IniY: y,
	}
}

func (t TouchStroke) CurrentPosition() (x, y int) {
	return ebiten.TouchPosition(t.ID)
}

func (t TouchStroke) DragDelta() (x, y float64) {
	currentX, currentY := ebiten.TouchPosition(t.ID)
	deltaX := currentX - t.IniX
	deltaY := currentY - t.IniY
	return float64(deltaX), float64(deltaY)
}

func (t TouchStroke) IsJustReleased() bool {
	return inpututil.IsTouchJustReleased(t.ID)
}
