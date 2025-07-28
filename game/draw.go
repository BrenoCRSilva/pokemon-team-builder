package game

import (
	"image/color"
	"strconv"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"github.com/hajimehoshi/ebiten/v2/vector"
)

func (g *Game) DrawInfoBackground(screen *ebiten.Image) {
	if g.CurrentPokemon == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.ogX, 140)
	background := ebiten.NewImage(600, 540)
	background.Fill(color.RGBA{20, 27, 55, 1})
	screen.DrawImage(background, op)
}

func (g *Game) DrawInfoWindow(screen *ebiten.Image) {
	if g.CurrentPokemon == nil {
		return
	}
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(g.ogX+20, 140+20)
	background := ebiten.NewImage(560, 250)
	background.Fill(color.RGBA{40, 56, 112, 255})
	screen.DrawImage(background, op)
}

func (g *Game) DrawSearchBar(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(500, 75)
	screen.DrawImage(g.Bar, op)
}

func (g *Game) DrawTitle(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(0.12, 0.12)
	op.GeoM.Translate(430, 6)
	tOpts := &text.DrawOptions{}
	tOpts.GeoM.Translate(605, 30)
	tOpts.ColorScale.ScaleWithColor(color.RGBA{R: 255, G: 204, B: 0, A: 255})
	screen.DrawImage(g.Title, op)
	text.Draw(screen, "  Team Builder", g.FontFace, tOpts)
}

func (g *Game) drawCoverageScore(screen *ebiten.Image) {
	scoreText := "Coverage Score"
	x := (GRID_COLS)*g.TypeChartGrid.cellWidth + TEXT_CELL_WIDTH + 50
	percentage := g.CoverageScore * 100
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(x), 730)
	opts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	opts2 := &text.DrawOptions{}
	opts2.GeoM.Translate(float64(x)+40, 770)
	text.Draw(screen, scoreText, g.FontFace16, opts)
	text.Draw(screen, strconv.Itoa(int(percentage))+"%", g.FontFace30, opts2)
}

func (g *Game) DrawSearchBarText(screen *ebiten.Image) {
	t := g.Text
	if g.Counter%60 < 30 {
		t += "|"
	}
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(540, 80)
	opts.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 255})
	text.Draw(screen, t, g.FontFace, opts)
}

func (g *Game) drawSprite(screen *ebiten.Image) {
	if !g.spriteLoaded || g.sprite == nil {
		return
	}
	spriteW, spriteH := g.sprite.Bounds().Dx(), g.sprite.Bounds().Dy()
	scalex := float64(200) / float64(spriteW)
	scaley := float64(200) / float64(spriteH)
	scale := min(scalex, scaley)
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Scale(scale, scale)
	op.GeoM.Translate(g.spriteX, g.spriteY)
	screen.DrawImage(g.sprite, op)
}

func (g *Game) drawNotFound(screen *ebiten.Image) {
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(g.ogX+100, g.ogY+100)
	opts.ColorScale.ScaleWithColor(color.RGBA{255, 0, 0, 255})
	text.Draw(screen, "Not found", g.FontFace30, opts)
}

func (g *Game) drawSpriteName(screen *ebiten.Image) {
	if g.sprite == nil {
		return
	}
	name := caserString(g.CurrentPokemon.Name)
	opts := &text.DrawOptions{}
	opts.GeoM.Translate(g.ogX+320, g.ogY+70) // Adjust position as needed
	opts.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})
	text.Draw(screen, name, g.FontFace30, opts)
}

func (g *Game) drawAbilities(screen *ebiten.Image) {
	if g.CurrentPokemon == nil {
		return
	}
	abilities := g.CurrentPokemon.Abilities
	for i, ability := range abilities {
		if i == 2 {
			break
		}
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(g.ogX+330, g.ogY+120+float64(i*30))
		text.Draw(screen, caserString(ability.Ability.Name), g.FontFace16, opts)
	}
}

