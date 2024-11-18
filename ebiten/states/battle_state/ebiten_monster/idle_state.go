package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
)

type MonsterIdleState struct {
	sprite  *ebiten_sprite.EbitenAnimatedSprite
	context *EbitenMonster
}

func newIdleMonster(coordinate ebiten_sprite.ScreenCoordinate) MonsterIdleState {
	beholder := ebiten_sprite.NewFromCentralPoint(
		coordinate.X,
		coordinate.Y,
		LoadedImages[BeholderIdle],
		64,
		64,
		1,
		0.1)
	return MonsterIdleState{
		sprite: &beholder,
	}
}

func (m *MonsterIdleState) setContext(context *EbitenMonster) {
	m.context = context
	m.context.LifeLine.Show()
}

func (m *MonsterIdleState) draw(screen *ebiten.Image) {
	m.sprite.Draw(screen)
}

func (m *MonsterIdleState) update() {
	m.sprite.Update()
	for _, bullet := range m.context.bulletsInField {
		if ebiten_sprite.Collision(m.sprite, bullet.BulletSprite) {
			hitState := newMonsterHitState(m, bullet)
			m.context.transitionTo(hitState)
		}
	}
}

func (m *MonsterIdleState) stateName() string {
	return "Idle State"
}

func (m *MonsterIdleState) Coordinates() ebiten_sprite.ScreenCoordinate {
	return m.sprite.Position()
}

func (m *MonsterIdleState) GetRectangle() image.Rectangle {
	return m.sprite.GetRectangle()
}
