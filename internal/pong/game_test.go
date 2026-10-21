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