func (g *Game) drawStats(screen *ebiten.Image) {
	if g.sprite == nil {
		return
	}
	stats := g.CurrentPokemon.Stats
	top30 := []int{96, 107, 95, 95, 92, 95}
	titleOpts := &text.DrawOptions{}
	titleOpts.GeoM.Translate(g.ogX+40+230, g.ogY+340-55)
	text.Draw(screen, "Current", g.FontFace, titleOpts)
	for i, stat := range stats {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(g.ogX+40, g.ogY+340+float64(i*30))
		text.Draw(screen, caserString(stat.Stat.Name), g.FontFace, opts)

		tOpts := &text.DrawOptions{}
		tOpts.GeoM.Translate(g.ogX+290, g.ogY+340+float64(i*30))
		tOpts.ColorScale.ScaleWithColor(statColor(stat.BaseStat, top30[i]))

		text.Draw(screen, strconv.Itoa(stat.BaseStat), g.FontFace, tOpts)
	}
}

func (g *Game) drawTop30(screen *ebiten.Image) {
	if g.sprite == nil {
		return
	}

	top30 := []int{96, 107, 95, 95, 92, 95}

	baseX := g.ogX + 250 + 200
	baseY := g.ogY + 340

	titleOpts := &text.DrawOptions{}
	titleOpts.GeoM.Translate(baseX-15, baseY-55)
	text.Draw(screen, "Top 30", g.FontFace, titleOpts)

	for i, val := range top30 {
		opts := &text.DrawOptions{}
		opts.GeoM.Translate(baseX, baseY+float64(i*30))
		text.Draw(screen, strconv.Itoa(val), g.FontFace, opts)
	}
}

func statColor(current, top30 int) color.Color {
	switch {
	case current > top30:
		return color.RGBA{R: 96, G: 200, B: 96, A: 255}
	case current < top30:
		return color.RGBA{R: 220, G: 70, B: 70, A: 255}
	default:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	}
}

func (g *Game) drawSpriteTypes(screen *ebiten.Image) {
	if g.sprite == nil {
		return
	}
	types := make([]string, len(g.CurrentPokemon.Types))
	for i, t := range g.CurrentPokemon.Types {
		types[i] = t.Details.Name
	}
	type1 := g.TypeChart[types[0]]
	opts1 := &ebiten.DrawImageOptions{}
	opts1.GeoM.Translate(g.ogX+280, g.ogY+170)
	opts1.GeoM.Scale(1.05, 1.05)
	screen.DrawImage(type1, opts1)
	if len(types) > 1 {
		type2 := g.TypeChart[types[1]]
		opts2 := &ebiten.DrawImageOptions{}
		opts2.GeoM.Translate(g.ogX+340, g.ogY+170)
		opts2.GeoM.Scale(1.05, 1.05)
		screen.DrawImage(type2, opts2)
	}
}

func (g *Game) drawBorder(screen *ebiten.Image) {
	totalGridWidth := (GRID_COLS)*g.TypeChartGrid.cellWidth + TEXT_CELL_WIDTH
	totalGridHeight := GRID_ROWS * g.TypeChartGrid.cellHeight

	borderX := float32(GRID_OFFSET_X - 30)
	borderY := float32(GRID_OFFSET_Y - 10)
	borderWidth := float32(totalGridWidth + 180)
	borderHeight := float32(totalGridHeight + 5)

	vector.StrokeRect(screen, borderX, borderY, borderWidth, borderHeight, 3,
		color.RGBA{0, 0, 0, 100}, false)
}

func (g *Game) drawGrid(screen *ebiten.Image) {
	for row := 0; row < GRID_ROWS; row++ {
		for col := 0; col < GRID_COLS; col++ {
			effectivenessValue := g.TypeChartGrid.effectiveness[row][col]

			g.drawImageCell(
				screen,
				g.TypeChartGrid.typeImages[row][col],
				effectivenessValue,
				row,
				col,
			)
		}
	}
}

