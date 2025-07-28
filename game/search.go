package game

import (
	"log"
	"strings"

	"github.com/BrenoCRSilva/pokemon-team-builder/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

func (g *Game) SearchPokemon() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.spriteLoaded = false

		pokemon, err := g.Api.FetchPokemon(g.Text)
		if err != nil || g.Text == "" {
			g.NotFound = true
			log.Printf("Error fetching Pokemon: %v", err)
		} else {
			g.NotFound = false
		}

		g.CurrentPokemon = &pokemon
		g.Text = ""

		sprite, err := GetSprite(g.CurrentPokemon)
		if err != nil {
			log.Printf("Error getting sprite: %v", err)
		}
		g.spriteLoaded = true
		g.sprite = sprite
		g.ogX = 500
		g.ogY = 140
		g.spriteX = g.ogX + 60
		g.spriteY = g.ogY + 30

	}
}

func (g *Game) GetSearchInput() {
	g.Runes = ebiten.AppendInputChars(g.Runes[:0])
	g.Text += string(g.Runes)
	g.Text = strings.TrimSpace(g.Text)
	if util.RepeatingKeyPressed(ebiten.KeyBackspace) {
		if len(g.Text) >= 1 {
			g.Text = g.Text[:len(g.Text)-1]
		}
	}
}
