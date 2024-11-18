package ebiten_canon

import (
	"bytes"
	"canon-tower-defense/ebiten/assets"
	"canon-tower-defense/ebiten/constants"
	"canon-tower-defense/ebiten/ebiten_sprite"
	"canon-tower-defense/game"
	"fmt"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"image"
	"log"
)

const canonYPlacement float64 = 550

type EbitenCanonDeck struct {
	ebitenCanons  map[int]*ebitenCanon
	deployAreas   map[int]*deployArea
	actionButton  ebitenActionButton
	game          *game.CanonTDGame
	Firing        bool
	draggedSprite *ebiten_sprite.EbitenDraggableSprite
}

func NewEbitenCanonDeck(g game.CanonTDGame) EbitenCanonDeck {
	img, _, err := image.Decode(bytes.NewReader(assets.RegularCanon))
	if err != nil {
		log.Fatal(err)
	}
	canonImage := ebiten.NewImageFromImage(img)

	img2, _, err := image.Decode(bytes.NewReader(assets.CanonPedestal))
	if err != nil {
		log.Fatal(err)
	}
	pedestalImage := ebiten.NewImageFromImage(img2)

	availableWidth := float64(constants.ScreenWidth / g.CanonDeck.CanonCapacity())

	cs := make(map[int]*ebitenCanon, g.CanonDeck.CanonCapacity())
	das := make(map[int]*deployArea, g.CanonDeck.CanonCapacity())

	color := ebiten_sprite.RandomColor()
	for formationPlacement := 0; formationPlacement < g.CanonDeck.CanonCapacity(); formationPlacement++ {
		centerX := getCanonCenterX(formationPlacement, g.CanonDeck.CanonCapacity())
		canon := g.CanonDeck.Canons[game.BattleGroundColumn(formationPlacement)]
		if canon != nil {
			ec := newEbitenCanon(
				*canon,
				formationPlacement,
				canonImage,
				centerX,
				canonYPlacement,
			)
			cs[formationPlacement] = &ec
		}

		da := NewDeployAreaFromCentralPoint(
			centerX, canonYPlacement,
			availableWidth-20, 60,
			color,
		)
		das[formationPlacement] = &da
	}

	ab := newEbitenActionButton(canonImage, pedestalImage, constants.ScreenWidth)

	return EbitenCanonDeck{
		ebitenCanons: cs,
		deployAreas:  das,
		actionButton: ab,
		game:         &g,
		Firing:       false,
	}
}

func getCanonCenterX(formationPlacement, numberCanons int) float64 {
	availableWidth := float64(constants.ScreenWidth / numberCanons)
	return availableWidth*float64(formationPlacement) + availableWidth/2
}

func (ecd *EbitenCanonDeck) Draw(screen *ebiten.Image) {
	for _, canon := range ecd.ebitenCanons {
		canon.draw(screen)
	}
	for _, deployArea := range ecd.deployAreas {
		deployArea.draw(screen)
	}
	ecd.actionButton.draw(screen)
}

func (ecd *EbitenCanonDeck) FiringUpdate() {
	bulletsInField := false
	for _, canon := range ecd.ebitenCanons {
		canon.firingUpdate()
		if canon.bullet != nil {
			bulletsInField = true
		}
	}
	if bulletsInField == false {
		ecd.Firing = false
	}
}

func (ecd *EbitenCanonDeck) Update() {
	if inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonLeft) {
		ecd.initDrag()
	}
	for _, canon := range ecd.ebitenCanons {
		canon.update()
	}

	ecd.actionButton.update()

	if ecd.draggedSprite != nil {
		for _, deployArea := range ecd.deployAreas {
			deployArea.update(&ecd.draggedSprite.EbitenSprite)
		}
	}

	if inpututil.IsMouseButtonJustReleased(ebiten.MouseButtonLeft) {
		ecd.releaseDrag()
	}
}

func (ecd *EbitenCanonDeck) initDrag() {
	ecd.actionButton.canonSprite.InitDrag()
	if ecd.actionButton.canonSprite.IsDragged {
		ecd.draggedSprite = ecd.actionButton.canonSprite
		return
	}
	for _, canon := range ecd.ebitenCanons {
		canon.sprite.InitDrag()
		if canon.sprite.IsDragged {
			ecd.draggedSprite = canon.sprite
			return
		}
	}
}

func (ecd *EbitenCanonDeck) releaseDrag() {
	if ecd.draggedSprite != nil {
		if ecd.actionButton.canonSprite.IsDragged {
			ecd.deploy()
			ecd.actionButton.canonSprite.ReleaseDrag()
		}
		for _, canon := range ecd.ebitenCanons {
			if canon.sprite.IsDragged {
				ecd.moveCanon(canon)
				canon.sprite.ReleaseDrag()
			}
		}
	}
}

func (ecd *EbitenCanonDeck) moveCanon(canon *ebitenCanon) {
	deployedArea := ecd.getDeployedAreaPosition()
	if deployedArea != nil {
		formationPlacement := *deployedArea
		if canon.formationPlacement == formationPlacement {
			return
		}
		ecd.game.MoveCannon(canon.formationPlacement, formationPlacement)
		ecd.finishTurn()
	}
}

func (ecd *EbitenCanonDeck) deploy() {
	deployedArea := ecd.getDeployedAreaPosition()
	if deployedArea != nil {
		formationPlacement := *deployedArea
		ecd.game.DeployCannon(formationPlacement)
		ecd.finishTurn()
	}
}

func (ecd *EbitenCanonDeck) finishTurn() {
	for formationPlacement := 0; formationPlacement < ecd.game.CanonDeck.CanonCapacity(); formationPlacement++ {
		canon := ecd.game.CanonDeck.Canons[game.BattleGroundColumn(formationPlacement)]
		if canon != nil {
			println(fmt.Sprintf("Setting cannon to position %d", formationPlacement))
			ec := newEbitenCanon(
				*canon,
				formationPlacement,
				ecd.actionButton.canonSprite.Image,
				getCanonCenterX(formationPlacement, ecd.game.CanonDeck.CanonCapacity()),
				canonYPlacement,
			)
			ecd.ebitenCanons[formationPlacement] = &ec
			ec.fire()
		} else {
			println(fmt.Sprintf("Deleting cannon to position %d", formationPlacement))
			delete(ecd.ebitenCanons, formationPlacement)
		}
	}
	ecd.Firing = true
}

func (ecd *EbitenCanonDeck) getDeployedAreaPosition() *int {
	for position, da := range ecd.deployAreas {
		if ebiten_sprite.Collision(da, ecd.draggedSprite) {
			return &position
		}
	}
	return nil
}

func (ecd *EbitenCanonDeck) CurrentBullets() []*EbitenCanonBullet {
	bullets := make([]*EbitenCanonBullet, 0, len(ecd.ebitenCanons))
	for _, ec := range ecd.ebitenCanons {
		if ec.bullet != nil {
			bullets = append(bullets, ec.bullet)
		}
	}
	return bullets
}
