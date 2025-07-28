package game

import (
	"bytes"
	"log"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type Assets struct {
	Title    *ebiten.Image
	Bar      *ebiten.Image
	Pokeball *ebiten.Image
	Normal   *ebiten.Image
	Fire     *ebiten.Image
	Water    *ebiten.Image
	Electric *ebiten.Image
	Grass    *ebiten.Image
	Ice      *ebiten.Image
	Fighting *ebiten.Image
	Poison   *ebiten.Image
	Ground   *ebiten.Image
	Flying   *ebiten.Image
	Psychic  *ebiten.Image
	Bug      *ebiten.Image
	Rock     *ebiten.Image
	Ghost    *ebiten.Image
	Dragon   *ebiten.Image
	Dark     *ebiten.Image
	Steel    *ebiten.Image
	Fairy    *ebiten.Image
}

func NewFonts() *Fonts {
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	if err != nil {
		log.Fatal(err)
	}

	return &Fonts{
		FontFace12: &text.GoTextFace{
			Source: fontSource,
			Size:   12,
		},
		FontFace16: &text.GoTextFace{
			Source: fontSource,
			Size:   16,
		},
		FontFace24: &text.GoTextFace{
			Source: fontSource,
			Size:   24,
		},
		FontFace34: &text.GoTextFace{
			Source: fontSource,
			Size:   34,
		},
	}
}

func NewAssets() *Assets {
	assets := &Assets{}

	var err error

	assets.Title, _, err = ebitenutil.NewImageFromFile("./assets/title.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Bar, _, err = ebitenutil.NewImageFromFile("./assets/searchbar.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Pokeball, _, err = ebitenutil.NewImageFromFile("./assets/pokeball.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Normal, _, err = ebitenutil.NewImageFromFile("./assets/normal.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Fire, _, err = ebitenutil.NewImageFromFile("./assets/fire.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Water, _, err = ebitenutil.NewImageFromFile("./assets/water.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Electric, _, err = ebitenutil.NewImageFromFile("./assets/electric.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Grass, _, err = ebitenutil.NewImageFromFile("./assets/grass.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Ice, _, err = ebitenutil.NewImageFromFile("./assets/ice.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Fighting, _, err = ebitenutil.NewImageFromFile("./assets/fighting.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Poison, _, err = ebitenutil.NewImageFromFile("./assets/poison.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Ground, _, err = ebitenutil.NewImageFromFile("./assets/ground.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Flying, _, err = ebitenutil.NewImageFromFile("./assets/flying.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Psychic, _, err = ebitenutil.NewImageFromFile("./assets/psychic.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Bug, _, err = ebitenutil.NewImageFromFile("./assets/bug.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Rock, _, err = ebitenutil.NewImageFromFile("./assets/rock.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Ghost, _, err = ebitenutil.NewImageFromFile("./assets/ghost.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Dragon, _, err = ebitenutil.NewImageFromFile("./assets/dragon.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Dark, _, err = ebitenutil.NewImageFromFile("./assets/dark.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Steel, _, err = ebitenutil.NewImageFromFile("./assets/steel.png")
	if err != nil {
		log.Fatal(err)
	}

	assets.Fairy, _, err = ebitenutil.NewImageFromFile("./assets/fairy.png")
	if err != nil {
		log.Fatal(err)
	}

	return assets
}
