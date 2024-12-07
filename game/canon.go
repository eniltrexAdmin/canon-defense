package game

type CanonDamage int64

// TODO this one needs another iteration, like monster plus monster in field, and canon deck might have less logic.

type Canon struct {
	Damage CanonDamage
}

func BuildCanon(damage CanonDamage) Canon {
	return Canon{
		damage,
	}
}

func (c1 *Canon) merge(c2 *Canon) {
	c1.Damage += c2.Damage
	c2 = nil
}
