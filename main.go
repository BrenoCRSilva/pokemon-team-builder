package main

import (
	"log"

	"github.com/BrenoCRSilva/pokemon-team-builder/api"
	"github.com/BrenoCRSilva/pokemon-team-builder/game"
	"github.com/hajimehoshi/ebiten/v2"
)

func main() {
	fonts := game.NewFonts()
	assets := game.NewAssets()
	typeChart := game.NewTypeChart(assets)
	resistances := game.NewTeamResistances()
	grid := game.NewSlotGrid(2, 3, 40, 80, 400, 600, 10)
	client := api.NewClient(10 * 60 * 1000)
	game := game.NewGame(
		typeChart,
		resistances,
		assets.Bar,
		assets.Title,
		grid,
		assets.Pokeball,
		fonts,
		client,
	)

	ebiten.SetWindowSize(1200, 900)
	ebiten.SetWindowTitle("Boot.dev Hackathon")

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
