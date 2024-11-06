package game

import (
	"errors"
	"fmt"
)

type canonDamage int64

type Canon struct {
	Damage canonDamage
}

func BuildCanon(damage canonDamage) Canon {
	return Canon{
		damage,
	}
}

func (c1 *Canon) merge(c2 *Canon) {
	c1.Damage += c2.Damage
	c2 = nil
}

type CanonDeck struct {
	Canons        map[BattleGroundColumn]*Canon
	canonCapacity int
}

func (cd *CanonDeck) CanonCapacity() int {
	return cd.canonCapacity
}

func NewCanonDeck(b Battleground) CanonDeck {
	canons := make(map[BattleGroundColumn]*Canon, b.Columns)
	return CanonDeck{
		Canons:        canons,
		canonCapacity: int(b.Columns),
	}
}

func (cd *CanonDeck) DeployCannon(position BattleGroundColumn, canon *Canon) {
	if cd.Canons[position] != nil {
		cd.Canons[position].merge(canon)
	} else {
		cd.Canons[position] = canon
	}
}

// TODO review this function, it's weird, probably needs different API taking into account the FE

func (cd *CanonDeck) MoveCanon(origin, destination BattleGroundColumn) error {
	println(fmt.Sprintf("Moving cannon from %d to %d", origin, destination))
	if origin == destination {
		return nil
	}
	canon := cd.Canons[origin]
	if canon == nil {
		return errors.New("no canon to move here")
	}

	if cd.Canons[destination] == nil {
		cd.Canons[destination] = canon
	} else {
		cd.Canons[destination].merge(canon)
	}
	delete(cd.Canons, origin)
	return nil
}
