package game

import "errors"

type canonDamage int64

type Canon struct {
	Damage canonDamage
}

func (c1 *Canon) merge(c2 *Canon) {
	c1.Damage += c2.Damage
	c2 = nil
}

type canonDeck struct {
	canons []*Canon
}

func NewCanonDeck(b battleground) canonDeck {
	canons := make([]*Canon, b.columns)
	return canonDeck{canons: canons}
}

func (cd *canonDeck) placeCanon(position BattleGroundColumn, canon *Canon) error {
	cd.canons[position] = canon
	return nil
}

func (cd *canonDeck) moveCanon(origin, destination BattleGroundColumn) error {
	canon := cd.canons[origin]
	if canon == nil {
		return errors.New("no canon to move here")
	}

	if cd.canons[destination] == nil {
		cd.canons[destination] = canon
	} else {
		cd.canons[destination].merge(canon)
	}
	return nil
}
