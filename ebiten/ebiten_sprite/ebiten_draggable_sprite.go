package ebiten_sprite

type EbitenDraggableSprite struct {
	initialPlacementX, initialPlacementY float64
	dragIniX, dragIniY                   float64
	IsDragged                            bool
	CurrentStroke                        Stroke
	EbitenSprite
}

type Stroke interface {
	CurrentPosition() (x, y int)
	DragDelta() (x, y float64)
	IsJustReleased() bool
}

func NewFromSprite(sprite EbitenSprite) *EbitenDraggableSprite {
	return &EbitenDraggableSprite{
		initialPlacementX: sprite.PosX,
		initialPlacementY: sprite.PosY,
		dragIniX:          0,
		dragIniY:          0,
		IsDragged:         false,
		EbitenSprite:      sprite,
		CurrentStroke:     nil,
	}
}

func (ds *EbitenDraggableSprite) StrokeStart(st Stroke) {
	x, y := st.CurrentPosition()
	if ds.EbitenSprite.InBounds(x, y) {
		ds.CurrentStroke = st
		ds.IsDragged = true
		return
	}
}

func (ds *EbitenDraggableSprite) ReleaseDrag() {
	ds.EbitenSprite.PosX = ds.initialPlacementX
	ds.EbitenSprite.PosY = ds.initialPlacementY
	ds.IsDragged = false
	ds.CurrentStroke = nil
}

func (ds *EbitenDraggableSprite) Update() {
	if ds.IsDragged == false {
		return
	}

	dragDeltaX, dragDeltaY := ds.CurrentStroke.DragDelta()

	ds.EbitenSprite.PosX = ds.initialPlacementX + dragDeltaX
	ds.EbitenSprite.PosY = ds.initialPlacementY + dragDeltaY
}
