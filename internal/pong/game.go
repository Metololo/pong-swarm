package pong

/*
using hard-coded values and unchangeable config for now to not overcomplicate the project as pong itself is not the main focus.
*/
const (
	ArenaWidth  = 800
	ArenaHeight = 600
)

type Game struct {
	ball        Ball
	leftPaddle  Paddle
	rightPaddle Paddle
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
	}
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
