package ebiten_monster

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"image"
	"image/color"
)

type LifeLine struct {
	maxLife          int
	currentLife      int
	destinationWidth float32
	lifeMarker       float32
	width, height    float32
	posX, posY       float32
	animationSpeed   float32
}

func NewLifeLineFromRectangle(maxLife int, rectangle image.Rectangle) LifeLine {
	posX := float32(rectangle.Min.X) + 5
	posY := rectangle.Max.Y

	startingWidth := float32(rectangle.Dx()) - 10

	return LifeLine{
		maxLife:          maxLife,
		currentLife:      maxLife,
		destinationWidth: startingWidth,
		lifeMarker:       startingWidth,
		width:            startingWidth,
		height:           5, // to say something
		posX:             posX,
		posY:             float32(posY),
		animationSpeed:   1,
	}
}

func (l *LifeLine) SetCurrentLife(cl int) {
	l.currentLife = cl
	l.destinationWidth = float32(l.currentLife) / float32(l.maxLife) * l.width
	println(fmt.Sprintf("Operation is: %d / %d * %f", l.currentLife, l.maxLife, l.width))
	//
	println(fmt.Sprintf("Setting destination width to be: %f", l.destinationWidth))
}

func (l *LifeLine) Update() {
	if l.lifeMarker < l.destinationWidth {
		l.lifeMarker += l.animationSpeed
		if l.lifeMarker > l.destinationWidth {
			l.lifeMarker = l.destinationWidth
		}
	} else if l.lifeMarker > l.destinationWidth {
		l.lifeMarker -= l.animationSpeed
		if l.lifeMarker < l.destinationWidth {
			l.lifeMarker = l.destinationWidth
		}
	}
}

func (l *LifeLine) Draw(screen *ebiten.Image) {
	vector.DrawFilledRect(screen,
		l.posX,
		l.posY,
		l.lifeMarker,
		l.height,
		color.RGBA{R: 255, G: 0, B: 0, A: 255},
		false)

	vector.StrokeRect(screen,
		l.posX-1,
		l.posY-1,
		l.width+2,
		l.height+2,
		1,
		color.RGBA{R: 255, G: 255, B: 255, A: 255}, // White color for border
		false,
	)
}
