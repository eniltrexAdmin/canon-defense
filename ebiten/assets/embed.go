package assets

import (
	_ "embed"
)

//go:embed logo.png
var LogoPng []byte

//go:embed canon_deck/regular_canon.png
var RegularCanon []byte

//go:embed canon_deck/canon_placement.png
var CanonPedestal []byte

//go:embed canon_deck/Rocket_108.png
var Bullet []byte

//go:embed "background/moving_platform_H.png"
var Highland []byte

// Monsters:

//go:embed "monsters/Beholder_idle.png"
var Beholder []byte

//go:embed "monsters/Beholder_hit.png"
var BeholderHit []byte

//go:embed "monsters/Beholder_die.png"
var BeholderDie []byte

//go:embed "monsters/Beholder_attack_green.png"
var BeholderAttack []byte

//go:embed "monsters/Lizardfolk_idle.png"
var LizardFolk []byte

//go:embed "monsters/Lizardfolk_hit.png"
var LizardFolkHit []byte

//go:embed "monsters/Lizardfolk_die.png"
var LizardFolkDie []byte

//go:embed "monsters/Lizardfolk_attack.png"
var LizardFolkAttack []byte

//go:embed "monsters/PurpleWorm_idle.png"
var PurpleWorm []byte

//go:embed "monsters/PurpleWorm_hit.png"
var PurpleWormHit []byte

//go:embed "monsters/PurpleWorm_die.png"
var PurpleWormDie []byte

//go:embed "monsters/PurpleWorm_move.png"
var PurpleWormMove []byte

//go:embed "monsters/PurpleWorm_attack.png"
var PurpleWormAttack []byte

//go:embed "monsters/AdultRedDragon_idle.png"
var Dragon []byte

//go:embed "monsters/AdultRedDragon_hit.png"
var DragonHit []byte

//go:embed "monsters/AdultRedDragon_die.png"
var DragonDie []byte

//go:embed "monsters/AdultRedDragon_fly.png"
var DragonAttack []byte
