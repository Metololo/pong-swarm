package pong

type Ball struct {
	x, y   float64
	vx, vy float64
	size   float64
}

func (b Ball) Position() (float64, float64) {
	return b.x, b.y
}

func (b Ball) Velocity() (float64, float64) {
	return b.vx, b.vy
}

func (b Ball) Size() float64 {
	return b.size
}
