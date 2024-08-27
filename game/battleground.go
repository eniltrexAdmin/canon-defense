package game

const defaultBattleGroundSize BattleGroundColumn = 5

type BattleGroundColumn int8
type BattleGroundLength uint64
type BattleGroundVisibleRows int8

type Battleground struct {
	Columns     BattleGroundColumn
	Rows        BattleGroundLength
	VisibleRows BattleGroundVisibleRows
	Monsters    [][]*Monster
}

type LevelGenerator interface {
	Generate(level int) Battleground
}

func ToBattleGroundColumn(input int, bg Battleground) BattleGroundColumn {
	// TODO encapsulate panics in probably all in game, and return errors here instead.
	if input < 0 || input > int(bg.Columns) {
		panic("value cannot be converted to column")
	}
	return BattleGroundColumn(input)
}
