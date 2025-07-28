package game

import (
	"image/color"
	"log"

	"github.com/BrenoCRSilva/pokemon-team-builder/api"
	"github.com/hajimehoshi/ebiten/v2"
)

func NewGame(
	typeChart TypeChart,
	resistances *TeamResistances,
	bar *ebiten.Image,
	title *ebiten.Image,
	grid *SlotGrid,
	img *ebiten.Image,
	fonts *Fonts,
	client *api.Client,
) *Game {
	return &Game{
		TypeChart:       typeChart,
		TeamResistances: resistances,
		Bar:             bar,
		Title:           title,
		Grid:            grid,
		Img:             img,
		Fonts:           fonts,
		Api:             client,
	}
}

func (g *Game) Update() error {
	g.GetSearchInput()
	g.SearchPokemon()
	g.DragSprite()
	g.DeleteSlotted()

	for i := range g.Grid.Slots {
		if g.Grid.Slots[i].Pokemon == nil && g.Grid.Slots[i].Image == nil {
			g.Grid.SetSlotImage(i, g.Img)
		} else if !g.Grid.Slots[i].slotted && g.Grid.Slots[i].Pokemon != nil {
			sprite, err := GetSprite(g.Grid.Slots[i].Pokemon)
			if err != nil {
				log.Printf("Error getting sprite: %v", err)
			}
			g.Grid.SetSlotImage(i, sprite)
			g.Grid.Slots[i].slotted = true
			g.Rescore = true
		} else if g.Grid.Slots[i].changed {
			if g.Grid.Slots[i].Pokemon == nil {
				g.Grid.SetSlotImage(i, g.Img)
				g.Grid.Slots[i].slotted = false
				g.Grid.Slots[i].changed = false
				g.Rescore = true
				continue
			}
			sprite, err := GetSprite(g.Grid.Slots[i].Pokemon)
			if err != nil {
				log.Printf("Error getting sprite: %v", err)
			}
			g.Grid.SetSlotImage(i, sprite)
			g.Grid.Slots[i].changed = false
			g.Rescore = true
		}
	}

	if g.Rescore {
		g.resetResistances()
		for i := range g.Grid.Slots {
			if g.Grid.Slots[i].Pokemon != nil {
				g.getTeamResistances(i)
			}
		}
		g.getCoverageScore()
		g.Rescore = false
	}

	g.Counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 56, 112, 1})
	g.DrawSearchBar(screen)
	g.DrawSearchBarText(screen)
	g.DrawTitle(screen)
	g.Grid.Draw(screen)
	g.DrawInfoBackground(screen)
	g.DrawInfoWindow(screen)
	g.drawSprite(screen)
	g.drawSpriteName(screen)
	g.drawSpriteTypes(screen)
	g.drawStats(screen)
	g.drawTop30(screen)
	g.drawAbilities(screen)
	g.drawCoverageScore(screen)
	g.drawTypeChart(screen)
	if g.NotFound {
		g.drawNotFound(screen)
	}
}

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1200, 900
}
