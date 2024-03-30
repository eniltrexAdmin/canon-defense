package game

const defaultBattleGroundSize BattleGroundColumn = 5

type BattleGroundColumn int8
type BattleGroundLength uint64

type battleground struct {
	columns  BattleGroundColumn
	rows     BattleGroundLength
	monsters [][]*Monster
}

type levelGenerator interface {
	Generate() battleground
}

func (bg battleground) fire(deck canonDeck) {

}
