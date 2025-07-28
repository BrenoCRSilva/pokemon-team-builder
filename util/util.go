package util

import (
	"fmt"
	"image/color"
	"math"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func RepeatingKeyPressed(key ebiten.Key) bool {
	const (
		delay    = 30
		interval = 3
	)
	d := inpututil.KeyPressDuration(key)
	if d == 1 {
		return true
	}
	if d >= delay && (d-delay)%interval == 0 {
		return true
	}
	return false
}

func CaserString(s string) string {
	caser := cases.Title(language.English)
	return caser.String(s)
}

func FormatFloat(value float64) string {
	if math.Abs(value-math.Round(value)) < 0.001 {
		return fmt.Sprintf("%.0f", value)
	}
	return fmt.Sprintf("%.1f", value)
}

func ColorComparisonRG(a, b int) color.Color {
	switch {
	case a > b:
		return color.RGBA{R: 96, G: 200, B: 96, A: 255}
	case a < b:
		return color.RGBA{R: 220, G: 70, B: 70, A: 255}
	default:
		return color.RGBA{R: 255, G: 255, B: 255, A: 255}
	}
}
