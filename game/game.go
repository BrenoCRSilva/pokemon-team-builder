package game

import (
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"strings"

	"github.com/BrenoCRSilva/pokemon-team-builder/api"
	"github.com/BrenoCRSilva/pokemon-team-builder/util"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

const (
	GRID_COLS        = 6
	GRID_ROWS        = 3
	GRID_OFFSET_X    = 70
	GRID_OFFSET_Y    = 700
	TEXT_CELL_WIDTH  = 40
	TEXT_CELL_HEIGHT = 30
	TITLE_WIDTH      = 20
)

var RECT_COLOR = color.RGBA{80, 80, 80, 255}

func (g *Game) calculateMaxImageDimensions() (int, int) {
	maxImgWidth, maxImgHeight := 0, 0

	for row := 0; row < GRID_ROWS; row++ {
		for col := 0; col < GRID_COLS-1; col++ {
			if g.TypeChartGrid.typeImages[row][col] != nil {
				bounds := g.TypeChartGrid.typeImages[row][col].Bounds()
				if bounds.Dx() > maxImgWidth {
					maxImgWidth = bounds.Dx()
				}
				if bounds.Dy() > maxImgHeight {
					maxImgHeight = bounds.Dy()
				}
			}
		}
	}

	return maxImgWidth, maxImgHeight
}

func (g *Game) formatFloat(value float64) string {
	if math.Abs(value-math.Round(value)) < 0.001 {
		return fmt.Sprintf("%.0f", value)
	}
	return fmt.Sprintf("%.1f", value)
}

func (g *Game) isInsideSlot(i int) bool {
	rect2 := image.Rect(
		g.Grid.Slots[i].X,
		g.Grid.Slots[i].Y,
		g.Grid.Slots[i].X+g.Grid.Slots[i].Width,
		g.Grid.Slots[i].Y+g.Grid.Slots[i].Height,
	)
	pointInSlot := image.Point{X: int(g.MouseX), Y: int(g.MouseY)}
	return pointInSlot.In(rect2)
}

func (g *Game) DeleteSlotted() {
	if !inpututil.IsMouseButtonJustPressed(ebiten.MouseButtonRight) {
		return
	}

	for i := range g.Grid.Slots {
		if g.Grid.Slots[i].slotted && g.isInsideSlot(i) {
			g.Grid.Set[g.Grid.Slots[i].Pokemon.Name] = false
			g.CurrentPokemon = g.Grid.Slots[i].Pokemon
			g.Grid.Slots[i].Pokemon = nil
			g.Grid.Slots[i].slotted = false
			g.Grid.Slots[i].changed = true
			break
		}
	}
}

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

func (g *Game) SearchPokemon() {
	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.spriteLoaded = false

		pokemon, err := g.Api.FetchPokemon(g.Text)
		if err != nil {
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

func NewTypeChart() TypeChart {
	return TypeChart{}
}

func NewSlotGrid(cols, rows, startX, startY, gridW, gridH, spacing int) *SlotGrid {
	grid := &SlotGrid{
		Cols:    cols,
		Rows:    rows,
		StartX:  startX,
		StartY:  startY,
		Width:   gridW,
		Height:  gridH,
		Spacing: spacing,
		Slots:   make([]Slot, cols*rows),
		Set:     make(map[string]bool),
	}

	slotWidth := (gridW - (cols+1)*spacing) / cols
	slotHeight := (gridH - (rows+1)*spacing) / rows

	for i := 0; i < len(grid.Slots); i++ {
		col := i % cols
		row := i / cols
		x := startX + col*(slotWidth+spacing) + spacing
		y := startY + row*(slotHeight+spacing) + spacing

		grid.Slots[i] = Slot{
			X:       x,
			Y:       y,
			Width:   slotWidth,
			Height:  slotHeight,
			changed: false,
		}
	}

	return grid
}

func (sg *SlotGrid) SetSlotImage(slotIndex int, img *ebiten.Image) {
	if slotIndex < 0 || slotIndex >= len(sg.Slots) {
		return
	}

	slot := &sg.Slots[slotIndex]
	slot.Image = img

	if img != nil {
		imgW, imgH := img.Bounds().Dx(), img.Bounds().Dy()
		scaleX := float64(slot.Width) / float64(imgW)
		scaleY := float64(slot.Height) / float64(imgH)
		slot.Scale = min(scaleX, scaleY)
	}
}

func (g *Game) resetResistances() {
	for key := range *g.TeamResistances {
		(*g.TeamResistances)[key] = 0
	}
}

func (g *Game) getTeamResistances(i int) {
	if g.Grid.Slots[i].Pokemon == nil {
		return
	}

	for _, typing := range g.Grid.Slots[i].Pokemon.Types {
		for _, half := range typing.Details.DamageRelations.HalfDamageFrom {
			(*g.TeamResistances)[half.Name] += 1
		}
		for _, no := range typing.Details.DamageRelations.NoDamageFrom {
			(*g.TeamResistances)[no.Name] += 1.5
		}
		for _, double := range typing.Details.DamageRelations.DoubleDamageFrom {
			(*g.TeamResistances)[double.Name] -= 1
		}
	}

	log.Printf("Final resistances after slot %d:", i)
	for typeName, value := range *g.TeamResistances {
		if value != 0 {
			log.Printf("  %s: %f", typeName, value)
		}
	}
}

func (g *Game) getCoverageScore() {
	var coverageSum float64
	var members int
	for _, value := range *g.TeamResistances {
		coverageSum += value
	}
	for i := range g.Grid.Slots {
		if g.Grid.Slots[i].Pokemon != nil {
			members++
		}
	}

	coverageScore := coverageSum / float64(members*4)
	if math.IsInf(coverageScore, 0) || math.IsNaN(coverageScore) {
		coverageScore = 0
	}
	g.CoverageScore = coverageScore
}

func (sg *SlotGrid) Draw(screen *ebiten.Image) {
	options := &ebiten.DrawImageOptions{}
	backgroundImage := ebiten.NewImage(sg.Width, sg.Height)
	backgroundImage.Fill(color.RGBA{20, 27, 55, 1})
	options.GeoM.Translate(float64(sg.StartX), float64(sg.StartY))
	screen.DrawImage(backgroundImage, options)

	for i := range sg.Slots {
		slot := &sg.Slots[i]

		if slot.Image == nil {
			continue
		}

		imgW, imgH := slot.Image.Bounds().Dx(), slot.Image.Bounds().Dy()
		imgDrawW := float64(imgW) * slot.Scale
		imgDrawH := float64(imgH) * slot.Scale

		offsetX := float64(slot.X) + float64(slot.Width)/2 - imgDrawW/2
		offsetY := float64(slot.Y) + float64(slot.Height)/2 - imgDrawH/2

		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(slot.Scale, slot.Scale)
		op.GeoM.Translate(offsetX, offsetY)
		screen.DrawImage(slot.Image, op)
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

func caserString(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{40, 56, 112, 1})
	g.DrawSearchBar(screen)
	g.DrawSearchBarText(screen)
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
