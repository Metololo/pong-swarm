package pong

import "fmt"

var ErrInvalidArenaWidth = fmt.Errorf("invalid arena width: must be positive")
var ErrInvalidArenaHeight = fmt.Errorf("invalid arena height: must be positive")

type Game struct {
	config Config
}

func NewGame(cfg Config) (*Game, error) {
	if cfg.ArenaWidth <= 0 {
		return nil, ErrInvalidArenaWidth
	}
	if cfg.ArenaHeight <= 0 {
		return nil, ErrInvalidArenaHeight
	}
	return &Game{config: cfg}, nil
}

func (g *Game) Dimensions() (width, height int) {
	return g.config.ArenaWidth, g.config.ArenaHeight
}
