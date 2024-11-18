package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

const DyingFadeOutSpeed = 0.02

type MonsterDeadState struct {
	sprite  *ebiten_sprite.EbitenAnimatedSprite
	context *EbitenMonster
	fadeOut float32
}

func newDeadFromHitState(hs *MonsterHitState) MonsterDeadState {
	coordinate := hs.sprite.Position()

	beholderDead := ebiten_sprite.NewFromCentralPoint(
		coordinate.X,
		coordinate.Y,
		LoadedImages[BeholderDead],
		64,
		64,
		1.1,
		0.1)

	return MonsterDeadState{
		sprite:  &beholderDead,
		fadeOut: 1,
	}
}

func (m *MonsterDeadState) setContext(context *EbitenMonster) {
	m.context = context
}

func (m *MonsterDeadState) draw(screen *ebiten.Image) {
	m.sprite.DrawWithFade(screen, m.fadeOut)
}

func (m *MonsterDeadState) update() {
	m.fadeOut -= DyingFadeOutSpeed
	if m.fadeOut < 0 {
		m.fadeOut = 0
	}
	m.sprite.Update()
}

func (m *MonsterDeadState) stateName() string {
	return "Dead State"
}
