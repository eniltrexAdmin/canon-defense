package battle_state

import (
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image/color"
)

type screenBlock struct {
	posX   float32
	posY   float32
	width  float32
	height float32
	image  *ebiten.Image
}

func (b screenBlock) inBounds(x, y float32) bool {
	return x >= b.posX && x <= b.posX+b.width && y >= b.posY && y <= b.posY+b.height
}

func (b screenBlock) draw(screen *ebiten.Image) {
	if b.image == nil {
		vector.DrawFilledRect(screen,
			b.posX,
			b.posY,
			b.width,
			b.height,
			color.White,
			false)
	} else {
		op := &ebiten.DrawImageOptions{}
		screen.DrawImage(b.image, op)
	}
}

type highland struct {
	screenBlock
}
