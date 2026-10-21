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
	WinningScore   = 5
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
	serveDelay  float64
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
	if g.status != NotStarted {
		return
	}

	g.status = InProgress

	direction := 1.0
	if rand.IntN(2) == 0 {
		direction = -1
	}

	g.serveBall(direction, 0)
}

func (g *Game) serveBall(direction, angle float64) {
	g.ball.x = (ArenaWidth - g.ball.size) / 2
	g.ball.y = (ArenaHeight - g.ball.size) / 2

	g.ball.vx = direction * BallServeSpeed * math.Cos(angle)
	g.ball.vy = BallServeSpeed * math.Sin(angle)
}

func (g *Game) Step(dt float64) {
	if g.status != InProgress {
		return
	}

	if g.serveDelay > 0 {
		g.serveDelay -= dt
		if g.serveDelay >= 0 {
			return
		}

		dt = -g.serveDelay
		g.serveDelay = 0
	}

	g.ball.move(dt)
	g.handleBallCollision()
	g.checkScore()
}
func (g *Game) checkScore() {
	var serveDirection float64

	if g.ball.x+g.ball.size <= 0 {
		g.rightScore++
		serveDirection = -1
	} else if g.ball.x >= ArenaWidth {
		g.leftScore++
		serveDirection = 1
	} else {
		return
	}

	if g.leftScore >= WinningScore || g.rightScore >= WinningScore {
		g.status = Finished
		g.ball.vx = 0
		g.ball.vy = 0
		return
	}

	const maxServeAngle = math.Pi / 6
	angle := (rand.Float64()*2 - 1) * maxServeAngle

	g.serveBall(serveDirection, angle)
	g.serveDelay = 1.0
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
	if g.status != InProgress {
		return
	}
	g.leftPaddle.moveUp(dt)
}

func (g *Game) MoveLeftPaddleDown(dt float64) {
	if g.status != InProgress {
		return
	}
	g.leftPaddle.moveDown(dt)
}

func (g *Game) MoveRightPaddleUp(dt float64) {
	if g.status != InProgress {
		return
	}
	g.rightPaddle.moveUp(dt)
}

func (g *Game) MoveRightPaddleDown(dt float64) {
	if g.status != InProgress {
		return
	}
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

func (g *Game) LeftScore() int {
	return g.leftScore
}

func (g *Game) RightScore() int {
	return g.rightScore
}
