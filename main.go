package main

import (
	"bytes"
	"log"

	"github.com/BrenoCRSilva/pokemon-team-builder/api"
	"github.com/BrenoCRSilva/pokemon-team-builder/game"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

func main() {
	bar, _, err := ebitenutil.NewImageFromFile("./assets/searchbar.png")
	if err != nil {
		log.Fatal(err)
	}
	img, _, err := ebitenutil.NewImageFromFile("./assets/pokeball.png")
	if err != nil {
		log.Fatal(err)
	}
	normal, _, err := ebitenutil.NewImageFromFile("./assets/normal.png")
	if err != nil {
		log.Fatal(err)
	}

	fire, _, err := ebitenutil.NewImageFromFile("./assets/fire.png")
	if err != nil {
		log.Fatal(err)
	}

	water, _, err := ebitenutil.NewImageFromFile("./assets/water.png")
	if err != nil {
		log.Fatal(err)
	}

	electric, _, err := ebitenutil.NewImageFromFile("./assets/electric.png")
	if err != nil {
		log.Fatal(err)
	}

	grass, _, err := ebitenutil.NewImageFromFile("./assets/grass.png")
	if err != nil {
		log.Fatal(err)
	}

	ice, _, err := ebitenutil.NewImageFromFile("./assets/ice.png")
	if err != nil {
		log.Fatal(err)
	}

	fighting, _, err := ebitenutil.NewImageFromFile("./assets/fighting.png")
	if err != nil {
		log.Fatal(err)
	}

	poison, _, err := ebitenutil.NewImageFromFile("./assets/poison.png")
	if err != nil {
		log.Fatal(err)
	}

	ground, _, err := ebitenutil.NewImageFromFile("./assets/ground.png")
	if err != nil {
		log.Fatal(err)
	}

	flying, _, err := ebitenutil.NewImageFromFile("./assets/flying.png")
	if err != nil {
		log.Fatal(err)
	}

	psychic, _, err := ebitenutil.NewImageFromFile("./assets/psychic.png")
	if err != nil {
		log.Fatal(err)
	}

	bug, _, err := ebitenutil.NewImageFromFile("./assets/bug.png")
	if err != nil {
		log.Fatal(err)
	}

	rock, _, err := ebitenutil.NewImageFromFile("./assets/rock.png")
	if err != nil {
		log.Fatal(err)
	}

	ghost, _, err := ebitenutil.NewImageFromFile("./assets/ghost.png")
	if err != nil {
		log.Fatal(err)
	}

	dragon, _, err := ebitenutil.NewImageFromFile("./assets/dragon.png")
	if err != nil {
		log.Fatal(err)
	}

	dark, _, err := ebitenutil.NewImageFromFile("./assets/dark.png")
	if err != nil {
		log.Fatal(err)
	}

	steel, _, err := ebitenutil.NewImageFromFile("./assets/steel.png")
	if err != nil {
		log.Fatal(err)
	}

	fairy, _, err := ebitenutil.NewImageFromFile("./assets/fairy.png")
	if err != nil {
		log.Fatal(err)
	}
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	client := api.NewClient(5 * 60 * 1000) // Cache for 5 minutes
	if err != nil {
		log.Fatal(err)
	}
	fontFace := &text.GoTextFace{
		Source: fontSource,
		Size:   24,
	}
	fontFace12 := &text.GoTextFace{
		Source: fontSource,
		Size:   12,
	}
	fontFace16 := &text.GoTextFace{
		Source: fontSource,
		Size:   16,
	}
	fontFace30 := &text.GoTextFace{
		Source: fontSource,
		Size:   34,
	}
	if err != nil {
		log.Fatal(err)
	}
	resistances := &game.TeamResistances{
		"fire":     0,
		"water":    0,
		"grass":    0,
		"electric": 0,
		"ice":      0,
		"fighting": 0,
		"poison":   0,
		"ground":   0,
		"flying":   0,
		"psychic":  0,
		"bug":      0,
		"rock":     0,
		"ghost":    0,
		"dragon":   0,
		"dark":     0,
		"steel":    0,
		"fairy":    0,
		"normal":   0,
	}
	grid := game.NewSlotGrid(2, 3, 40, 80, 400, 600, 10)
	ebiten.SetWindowSize(1200, 900)
	ebiten.SetWindowTitle("Pokemon Team Builder")
	typeChart := game.TypeChart{
		"fire":     fire,
		"water":    water,
		"grass":    grass,
		"electric": electric,
		"ice":      ice,
		"fighting": fighting,
		"poison":   poison,
		"ground":   ground,
		"flying":   flying,
		"psychic":  psychic,
		"bug":      bug,
		"rock":     rock,
		"ghost":    ghost,
		"dragon":   dragon,
		"dark":     dark,
		"steel":    steel,
		"fairy":    fairy,
		"normal":   normal,
	}
	game := &game.Game{
		TypeChart:       typeChart,
		TeamResistances: resistances,
		Bar:             bar,
		Grid:            grid,
		Img:             img,
		FontFace:        fontFace,
		FontFace12:      fontFace12,
		FontFace16:      fontFace16,
		FontFace30:      fontFace30,
		Api:             client,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
