package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type MonsterHitState struct {
	context       *EbitenMonster
	hittingBullet *ebiten_canon.EbitenCanonBullet
}

func NewMonsterHitState(
	context *EbitenMonster,
	bullet *ebiten_canon.EbitenCanonBullet,
) *MonsterHitState {
	sprite := ebiten_sprite.NewFromCentralPoint(
		context.sprite.Position().X,
		context.sprite.Position().Y,
		context.animationsSprites.Hit,
		1.1,
		0.1)
	context.sprite = &sprite
	context.lifeLine.SetCurrentLife(int(context.monster.HealthPoints))
	return &MonsterHitState{
		context:       context,
		hittingBullet: bullet,
	}
}

func (m *MonsterHitState) draw(screen *ebiten.Image) {
	m.context.sprite.Draw(screen)
	m.context.lifeLine.Draw(screen)
}

func (m *MonsterHitState) update() {
	if !ebiten_sprite.Collision(m.context.sprite, m.hittingBullet.BulletSprite) {
		if m.context.monster.IsAlive() {
			println("Sapparently monser is still alive here!!!!")
			m.context.setState(
				NewIdleState(m.context, m.context.sprite.Position()))
		} else {
			println("SMONSER IS NOOOOT ALIVE I SHOULD TRANSITION TO DEAD!")
			m.context.setState(NewMonsterDeadState(m.context))
		}
	} else {
		println("COLLISION IS OOOONNNNN BABE")
		rect := m.hittingBullet.BulletSprite.GetRectangle()
		println(fmt.Sprintf("coordinates %d %d %d %d", rect.Min.X, rect.Min.Y, rect.Max.X, rect.Max.Y))
	}

	m.context.sprite.Update()
}

func (m *MonsterHitState) stateName() string {
	return "hit State"
}
