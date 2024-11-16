package game

const defaultBattleGroundSize BattleGroundColumn = 5

type BattleGroundColumn int8
type BattleGroundRow uint64

type Battleground struct {
	Columns     BattleGroundColumn
	Rows        BattleGroundRow
	VisibleRows BattleGroundRow
}

func NewBattleGround(columns, rows, visibleRows int) Battleground {
	if visibleRows > rows {
		panic("visibleRows > rows")
	}
	if columns < 1 || rows < 1 || visibleRows < 1 {
		panic("battleground needs to exist in all dimensions!")
	}
	return Battleground{
		Columns:     BattleGroundColumn(columns),
		Rows:        BattleGroundRow(rows),
		VisibleRows: BattleGroundRow(visibleRows),
	}
}

func (bg Battleground) checkIndexPosition(row, column int) {
	if row >= int(bg.Rows) || row < 0 {
		panic("row index out of bounds")
	}
	if column >= int(bg.Columns) || column < 0 {
		panic("column index out of bounds")
	}
}
