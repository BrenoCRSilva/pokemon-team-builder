package game

import (
	"github.com/BrenoCRSilva/pokemon-team-builder/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
)

type Game struct {
	Sprite
	TypeChart       TypeChart
	TeamResistances *TeamResistances
	TypeChartGrid   TypeChartGrid
	Grid            *SlotGrid
	Img             *ebiten.Image
	Title           *ebiten.Image
	Bar             *ebiten.Image
	Runes           []rune
	Text            string
	Counter         int
	Api             *api.Client
	Fonts           *Fonts
	CurrentPokemon  *api.Pokemon
	NotFound        bool
	CoverageScore   float64
	Rescore         bool
	MouseX, MouseY  float64
}
type Fonts struct {
	FontFace12 *text.GoTextFace
	FontFace16 *text.GoTextFace
	FontFace24 *text.GoTextFace
	FontFace34 *text.GoTextFace
}
type Sprite struct {
	sprite           *ebiten.Image
	spriteLoaded     bool
	isDragging       bool
	ogX, ogY         float64
	spriteX, spriteY float64
	offsetX, offsetY float64
}

type SlotGrid struct {
	Slots   []Slot
	Set     map[string]bool
	Cols    int
	Rows    int
	StartX  int
	StartY  int
	Height  int
	Width   int
	Spacing int
}

type Slot struct {
	Pokemon *api.Pokemon
	Width   int
	Height  int
	X, Y    int
	Image   *ebiten.Image
	Scale   float64
	slotted bool
	changed bool
}

type TypeChartGrid struct {
	typeImages                [GRID_ROWS][GRID_COLS]*ebiten.Image
	effectiveness             [GRID_ROWS][GRID_COLS]float64
	maxImgWidth, maxImgHeight int
	cellWidth, cellHeight     int
}

type TypeChart map[string]*ebiten.Image

type TeamResistances map[string]float64
