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

type ebitenBattleGround struct {
	tiles           []ebiten_sprite.EbitenSprite
	visibleMonsters []*ebitenMonster
}

func (ecd ebitenBattleGround) draw(screen *ebiten.Image) {
	for _, tile := range ecd.tiles {
		tile.Draw(screen)
	}
	for _, visibleMonsters := range ecd.visibleMonsters {
		visibleMonsters.draw(screen)
	}
}

func (ecd ebitenBattleGround) update(bullets []*ebitenCanonBullet) {
	for _, visibleMonsters := range ecd.visibleMonsters {
		visibleMonsters.update(bullets)
	}
}

func newEbitenBattleGround(bg game.Battleground) ebitenBattleGround {
	availableWidth := constants.ScreenWidth / int(bg.Columns)
	availableHeight := int(canonYPlacement) / int(bg.VisibleRows)

	img, _, err := image.Decode(bytes.NewReader(assets.Highland))
	if err != nil {
		log.Fatal(err)
	}
	platformImage := ebiten.NewImageFromImage(img)

	var tiles []ebiten_sprite.EbitenSprite

	var visibleMonsters []*ebitenMonster

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

			possibleMonster := monsterFromGameInEbitenBattleGround(i, j, bg)

			if possibleMonster != nil {
				monster := NewEbitenMonster(possibleMonster, posX, posY)
				visibleMonsters = append(visibleMonsters, &monster)
			}
		}
	}
	return ebitenBattleGround{
		tiles:           tiles,
		visibleMonsters: visibleMonsters,
	}
}

func monsterFromGameInEbitenBattleGround(ebitenRow, ebitenColumn int, bg game.Battleground) *game.Monster {
	gameI, gameJ := translateEbitenToGame(ebitenRow, ebitenColumn, int(bg.VisibleRows))
	return bg.Monsters[gameI][gameJ]
}

func translateEbitenToGame(ebitenRow, ebitenColumn int, gameVisibleRows int) (int, int) {

	return gameVisibleRows - (ebitenRow + 1), ebitenColumn
}
