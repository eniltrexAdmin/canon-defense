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

//go:embed "monsters/Beholder_idle.png"
var Beholder []byte

//go:embed "monsters/Beholder_hit.png"
var BeholderHit []byte

//go:embed "monsters/Beholder_die.png"
var BeholderDie []byte
