package main

import (
	"image/color"
	"log"
	"strconv"

	"github.com/Metololo/pong-swarm/internal/pong"
	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
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
	dt := 1.0 / float64(ebiten.TPS())

	if ebiten.IsKeyPressed(ebiten.KeyW) {
		g.pongGame.MoveLeftPaddleUp(dt)
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		g.pongGame.MoveLeftPaddleDown(dt)
	}

	if ebiten.IsKeyPressed(ebiten.KeyUp) {
		g.pongGame.MoveRightPaddleUp(dt)
	}
	if ebiten.IsKeyPressed(ebiten.KeyDown) {
		g.pongGame.MoveRightPaddleDown(dt)
	}

	g.pongGame.Step(dt)
	return nil
}

func (g *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	g.drawCenterLine(screen)
	g.DrawPaddles(screen)
	g.DrawBall(screen)
	g.DrawScore(screen)

	if g.pongGame.Status() == pong.Finished {
		winner := "Left player wins!"
		if g.pongGame.RightScore() >= pong.WinningScore {
			winner = "Right player wins!"
		}

		ebitenutil.DebugPrintAt(screen, winner, screenWidth/2-60, screenHeight/2)
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

func (g *Game) DrawScore(screen *ebiten.Image) {
	leftScore := strconv.Itoa(g.pongGame.LeftScore())
	rightScore := strconv.Itoa(g.pongGame.RightScore())

	ebitenutil.DebugPrintAt(screen, leftScore, screenWidth/2-80, 50)
	ebitenutil.DebugPrintAt(screen, rightScore, screenWidth/2+80, 50)
}

func (g *Game) Layout(_, _ int) (int, int) {
	return screenWidth, screenHeight
}

func main() {
	game := pong.NewGame()

	ebiten.SetWindowSize(screenWidth, screenHeight)
	ebiten.SetWindowTitle("Pong Swarm")

	g := &Game{pongGame: game}
	if g.pongGame.Status() == pong.NotStarted {
		g.pongGame.Start()
	}
	if err := ebiten.RunGame(g); err != nil {
		log.Fatal(err)
	}
}
