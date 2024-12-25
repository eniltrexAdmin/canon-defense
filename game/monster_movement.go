package game

import "fmt"

type LateralMovement interface {
	nextColumnPlacement(m Monster, bg Battleground) BattleGroundColumn
}

type NoLateralMovement struct{}

func (nlm NoLateralMovement) nextColumnPlacement(m Monster, bg Battleground) BattleGroundColumn {
	return m.CurrentColumn
}

type ZigZagMovement struct {
	direction             string
	limitRight, limitLeft BattleGroundColumn
}

func NewZigZagMovement(
	direction string,
	limitLeft, limitRight BattleGroundColumn,
	bg Battleground,
) *ZigZagMovement {
	// todo enum, but I hate enums in Go.
	// TODO check limit left < limit right < bg.columns.
	if limitLeft < 0 {
		panic(fmt.Sprintf("zig zag movement cant go below 0"))
	}
	if limitRight > bg.Columns {
		panic(fmt.Sprintf("zig zag movement cant go above %d", bg.Rows))
	}
	if limitRight <= limitLeft {
		panic(fmt.Sprintf("zig zag movement invalid setting, from %d, to %d invalid", limitLeft, limitRight))
	}

	return &ZigZagMovement{
		direction:  direction,
		limitRight: limitRight,
		limitLeft:  limitLeft,
	}
}

// TODO we should start adding tests... (creator and this)

func (zzm *ZigZagMovement) nextColumnPlacement(m Monster, bg Battleground) BattleGroundColumn {
	var c BattleGroundColumn
	zzm.fixDirection(m, bg)

	if zzm.direction == "right" {
		c = m.CurrentColumn + 1
	} else {
		c = m.CurrentColumn - 1
	}

	return c
}
func (zzm *ZigZagMovement) fixDirection(m Monster, bg Battleground) {
	if m.CurrentColumn == 0 || m.CurrentColumn == zzm.limitLeft {
		zzm.direction = "right"
	}
	if m.CurrentColumn == bg.Columns || m.CurrentColumn == zzm.limitRight {
		zzm.direction = "left"
	}
}
