package pong

import (
	"math"
	"math/rand/v2"
)

/*
using hard-coded values and unchangeable config as pong itself is not the main focus.
*/
const (
	ArenaWidth     = 800
	ArenaHeight    = 600
	BallServeSpeed = 400.0
	BallPlaySpeed  = 700.0
	PaddleSpeed    = 350.0
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
			speed:  PaddleSpeed,
		},
		rightPaddle: Paddle{
			x:      768,
			y:      260,
			width:  12,
			height: 80,
			speed:  PaddleSpeed,
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
	g.handleBallCollision()
}
func (g *Game) handleBallCollision() {
	if g.ball.y <= 0 && g.ball.vy < 0 {
		g.ball.y = 0
		g.ball.vy = -g.ball.vy
	}

	if g.ball.y+g.ball.size >= ArenaHeight && g.ball.vy > 0 {
		g.ball.y = ArenaHeight - g.ball.size
		g.ball.vy = -g.ball.vy
	}

	if g.ball.vx < 0 && g.checkPaddleCollision(g.leftPaddle) {
		g.ball.x = g.leftPaddle.x + g.leftPaddle.width
		g.bounceOffPaddle(g.leftPaddle, 1)
	}

	if g.ball.vx > 0 && g.checkPaddleCollision(g.rightPaddle) {
		g.ball.x = g.rightPaddle.x - g.ball.size
		g.bounceOffPaddle(g.rightPaddle, -1)
	}
}

func (g *Game) checkPaddleCollision(paddle Paddle) bool {
	ballX, ballY := g.ball.Position()
	paddleX, paddleY := paddle.Position()
	paddleWidth, paddleHeight := paddle.Dimensions()

	return ballX < paddleX+paddleWidth &&
		ballX+g.ball.size > paddleX &&
		ballY < paddleY+paddleHeight &&
		ballY+g.ball.size > paddleY
}

func (g *Game) bounceOffPaddle(paddle Paddle, direction float64) {
	ballCenter := g.ball.y + g.ball.size/2
	paddleCenter := paddle.y + paddle.height/2

	offset := (ballCenter - paddleCenter) / (paddle.height / 2)
	offset = math.Max(-1, math.Min(1, offset))

	const maxAngle = math.Pi / 3
	angle := offset * maxAngle

	g.ball.vx = direction * BallPlaySpeed * math.Cos(angle)
	g.ball.vy = BallPlaySpeed * math.Sin(angle)
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
