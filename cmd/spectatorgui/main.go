package main

import (
	"image/color"
	"log"

	"github.com/Metololo/pong-swarm/internal/pong"
	"github.com/hajimehoshi/ebiten/v2"
)

const (
	screenWidth     = pong.ArenaWidth
	screenHeight    = pong.ArenaHeight
	gap             = 16
	lineWidth       = 8
	borderThickness = 16
)

var elementsColor = color.RGBA{R: 145, G: 112, B: 145, A: 150}

type Game struct {
	pongGame *pong.Game
}

func (g *Game) Update() error {
	if g.pongGame.Status() == pong.NotStarted {
		g.pongGame.Start()
	}
	dt := 1.0 / float64(ebiten.TPS())
	g.pongGame.Step(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.drawCenterLine(screen)
	g.drawBorders(screen)
	g.DrawPaddles(screen)
	g.DrawBall(screen)
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

	for y := borderThickness; y < screenHeight-borderThickness; y += int(g.pongGame.Ball().Size()) + gap {
		for dx := 0; dx < lineWidth; dx++ {
			for dy := 0; dy < int(g.pongGame.Ball().Size()) && y+dy < screenHeight; dy++ {
				screen.Set(x+dx, y+dy, elementsColor)
			}
		}
	}
}

func (g *Game) DrawPaddles(screen *ebiten.Image) {
	leftPaddle, rightPaddle := g.pongGame.Paddles()

	g.drawLeftPaddle(screen, leftPaddle)
	g.drawRightPaddle(screen, rightPaddle)
}

func (g *Game) drawLeftPaddle(screen *ebiten.Image, paddle pong.Paddle) {
	x, y := paddle.Position()
	width, height := paddle.Dimensions()

	for dx := 0; dx < int(width); dx++ {
		for dy := 0; dy < int(height); dy++ {
			screen.Set(int(x)+dx, int(y)+dy, color.RGBA{R: 200, G: 0, B: 0, A: 150})
		}
	}
}

func (g *Game) drawRightPaddle(screen *ebiten.Image, paddle pong.Paddle) {
	x, y := paddle.Position()
	width, height := paddle.Dimensions()

	for dx := 0; dx < int(width); dx++ {
		for dy := 0; dy < int(height); dy++ {
			screen.Set(int(x)+dx, int(y)+dy, color.RGBA{R: 0, G: 0, B: 200, A: 150})
		}
	}
}
func (g *Game) DrawBall(screen *ebiten.Image) {
	ball := g.pongGame.Ball()
	x, y := ball.Position()
	size := ball.Size()

	for dx := 0; dx < int(size); dx++ {
		for dy := 0; dy < int(size); dy++ {
			screen.Set(int(x)+dx, int(y)+dy, color.RGBA{R: 250, G: 250, B: 250, A: 150})
		}
	}
}
func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := pong.NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong Swarm")

	if err := ebiten.RunGame(&Game{pongGame: game}); err != nil {
		log.Fatal(err)
	}
}
