package pong

type Config struct {
	ArenaWidth  int
	ArenaHeight int
}

func DefaultConfig() Config {
	return Config{
		ArenaWidth:  800,
		ArenaHeight: 600,
	}
}
