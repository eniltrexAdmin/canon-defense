package assets

import (
	_ "embed"
)

//go:embed logo.png
var LogoPng []byte

//go:embed canon_deck/regular_canon.png
var RegularCanon []byte
