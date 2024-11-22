package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

const DyingFadeOutSpeed = 0.02

type MonsterDeadState struct {
	context *EbitenMonster
	fadeOut float32
}

func NewMonsterDeadState(context *EbitenMonster) *MonsterDeadState {
	sprite := ebiten_sprite.NewFromCentralPoint(
		context.sprite.Position().X,
		context.sprite.Position().Y,
		LoadedImages[BeholderDead],
		64,
		64,
		1.1,
		0.1)
	context.sprite = &sprite
	return &MonsterDeadState{
		context: context,
		fadeOut: 1,
	}
}

func (m *MonsterDeadState) draw(screen *ebiten.Image) {
	m.context.sprite.DrawWithFade(screen, m.fadeOut)
}

func (m *MonsterDeadState) update() {
	m.fadeOut -= DyingFadeOutSpeed
	if m.fadeOut < 0 {
		m.fadeOut = 0
	}
	m.context.sprite.Update()
}

func (m *MonsterDeadState) stateName() string {
	return "Dead State"
}
