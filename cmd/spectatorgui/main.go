package main

import (
	"image/color"
	"log"

	"github.com/Metololo/pong-swarm/internal/pong"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth  = 800
	screenHeight = 400
)

type Game struct {
	pongGame *pong.Game
}

func (g *Game) Update() error {
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.drawCenterLine(screen)
}

func (g *Game) drawCenterLine(screen *ebiten.Image) {
	const (
		dashHeight = 16
		gap        = 16
		lineWidth  = 16
	)

	x := (screenWidth - lineWidth) / 2

	for y := 0; y < screenHeight; y += dashHeight + gap {
		for dx := 0; dx < lineWidth; dx++ {
			for dy := 0; dy < dashHeight && y+dy < screenHeight; dy++ {
				screen.Set(x+dx, y+dy, color.White)
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
