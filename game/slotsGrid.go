package game

import (
	"image"
	"image/color"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

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
