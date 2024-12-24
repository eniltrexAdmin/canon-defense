package data

import "canon-tower-defense/game"

const BeholderMonster string = "beholder"
const LizardFolk string = "lizard"
const Medusa string = "medusa"
const Slime string = "slime"
const SerpentFly string = "serpent"
const PurpleWormName string = "worm"
const Dragon string = "dragon"
const Djinn string = "djinn"
const DeathKnight string = "death_knight"
const FrostGiant string = "frost_giant"
const Oni string = "oni"
const OwlBear string = "owl_bear"
const Samurai string = "samurai"
const Tengu string = "tengu"
const SlimeBoss string = "slimeBoss"
const AncientDragon string = "ancientDragon"

func SlimeTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Slime,
		HealthPoints:   game.CanonDamage(1),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func SerpentFlyTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           SerpentFly,
		HealthPoints:   game.CanonDamage(2),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func BeholderTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           BeholderMonster,
		HealthPoints:   game.CanonDamage(2),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func LizardFolkTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           LizardFolk,
		HealthPoints:   game.CanonDamage(2),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func MedusaTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Medusa,
		HealthPoints:   game.CanonDamage(6),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func PurpleWormTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           PurpleWormName,
		HealthPoints:   game.CanonDamage(5),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func DragonTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Dragon,
		HealthPoints:   game.CanonDamage(15),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func DjinnTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Djinn,
		HealthPoints:   game.CanonDamage(20),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func DeathKnightTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           DeathKnight,
		HealthPoints:   game.CanonDamage(20),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func FrostGianTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           FrostGiant,
		HealthPoints:   game.CanonDamage(20),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func OniTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Oni,
		HealthPoints:   game.CanonDamage(20),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func OwlBearTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           OwlBear,
		HealthPoints:   game.CanonDamage(20),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func SamuraiTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Samurai,
		HealthPoints:   game.CanonDamage(20),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func TenguTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Tengu,
		HealthPoints:   game.CanonDamage(20),
		ColumnMovement: 1,
		RowMovement:    1,
	}
}

func SlimeBossTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           SlimeBoss,
		HealthPoints:   game.CanonDamage(100),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func AncientDragonTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           AncientDragon,
		HealthPoints:   game.CanonDamage(50),
		ColumnMovement: 0,
		RowMovement:    2,
	}
}
