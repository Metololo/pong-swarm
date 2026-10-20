package pong_test

import (
	"testing"

	"github.com/Metololo/pong-swarm/internal/pong"
)

func TestGameShouldInitialize(t *testing.T) {
	cfg := pong.DefaultConfig()
	game, err := pong.NewGame(cfg)
	if err != nil {
		t.Fatalf("Failed to initialize game: %v", err)
	}

	if game == nil {
		t.Fatal("Game instance is nil after initialization")
	}

	if w, h := game.Dimensions(); w != cfg.ArenaWidth || h != cfg.ArenaHeight {
		t.Fatalf("Game dimensions mismatch: got (%d, %d), want (%d, %d)", w, h, cfg.ArenaWidth, cfg.ArenaHeight)
	}

}

func TestGameShouldFailWithInvalidConfig(t *testing.T) {
	invalidConfigs := []pong.Config{
		{ArenaWidth: -1, ArenaHeight: 600},
		{ArenaWidth: 800, ArenaHeight: -1},
		{ArenaWidth: 0, ArenaHeight: 600},
		{ArenaWidth: 800, ArenaHeight: 0},
	}

	for _, cfg := range invalidConfigs {
		_, err := pong.NewGame(cfg)
		if err == nil {
			t.Fatalf("Expected error for invalid config: %+v", cfg)
		}
	}
}
