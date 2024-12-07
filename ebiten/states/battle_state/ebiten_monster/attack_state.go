package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"github.com/hajimehoshi/ebiten/v2"
)

const AttackStateName = "Attacking!"
const MovementSpeed float64 = 2

type MonsterAttackState struct {
	context     *EbitenMonster
	destination ebiten_sprite.ScreenCoordinate
}

func NewMonsterAttackState(context *EbitenMonster, stepHeight float64) *MonsterAttackState {
	position := context.sprite.Position()
	sprite := ebiten_sprite.NewFromCentralPoint(
		position.X,
		position.Y,
		context.animations.Attack,
		128,
		128,
		1,
		0.1)
	context.sprite = &sprite
	stepSize := stepHeight * float64(context.monster.RowMovement)

	destination := ebiten_sprite.ScreenCoordinate{
		X: position.X,
		Y: stepSize + position.Y,
	}

	return &MonsterAttackState{
		context:     context,
		destination: destination,
	}
}

func (m *MonsterAttackState) draw(screen *ebiten.Image) {
	m.context.sprite.Draw(screen)
}

func (m *MonsterAttackState) update() {
	m.context.sprite.Move(m.destination, MovementSpeed)
	m.context.sprite.Update()
	if m.context.sprite.Position() == m.destination {
		m.context.setState(NewIdleState(m.context, m.context.sprite.Position()))
	}
}

func (m *MonsterAttackState) stateName() string {
	return AttackStateName
}
