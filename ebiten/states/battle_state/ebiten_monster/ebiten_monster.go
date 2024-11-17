package ebiten_monster

import (
	"canon-tower-defense/ebiten/states/battle_state/ebiten_canon"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
)

type EbitenMonsterState interface {
	setContext(context *EbitenMonster)
	draw(screen *ebiten.Image)
	update()
	updateDeckFiring(bullets []*ebiten_canon.EbitenCanonBullet)
	updateAttack()
	stateName() string
}

type EbitenMonster struct {
	state    EbitenMonsterState
	monster  *game.Monster
	LifeLine *LifeLine
	game     *game.CanonTDGame
}

func NewEbitenMonster(monster *game.Monster, posX, posY float64, game *game.CanonTDGame) *EbitenMonster {
	c := EbitenMonster{
		monster: monster,
		game:    game,
	}
	state := newIdleMonster(posX, posY)

	ll := NewLifeLineFromRectangle(int(monster.MaxLife), state.sprite.GetRectangle())
	c.LifeLine = &ll
	c.transitionTo(&state)
	// If I do not return the pointer, I am returning a copy, which is not the one where "state" is pointing to.
	// so then changing the context, from state.transtitionTo is calling the wrong one, not the one
	// I am returning here, so the one I am returning here cannot be changed of state....
	return &c
}

func (e *EbitenMonster) transitionTo(state EbitenMonsterState) {
	println(fmt.Sprintf("RTransitioning to : %s", state.stateName()))
	e.state = state
	e.state.setContext(e)
}

func (e *EbitenMonster) Draw(screen *ebiten.Image) {
	e.state.draw(screen)
	e.LifeLine.Draw(screen)
}

func (e *EbitenMonster) Update() {
	e.state.update()
	e.LifeLine.Update()
}

func (e *EbitenMonster) UpdateDeckFiring(bullets []*ebiten_canon.EbitenCanonBullet) {
	e.state.updateDeckFiring(bullets)
	e.LifeLine.Update()
}

func (e *EbitenMonster) UpdateAttack() {
	e.state.updateAttack()
}