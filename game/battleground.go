package game

const defaultBattleGroundSize BattleGroundColumn = 5

type BattleGroundColumn int8
type BattleGroundRow int64

const NoVisibleRow BattleGroundRow = -1

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

func (bg Battleground) checkIndexPosition(row BattleGroundRow, column BattleGroundColumn) {
	if row >= bg.Rows || row < 0 {
		panic("Row index out of bounds")
	}
	if column >= bg.Columns || column < 0 {
		panic("Column index out of bounds")
	}
}

func (bg Battleground) toVisibleRow(gameRow BattleGroundRow) BattleGroundRow {
	visibleRow := bg.VisibleRows - gameRow - 1
	return visibleRow
}