func (g *Game) drawImageCell(screen *ebiten.Image, typeImage *ebiten.Image,
	effectivenessValue float64, row, col int,
) {
	if typeImage == nil {
		return
	}

	imgBounds := typeImage.Bounds()
	imgWidth := imgBounds.Dx()
	imgHeight := imgBounds.Dy()

	x := GRID_OFFSET_X + col*g.TypeChartGrid.cellWidth
	y := GRID_OFFSET_Y + row*g.TypeChartGrid.cellHeight

	drawOpts := &ebiten.DrawImageOptions{}
	drawOpts.GeoM.Translate(float64(x), float64(y))
	screen.DrawImage(typeImage, drawOpts)

	textRectX := float32(x + imgWidth + 5)
	textRectY := float32(y + (imgHeight-TEXT_CELL_HEIGHT)/2)

	g.drawEffectivenessRect(screen, textRectX, textRectY, effectivenessValue)
}

func (g *Game) drawEffectivenessRect(
	screen *ebiten.Image,
	x, y float32,
	effectivenessValue float64,
) {
	rectColor := g.getEffectivenessColor(effectivenessValue)

	vector.DrawFilledRect(screen, x, y, TEXT_CELL_WIDTH, TEXT_CELL_HEIGHT, rectColor, false)
	vector.StrokeRect(screen, x, y, TEXT_CELL_WIDTH, TEXT_CELL_HEIGHT, 1,
		color.RGBA{0, 0, 0, 255}, false)

	effectivenessText := g.formatFloat(effectivenessValue)
	textX := int(x) + (TEXT_CELL_WIDTH-len(effectivenessText)*6)/2
	textY := int(y) + TEXT_CELL_HEIGHT/2 - 8

	options := &text.DrawOptions{}
	options.GeoM.Translate(float64(textX), float64(textY))
	options.ColorScale.ScaleWithColor(color.RGBA{255, 255, 255, 255})

	text.Draw(screen, effectivenessText, g.FontFace12, options)
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

func (g *Game) drawTypeChart(screen *ebiten.Image) {
	if g.TypeChart == nil || g.TeamResistances == nil {
		return
	}
	typeImages := [GRID_ROWS][GRID_COLS]*ebiten.Image{
		{
			g.TypeChart["fire"], g.TypeChart["water"], g.TypeChart["grass"],
			g.TypeChart["electric"], g.TypeChart["ice"], g.TypeChart["fighting"],
		},
		{
			g.TypeChart["poison"], g.TypeChart["ground"], g.TypeChart["flying"],
			g.TypeChart["psychic"], g.TypeChart["bug"], g.TypeChart["rock"],
		},
		{
			g.TypeChart["ghost"], g.TypeChart["dragon"], g.TypeChart["dark"],
			g.TypeChart["steel"], g.TypeChart["fairy"], g.TypeChart["normal"],
		},
	}

	effectiveness := [GRID_ROWS][GRID_COLS]float64{
		{
			(*g.TeamResistances)["fire"], (*g.TeamResistances)["water"], (*g.TeamResistances)["grass"],
			(*g.TeamResistances)["electric"], (*g.TeamResistances)["ice"], (*g.TeamResistances)["fighting"],
		},
		{
			(*g.TeamResistances)["poison"], (*g.TeamResistances)["ground"], (*g.TeamResistances)["flying"],
			(*g.TeamResistances)["psychic"], (*g.TeamResistances)["bug"], (*g.TeamResistances)["rock"],
		},
		{
			(*g.TeamResistances)["ghost"], (*g.TeamResistances)["dragon"], (*g.TeamResistances)["dark"],
			(*g.TeamResistances)["steel"], (*g.TeamResistances)["fairy"], (*g.TeamResistances)["normal"],
		},
	}
	g.TypeChartGrid.typeImages = typeImages
	g.TypeChartGrid.effectiveness = effectiveness
	g.TypeChartGrid.maxImgWidth, g.TypeChartGrid.maxImgHeight = g.calculateMaxImageDimensions()
	g.TypeChartGrid.cellWidth = g.TypeChartGrid.maxImgWidth + TEXT_CELL_WIDTH + 40
	g.TypeChartGrid.cellHeight = max(g.TypeChartGrid.maxImgHeight, TEXT_CELL_HEIGHT) + 20

	g.drawBorder(screen)
	g.drawGrid(screen)
}
