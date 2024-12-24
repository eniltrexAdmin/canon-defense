package ebiten_monster

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"canon-tower-defense/game/data"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"image"
	"log"
)

func LoadMonsterImages(m *game.Monster) EbitenMonsterAnimationsSprites {
	switch m.Name {
	case data.BeholderMonster:
		return BeholderImages()
	case data.LizardFolk:
		return LizardFolkImages()
	case data.PurpleWormName:
		return PurpleWorm()
	case data.Dragon:
		return Dragon()
	case data.SerpentFly:
		return SerpentFlyImages()
	case data.Medusa:
		return MedusaImages()
	case data.Slime:
		return SlimeImages()
	case data.Djinn:
		return DjinnImages()
	case data.SlimeBoss:
		return SlimeImages()
	case data.AncientDragon:
		return AncientDragonImages()
	default:
		panic(fmt.Sprintf("No images for monster: %s", m.Name))
	}
}

func BeholderImages() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.Beholder),
			64, 64,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.BeholderHit),
			64, 64,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.BeholderDie),
			64, 64,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.BeholderAttack),
			128, 128,
		),
	}
}

func LizardFolkImages() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.LizardFolk),
			48, 72,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.LizardFolkHit),
			48, 72,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.LizardFolkDie),
			48, 72,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.LizardFolkAttack),
			144, 168,
		),
	}
}

func MedusaImages() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.Medusa),
			64, 64,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.MedusaHit),
			64, 64,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.MedusaDie),
			64, 64,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.MedusaAttack),
			128, 128,
		),
	}
}

func SerpentFlyImages() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.SerpentFly),
			32, 32,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.SerpentFlyHit),
			32, 32,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.SerpentFlyDie),
			32, 32,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.SerpentFlyAttack),
			96, 96,
		),
	}
}

func SlimeImages() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.Slime),
			32, 32,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.SlimeHit),
			32, 32,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.SlimeDie),
			32, 32,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.SlimeAttack),
			96, 96,
		),
	}
}

func PurpleWorm() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.PurpleWorm),
			64, 64,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.PurpleWormHit),
			64, 64,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.PurpleWormDie),
			64, 64,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.PurpleWormAttack),
			128, 128,
		),
	}
}

func Dragon() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.Dragon),
			64, 64,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.DragonHit),
			64, 64,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.DragonDie),
			64, 64,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.DragonAttack),
			224, 224,
		),
	}
}

func DjinnImages() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.Djinn),
			64, 64,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.DjinnHit),
			64, 64,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.DjinnDie),
			64, 64,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.DjinnAttack),
			128, 64,
		),
	}
}

func AncientDragonImages() EbitenMonsterAnimationsSprites {
	return EbitenMonsterAnimationsSprites{
		Idle: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.AncientDragon),
			96, 96,
		),
		Hit: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.AncientDragonHit),
			96, 96,
		),
		Dead: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.AncientDragonDie),
			96, 96,
		),
		Attack: ebiten_sprite.NewAnimatedSprite(
			loadImage(assets.AncientDragonAttack),
			288, 288,
		),
	}
}

func loadImage(data []byte) *ebiten.Image {
	img, _, err := image.Decode(bytes.NewReader(data))
	if err != nil {
		log.Fatal(err)
	}
	return ebiten.NewImageFromImage(img)
}
