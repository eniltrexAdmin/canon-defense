package ebiten_monster

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/game"
	"canon-tower-defense/game/data"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

func LoadMonsterImages(m *game.Monster) EbitenMonsterImages {
	switch m.Name {
	case data.BeholderMonster:
		return BeholderImages()
	case data.LizardFolk:
		return LizardFolkImages()
	case data.PurpleWormName:
		return PurpleWorm()
	case data.Dragon:
		return Dragon()

	default:
		panic(fmt.Sprintf("No images for monster: %s", m.Name))
	}
}

func BeholderImages() EbitenMonsterImages {
	return EbitenMonsterImages{
		Idle:   loadImage(assets.Beholder),
		Hit:    loadImage(assets.BeholderHit),
		Dead:   loadImage(assets.BeholderDie),
		Attack: loadImage(assets.BeholderAttack),
	}
}

func LizardFolkImages() EbitenMonsterImages {
	return EbitenMonsterImages{
		Idle:   loadImage(assets.LizardFolk),
		Hit:    loadImage(assets.LizardFolkHit),
		Dead:   loadImage(assets.LizardFolkDie),
		Attack: loadImage(assets.LizardFolkAttack),
	}
}

func PurpleWorm() EbitenMonsterImages {
	return EbitenMonsterImages{
		Idle:   loadImage(assets.PurpleWorm),
		Hit:    loadImage(assets.PurpleWormHit),
		Dead:   loadImage(assets.PurpleWormDie),
		Attack: loadImage(assets.PurpleWormAttack),
	}
}

func Dragon() EbitenMonsterImages {
	return EbitenMonsterImages{
		Idle:   loadImage(assets.Dragon),
		Hit:    loadImage(assets.DragonHit),
		Dead:   loadImage(assets.DragonDie),
		Attack: loadImage(assets.DragonAttack),
	}
}

func loadImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
