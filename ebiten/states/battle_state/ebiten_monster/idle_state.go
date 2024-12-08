package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterIdleState struct {
	context *EbitenMonster
}

func NewIdleState(context *EbitenMonster, coordinate ebiten_sprite.ScreenCoordinate) *MonsterIdleState {
	sprite := ebiten_sprite.NewFromCentralPoint(
		coordinate.X,
		coordinate.Y,
		context.animationsSprites.Idle,
		1,
		0.1)
	context.sprite = &sprite
	return &MonsterIdleState{context: context}
}

func (m *MonsterIdleState) draw(screen *ebiten.Image) {
	m.context.sprite.Draw(screen)
}

func (m *MonsterIdleState) update() {
	m.context.sprite.Update()
	for _, bullet := range m.context.bulletsInField {
		if ebiten_sprite.Collision(m.context.sprite, bullet.BulletSprite) {

			m.context.hitTrigger(EbitenMonsterHitEvent{
				Canon:   &bullet.Canon.Canon,
				Monster: m.context.monster,
			})

			m.context.setState(NewMonsterHitState(m.context, bullet))
		}
	}
}

func (m *MonsterIdleState) stateName() string {
	return "Idle State"
}
