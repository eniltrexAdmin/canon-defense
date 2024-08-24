package game

import "errors"

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
	Canons []*Canon
}

func NewCanonDeck(b Battleground) CanonDeck {
	canons := make([]*Canon, b.Columns)
	return CanonDeck{Canons: canons}
}

func (cd *CanonDeck) placeCanon(position BattleGroundColumn, canon *Canon) error {
	// TODO what happens if position is already filled? errors!!!
	cd.Canons[position] = canon
	return nil
}

func (cd *CanonDeck) moveCanon(origin, destination BattleGroundColumn) error {
	canon := cd.Canons[origin]
	if canon == nil {
		return errors.New("no canon to move here")
	}

	if cd.Canons[destination] == nil {
		cd.Canons[destination] = canon
	} else {
		cd.Canons[destination].merge(canon)
	}
	return nil
}
