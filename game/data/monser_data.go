package data

import "canon-tower-defense/game"

const BeholderMonster string = "beholder"
const LizardFolk string = "lizard"
const Medusa string = "medusa"
const SerpentFly string = "serpent"
const PurpleWormName string = "worm"
const Dragon string = "dragon"

func Beholder() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           BeholderMonster,
		HealthPoints:   game.CanonDamage(1),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func PurpleWorm() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           PurpleWormName,
		HealthPoints:   game.CanonDamage(3),
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

func DragonTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           Dragon,
		HealthPoints:   game.CanonDamage(30),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}

func SerpentFlyTemplate() game.MonsterTemplate {
	return game.MonsterTemplate{
		Name:           SerpentFly,
		HealthPoints:   game.CanonDamage(3),
		ColumnMovement: 0,
		RowMovement:    1,
	}
}
