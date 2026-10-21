package pong

type Paddle struct {
	x, y          float64
	width, height float64
	speed         float64
}

func (p Paddle) Position() (float64, float64) {
	return p.x, p.y
}

func (p Paddle) Dimensions() (float64, float64) {
	return p.width, p.height
}

func (p Paddle) Speed() float64 {
	return p.speed
}
