package game

import (
	"image/color"
	"log"
	"math"
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

func NewTypeChart(assets *Assets) TypeChart {
	return TypeChart{
		"fire":     assets.Fire,
		"water":    assets.Water,
		"grass":    assets.Grass,
		"electric": assets.Electric,
		"ice":      assets.Ice,
		"fighting": assets.Fighting,
		"poison":   assets.Poison,
		"ground":   assets.Ground,
		"flying":   assets.Flying,
		"psychic":  assets.Psychic,
		"bug":      assets.Bug,
		"rock":     assets.Rock,
		"ghost":    assets.Ghost,
		"dragon":   assets.Dragon,
		"dark":     assets.Dark,
		"steel":    assets.Steel,
		"fairy":    assets.Fairy,
		"normal":   assets.Normal,
	}
}

func NewTeamResistances() *TeamResistances {
	return &TeamResistances{
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
}

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

func (g *Game) getEffectivenessColor(value float64) color.RGBA {
	switch {
	case value >= 2.0:
		return color.RGBA{50, 200, 50, 255}
	case value == 1.0:
		return color.RGBA{80, 80, 80, 255}
	case value == -1.0:
		return color.RGBA{80, 80, 80, 255}
	case value <= -2.0:
		return color.RGBA{200, 50, 50, 255}
	default:
		return color.RGBA{50, 50, 50, 255}
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
