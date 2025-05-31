package level_selection

import (
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/audio"
	"github.com/hajimehoshi/ebiten/v2/text"
	"github.com/hajimehoshi/ebiten/v2/vector"
	"golang.org/x/image/font/basicfont"
	"image/color"
)

const levelHeight float64 = 50
const levelHeightPadding float64 = 20

type Level struct {
	LevelNumber int
	PosX        float64
	PosY        float64
	width       float64
	height      float64
	color       color.Color
	strokeWidth float32
	enabled     bool
	hover       bool
	completed   bool
	hoverSound  *audio.Player
}

func NewLevelSet(completedLevels []bool, hoverSound *audio.Player) []*Level {
	var levels []*Level
	for levelNumber, completed := range completedLevels {
		levels = append(levels, NewEbitenLevel(levelNumber+1, completed, hoverSound))
	}
	return levels
}

func NewEbitenLevel(levelNumber int, completed bool, hoverSound *audio.Player) *Level {
	c := completedColor
	if !completed {
		c = enabledColor
	}
	column := 1
	if levelNumber > 10 {
		column = 2
	}
	var posX float64 = 20
	posY := 10 + float64(levelNumber-1)*(levelHeight+levelHeightPadding)
	if column > 1 {
		posX = 240
		posY = 10 + float64(levelNumber-11)*(levelHeight+levelHeightPadding)
	}

	return &Level{
		LevelNumber: levelNumber,
		PosX:        posX,
		PosY:        posY,
		width:       200,
		height:      levelHeight,
		color:       c,
		strokeWidth: 3,
		enabled:     true, // leaving enabled maybe for the future...yeah I know yagni.
		hover:       false,
		completed:   completed,
		hoverSound:  hoverSound,
	}
}

var enabledColor = color.RGBA{
	R: uint8(255),
	G: uint8(255),
	B: uint8(255),
	A: 0,
}

var disabledColor = color.RGBA{
	R: uint8(100),
	G: uint8(100),
	B: uint8(100),
	A: 0,
}

var completedColor = color.RGBA{
	R: uint8(0),
	G: uint8(255),
	B: uint8(0),
	A: 0,
}

func (l *Level) InBounds(xInt, yInt int) bool {
	x := float64(xInt)
	y := float64(yInt)
	return x >= l.PosX && x <= l.PosX+l.width && y >= l.PosY && y <= l.PosY+l.height
}

func (l *Level) Update(x, y int) {
	if !l.enabled {
		return
	}

	if l.InBounds(x, y) && !l.hover {
		l.hoverSound.Rewind()
		l.hoverSound.Play()
	}

	if l.InBounds(x, y) {
		l.hover = true
		l.strokeWidth = 15
	} else {
		l.hover = false
		l.strokeWidth = 3
	}
}

func (l *Level) Draw(screen *ebiten.Image) {
	vector.StrokeRect(screen,
		float32(l.PosX),
		float32(l.PosY),
		float32(l.width),
		float32(l.height),
		l.strokeWidth,
		l.color,
		false,
	)
	l.DrawText(screen)
}

func (l *Level) DrawText(screen *ebiten.Image) {
	basicFont := basicfont.Face7x13

	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(1.5, 1.5)
	op.GeoM.Translate(l.PosX+60, l.PosY+(levelHeight/2+10))

	text.DrawWithOptions(screen,
		fmt.Sprintf("LEVEL %d", l.LevelNumber),
		basicFont,
		op,
	)
}
