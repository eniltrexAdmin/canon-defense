package battle_state

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type deployArea struct {
	PosX        float64
	PosY        float64
	width       float64
	height      float64
	color       color.Color
	strokeWidth float32
}

func NewDeployAreaFromCentralPoint(posX, posY, width, height float64, color color.Color) deployArea {
	return deployArea{
		PosX:        posX - width/2,
		PosY:        posY - height/2,
		width:       width,
		height:      height,
		color:       color,
		strokeWidth: 0,
	}
}

func (d *deployArea) GetRectangle() image.Rectangle {
	x := int(d.PosX)
	y := int(d.PosY)
	width := int(d.width)
	height := int(d.height)
	return image.Rect(x, y, x+width, y+height)
}

func (d *deployArea) update(draggedSprite *ebiten_sprite.EbitenSprite) {
	if draggedSprite == nil {
		d.strokeWidth = 0
		return
	}
	d.strokeWidth = 2
	if ebiten_sprite.Collision(d, draggedSprite) {
		d.strokeWidth = 4
	}
}

func (d *deployArea) draw(screen *ebiten.Image) {
	vector.StrokeRect(screen,
		float32(d.PosX),
		float32(d.PosY),
		float32(d.width),
		float32(d.height),
		d.strokeWidth,
		d.color,
		false,
	)
}
