package game

const defaultBattleGroundSize BattleGroundColumn = 5

type BattleGroundColumn int8
type BattleGroundLength uint64

type Battleground struct {
	columns  BattleGroundColumn
	rows     BattleGroundLength
	monsters [][]*Monster
}

type LevelGenerator interface {
	Generate(level int) Battleground
}

func (bg Battleground) fire(deck canonDeck) {

}
