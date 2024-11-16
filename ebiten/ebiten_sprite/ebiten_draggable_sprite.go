package ebiten_sprite

import "github.com/hajimehoshi/ebiten/v2"

type EbitenDraggableSprite struct {
	initialPlacementX, initialPlacementY float64
	dragIniX, dragIniY                   float64
	IsDragged                            bool
	EbitenSprite
}

func NewFromSprite(sprite EbitenSprite) *EbitenDraggableSprite {
	return &EbitenDraggableSprite{
		initialPlacementX: sprite.PosX,
		initialPlacementY: sprite.PosY,
		dragIniX:          0,
		dragIniY:          0,
		IsDragged:         false,
		EbitenSprite:      sprite,
	}
}

func (ds *EbitenDraggableSprite) InitDrag() {
	// already couping with ebiten, probably better here
	x, y := ebiten.CursorPosition()
	if ds.EbitenSprite.InBounds(x, y) {
		ds.dragIniX = float64(x)
		ds.dragIniY = float64(y)
		ds.IsDragged = true
	}
}

func (ds *EbitenDraggableSprite) ReleaseDrag() {
	ds.EbitenSprite.PosX = ds.initialPlacementX
	ds.EbitenSprite.PosY = ds.initialPlacementY
	ds.IsDragged = false
}

func (ds *EbitenDraggableSprite) Update() {
	if ds.IsDragged == false {
		return
	}
	x, y := ebiten.CursorPosition()

	dragDeltaX := float64(x) - ds.dragIniX + ds.initialPlacementX
	dragDeltaY := float64(y) - ds.dragIniY + ds.initialPlacementY

	ds.EbitenSprite.PosX = dragDeltaX
	ds.EbitenSprite.PosY = dragDeltaY
}
