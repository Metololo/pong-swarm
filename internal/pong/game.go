package pong

import "math/rand/v2"

/*
using hard-coded values and unchangeable config for now to not overcomplicate the project as pong itself is not the main focus.
*/
const (
	ArenaWidth     = 800
	ArenaHeight    = 600
	BallServeSpeed = 200.0
	BallPlaySpeed  = 350.0
)

type GameStatus int

const (
	NotStarted GameStatus = iota
	InProgress
	Finished
)

type Game struct {
	ball        Ball
	leftPaddle  Paddle
	rightPaddle Paddle
	leftScore   int
	rightScore  int
	status      GameStatus
}

func (g *Game) Status() GameStatus {
	return g.status
}

func NewGame() *Game {
	return &Game{
		ball: Ball{
			x:    392,
			y:    292,
			size: 16,
		},
		leftPaddle: Paddle{
			x:      20,
			y:      260,
			width:  12,
			height: 80,
			speed:  250,
		},
		rightPaddle: Paddle{
			x:      768,
			y:      260,
			width:  12,
			height: 80,
			speed:  250,
		},
		status:     NotStarted,
		leftScore:  0,
		rightScore: 0,
	}
}

func (g *Game) Start() {
	g.status = InProgress
	g.serveBall()
}

func (g *Game) serveBall() {
	g.ball.vx = BallServeSpeed
	g.ball.vy = 0

	if rand.IntN(2) == 0 {
		g.ball.vx = -g.ball.vx
	}
}

func (g *Game) Step(dt float64) {
	if g.status != InProgress {
		return
	}

	g.ball.move(dt)
}

func (g *Game) MoveLeftPaddleUp(dt float64) {
	g.leftPaddle.moveUp(dt)
}

func (g *Game) MoveLeftPaddleDown(dt float64) {
	g.leftPaddle.moveDown(dt)
}

func (g *Game) MoveRightPaddleUp(dt float64) {
	g.rightPaddle.moveUp(dt)
}

func (g *Game) MoveRightPaddleDown(dt float64) {
	g.rightPaddle.moveDown(dt)
}

func (g *Game) Dimensions() (int, int) {
	return ArenaWidth, ArenaHeight
}

func (g *Game) Ball() Ball {
	return g.ball
}

func (g *Game) Paddles() (Paddle, Paddle) {
	return g.leftPaddle, g.rightPaddle
}
