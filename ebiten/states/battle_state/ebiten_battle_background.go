package battle_state

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

const TileSize float64 = 50

const PlatformYPadding float64 = 20

const BattleGroundHeight float64 = 500

type ebitenBattleGround struct {
	tiles       []ebiten_sprite.EbitenSprite
	rowDistance float64
}

func (ecd *ebitenBattleGround) draw(screen *ebiten.Image) {
	for _, tile := range ecd.tiles {
		tile.Draw(screen)
	}
}

func newEbitenBattleGround(bg game.Battleground) ebitenBattleGround {
	availableWidth := constants.ScreenWidth / int(bg.Columns)
	availableHeight := int(BattleGroundHeight) / int(bg.VisibleRows)

	img, _, err := image.Decode(bytes.NewReader(assets.Highland))
	if err != nil {
		log.Fatal(err)
	}
	platformImage := ebiten.NewImageFromImage(img)

	var tiles []ebiten_sprite.EbitenSprite

	for i := 0; i < int(bg.VisibleRows); i++ {
		for j := 0; j < int(bg.Columns); j++ {

			centerXSpace := availableWidth / 2
			tileCenterPointX := float64(centerXSpace)

			centerYSpace := availableHeight / 2
			tileSCenterPointY := float64(centerYSpace)

			posX := float64(availableWidth*j) + tileCenterPointX
			posY := float64(availableHeight*i) + tileSCenterPointY

			tile := ebiten_sprite.NewFromCentralPointScaleImage(
				posX,
				posY+PlatformYPadding,
				platformImage,
				TileSize,
			)
			tiles = append(tiles, tile)
		}
	}
	return ebitenBattleGround{
		tiles:       tiles,
		rowDistance: float64(availableHeight),
	}
}
