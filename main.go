package main

import (
	"bytes"
	"image/color"
	"log"
	"strings"

	"github.com/BrenoCRSilva/pokemon-team-builder/api"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
	"github.com/hajimehoshi/ebiten/v2/text/v2"
	"golang.org/x/image/font/gofont/goregular"
)

type Game struct {
	img              *ebiten.Image
	runes            []rune
	text             string
	counter          int
	api              *api.Client
	fontFace         *text.GoTextFace
	pokemon          *api.Pokemon
	spriteLoaded     bool
	draggable        bool
	sprite           *ebiten.Image
	spriteX, spriteY float64
	offsetX, offsetY float64
	mouseX, mouseY   float64
}

func (g *Game) Update() error {
	g.runes = ebiten.AppendInputChars(g.runes[:0])
	g.text += string(g.runes)
	mouseX, mouseY := ebiten.CursorPosition()
	g.mouseX, g.mouseY = float64(mouseX), float64(mouseY)
	if ebiten.IsMouseButtonPressed(ebiten.MouseButtonLeft) {
		if !g.draggable {
			if g.mouseX <= g.spriteX+float64(g.sprite.Bounds().Dx()) &&
				g.mouseX >= g.spriteX &&
				g.mouseY <= g.spriteY+float64(g.sprite.Bounds().Dy()) &&
				g.mouseY >= g.spriteY {
				g.draggable = true
				g.offsetX = float64(mouseX) - g.spriteX
				g.offsetY = float64(mouseY) - g.spriteY
				log.Printf(
					"Draggable: %v, OffsetX: %f, OffsetY: %f, MouseX: %d, MouseY: %d, SpriteX: %f, SpriteY: %f",
					g.draggable,
					g.offsetX,
					g.offsetY,
					g.mouseX,
					g.mouseY,
					g.spriteX,
					g.spriteY,
				)
			}
		}
		if g.draggable {
			g.spriteX = float64(mouseX) - g.offsetX
			g.spriteY = float64(mouseY) - g.offsetY
		}
	} else {
		g.draggable = false
	}

	check := strings.Split(g.text, " ")
	if len(check) > 1 {
		g.text = check[0]
	}

	if repeatingKeyPressed(ebiten.KeyBackspace) {
		if len(g.text) >= 1 {
			g.text = g.text[:len(g.text)-1]
		}
	}

	if inpututil.IsKeyJustPressed(ebiten.KeyEnter) {
		g.spriteLoaded = false

		pokemon, err := g.api.FetchPokemon(g.text)
		if err != nil {
			log.Printf("Error fetching Pokemon: %v", err)
			return nil
		}

		g.pokemon = &pokemon
		g.text = ""

		if g.pokemon.Sprites.FrontDefault != "" {
			sprite, err := ebitenutil.NewImageFromURL(g.pokemon.Sprites.FrontDefault)
			if err != nil {
				log.Printf("Error loading Pokemon sprite: %v", err)
			} else {
				g.sprite = sprite
				g.spriteLoaded = true
			}
		}

	}

	g.counter++
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.RGBA{240, 248, 255, 255}) // Light blue background

	// Grid configuration
	cols := 2
	rows := 3
	gridStartX := 20
	gridStartY := 300
	gridWidth := 400
	gridHeight := 600

	// Calculate slot dimensions
	slotWidth := (gridWidth - (cols+1)*10) / cols
	slotHeight := (gridHeight - (rows+1)*10) / rows

	for i := 0; i < 6; i++ {
		col := i % cols
		row := i / cols

		x := gridStartX + col*(slotWidth+10) + 10
		y := gridStartY + row*(slotHeight+10) + 10

		// Compute scaling factors
		imgW, imgH := g.img.Bounds().Dx(), g.img.Bounds().Dy()
		scaleX := float64(slotWidth) / float64(imgW)
		scaleY := float64(slotHeight) / float64(imgH)

		// Optionally preserve aspect ratio
		scale := min(scaleX, scaleY)

		// Center the image inside the slot
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)

		imgDrawW := float64(imgW) * scale
		imgDrawH := float64(imgH) * scale

		offsetX := float64(x) + float64(slotWidth)/2 - imgDrawW/2
		offsetY := float64(y) + float64(slotHeight)/2 - imgDrawH/2
		op.GeoM.Translate(offsetX, offsetY)

		screen.DrawImage(g.img, op)
	}

	t := g.text
	if g.counter%60 < 30 {
		t += "_"
	}

	opts := &text.DrawOptions{}
	opts.GeoM.Translate(float64(gridWidth+100), 40)          // Position text at the bottom left
	opts.ColorScale.ScaleWithColor(color.RGBA{0, 0, 0, 255}) // Black text
	text.Draw(screen, t, g.fontFace, opts)

	if g.spriteLoaded && g.sprite != nil {
		spriteW, spriteH := g.sprite.Bounds().Dx(), g.sprite.Bounds().Dy()
		scalex := float64(200) / float64(spriteW)
		scaley := float64(200) / float64(spriteH)
		scale := min(scalex, scaley)
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Scale(scale, scale)
		op.GeoM.Translate(g.spriteX, g.spriteY) // Apply draggable position
		screen.DrawImage(g.sprite, op)
	}
}

func repeatingKeyPressed(key ebiten.Key) bool {
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

func (g *Game) Layout(outsideWidth, outsideHeight int) (int, int) {
	return 1200, 900
}

func main() {
	img, _, err := ebitenutil.NewImageFromFile("./assets/pokeball.png")
	fontSource, err := text.NewGoTextFaceSource(bytes.NewReader(goregular.TTF))
	client := api.NewClient(5 * 60 * 1000) // Cache for 5 minutes
	if err != nil {
		log.Fatal(err)
	}
	fontFace := &text.GoTextFace{
		Source: fontSource,
		Size:   24,
	}
	if err != nil {
		log.Fatal(err)
	}
	ebiten.SetWindowSize(1200, 900)
	ebiten.SetWindowTitle("Pokemon Team Builder")

	game := &Game{
		img:      img,
		fontFace: fontFace,
		api:      client,
	}

	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
