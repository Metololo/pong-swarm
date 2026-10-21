package pong

import "testing"

// write test for ball move
func TestBallShouldMove(t *testing.T) {
	ball := Ball{
		x:  100,
		y:  100,
		vx: 50,
		vy: 50,
	}

	dt := 1.0
	ball.move(dt)

	expectedX := 150.0
	expectedY := 150.0

	if ball.x != expectedX || ball.y != expectedY {
		t.Errorf("Expected ball position to be (%f, %f), but got (%f, %f)", expectedX, expectedY, ball.x, ball.y)
	}
}
