package ebiten_monster

import (
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type EbitenMonsterHitEvent struct {
	Monster *game.Monster
	Canon   *game.Canon
}

type EbitenMonsterAnimationsSprites struct {
	Idle   ebiten_sprite.AnimatedSprite
	Hit    ebiten_sprite.AnimatedSprite
	Dead   ebiten_sprite.AnimatedSprite
	Attack ebiten_sprite.AnimatedSprite
}

type EbitenMonsterState interface {
	draw(screen *ebiten.Image)
	update()
	stateName() string
}

type EbitenMonster struct {
	state             EbitenMonsterState
	monster           *game.Monster
	sprite            *ebiten_sprite.EbitenAnimatedSprite
	animationsSprites EbitenMonsterAnimationsSprites
	lifeLine          *LifeLine
	bulletsInField    []*ebiten_canon.EbitenCanonBullet // not convinced if it should be here or only idle state
	hitTrigger        func(event EbitenMonsterHitEvent)
}

func NewEbitenMonster(monster *game.Monster, posX, posY float64, HitTrigger func(event EbitenMonsterHitEvent)) *EbitenMonster {

	coordinate := ebiten_sprite.ScreenCoordinate{X: posX, Y: posY}

	c := &EbitenMonster{
		monster:           monster,
		hitTrigger:        HitTrigger,
		animationsSprites: LoadMonsterImages(monster),
	}

	c.state = NewIdleState(c, coordinate)

	ll := NewLifeLineFromRectangle(int(monster.MaxLife), c.sprite.GetRectangle())
	c.lifeLine = &ll

	return c
}

func (e *EbitenMonster) setState(state EbitenMonsterState) {
	println(fmt.Sprintf("RTransitioning to : %s", state.stateName()))
	e.state = state
}

func (e *EbitenMonster) Draw(screen *ebiten.Image) {
	e.state.draw(screen)
}

func (e *EbitenMonster) Update() {
	e.state.update()
	e.lifeLine.Update(e.sprite.GetRectangle())
}

func (e *EbitenMonster) DeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	e.bulletsInField = bullets
}

func (e *EbitenMonster) Attack(stepHeight float64) {
	if !e.monster.IsAlive() {
		return
	}
	e.setState(NewMonsterAttackState(e, stepHeight))
}

func (e *EbitenMonster) IsAttacking() bool {
	return e.state.stateName() == AttackStateName
}

func (e *EbitenMonster) ReachedGameOver() bool {
	if !e.monster.IsAlive() {
		return false
	}
	println(fmt.Sprintf("row is : %d", e.monster.CurrentRow))
	return e.monster.CurrentRow < 0
}
