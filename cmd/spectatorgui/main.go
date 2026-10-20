package main

import (
	"image/color"
	"log"

	"github.com/Metololo/pong-swarm/internal/pong"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth     = 800
	screenHeight    = 400
	dashHeight      = 8
	gap             = 8
	lineWidth       = 8
	borderThickness = 16
)

var elementsColor = color.RGBA{R: 145, G: 112, B: 145, A: 150}

type Game struct {
	pongGame *pong.Game
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.drawCenterLine(screen)
	g.drawBorders(screen)
}

func (g *Game) drawBorders(screen *ebiten.Image) {

	for x := 0; x < screenWidth; x++ {
		for dy := 0; dy < borderThickness; dy++ {
			screen.Set(x, dy, elementsColor)                // Top
			screen.Set(x, screenHeight-1-dy, elementsColor) // Bottom
		}
	}
}

func (g *Game) drawCenterLine(screen *ebiten.Image) {
	x := (screenWidth - lineWidth) / 2

	for y := borderThickness; y < screenHeight-borderThickness; y += dashHeight + gap {
		for dx := 0; dx < lineWidth; dx++ {
			for dy := 0; dy < dashHeight && y+dy < screenHeight; dy++ {
				screen.Set(x+dx, y+dy, elementsColor)
			}
		}
	}
}
func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	gameCfg := pong.DefaultConfig()
	game, err := pong.NewGame(gameCfg)
	if err != nil {
		log.Fatalf("Failed to initialize game: %v", err)
	}

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong Swarm")

	if err := ebiten.RunGame(&Game{pongGame: game}); err != nil {
		log.Fatal(err)
	}
}
