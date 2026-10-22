package pong_test

import (
	"testing"

	"github.com/Metololo/pong-swarm/internal/pong"
)

func TestGameShouldInitialize(t *testing.T) {
	game := pong.NewGame()

	if game == nil {
		t.Fatal("Game instance is nil after initialization")
	}

}

func TestGameShouldBeInProgressWhenStarted(t *testing.T) {
	game := pong.NewGame()
	game.Start()

	if game.Status() != pong.InProgress {
		t.Errorf("Expected game status to be InProgress, but got %v", game.Status())
	}
}

func TestGameShouldSetBallDirectionOnStart(t *testing.T) {
	game := pong.NewGame()
	game.Start()

	ball := game.Ball()
	vx, _ := ball.Velocity()
	if vx == 0 {
		t.Errorf("Expected ball velocity to be non-zero after game start, but got %v", vx)
	}
}

func TestGameShouldUpdateBallPositionOnNextStep(t *testing.T) {
	game := pong.NewGame()
	game.Start()

	initialX, initialY := game.Ball().Position()
	dt := 1.0

	game.Step(dt)

	newX, newY := game.Ball().Position()

	if newX == initialX && newY == initialY {
		t.Errorf("Expected ball position to change after update, but it did not. Initial: (%f, %f), New: (%f, %f)", initialX, initialY, newX, newY)
	}
}
