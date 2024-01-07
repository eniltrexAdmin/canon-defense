package internal

type battlegroundMovement uint8

type Monster struct {
	name           string
	healthPoints   canonDamage
	rowMovement    battlegroundMovement
	columnMovement battlegroundMovement
}

func (m Monster) move(b battleground) {

}

func (m Monster) hit(damage canonDamage) {
	m.healthPoints -= damage
}

func (m Monster) alive() bool {
	return m.healthPoints > 0
}

func Skeleton() Monster {
	return Monster{
		name:           "Skeleton",
		healthPoints:   1,
		columnMovement: 0,
		rowMovement:    1,
	}
}
