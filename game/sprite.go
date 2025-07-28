package game

import (
	"github.com/BrenoCRSilva/pokemon-team-builder/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
)

func (g *Game) DragSprite() {
	mouseX, mouseY := ebiten.CursorPosition()
	g.MouseX, g.MouseY = float64(mouseX), float64(mouseY)
	if !g.spriteLoaded || g.NotFound {
		return
	}
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.isDragging {
			if g.MouseX <= 1.2*(g.spriteX+float64(g.sprite.Bounds().Dx())) &&
				g.MouseX >= g.spriteX &&
				g.MouseY <= 1.5*(g.spriteY+float64(g.sprite.Bounds().Dy())) &&
				g.MouseY >= g.spriteY {
				g.isDragging = true
				g.offsetX = float64(mouseX) - g.spriteX
				g.offsetY = float64(mouseY) - g.spriteY
			}
		}
		if g.isDragging {
			for i := range g.Grid.Slots {
				if g.isInsideSlot(i) {
				}
			}
			g.spriteX = float64(mouseX) - g.offsetX
			g.spriteY = float64(mouseY) - g.offsetY
		}
	} else {
		if g.isDragging {
			dropped := false

			for i := range g.Grid.Slots {
				if g.isInsideSlot(i) && !g.Grid.Set[g.CurrentPokemon.Name] {
					if g.Grid.Slots[i].slotted {
						g.Grid.Slots[i].changed = true
						g.Grid.Set[g.Grid.Slots[i].Pokemon.Name] = false
					}
					g.Grid.Slots[i].Pokemon = &api.Pokemon{
						Name:      g.CurrentPokemon.Name,
						Abilities: append([]api.PokemonAbility{}, g.CurrentPokemon.Abilities...),
						Sprites:   g.CurrentPokemon.Sprites,
						Stats:     append([]api.PokemonStat{}, g.CurrentPokemon.Stats...),
						Types:     append([]api.PokemonType{}, g.CurrentPokemon.Types...),
					}
					g.Grid.Set[g.CurrentPokemon.Name] = true
					dropped = true
					g.spriteLoaded = false
					break
				}
			}

			if !dropped {
				g.spriteX = g.ogX + 60
				g.spriteY = g.ogY + 30
			}
		}
		g.isDragging = false
	}
}

func GetSprite(pokemon *api.Pokemon) (*ebiten.Image, error) {
	if pokemon.Sprites.FrontDefault == "" {
		return nil, nil
	}
	sprite, err := ebitenutil.NewImageFromURL(pokemon.Sprites.FrontDefault)
	if err != nil {
		return nil, err
	}
	return sprite, nil
}
