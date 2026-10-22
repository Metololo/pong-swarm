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

func (p *Paddle) moveUp(dt float64) {
	p.y -= p.speed * dt
	if p.y < 0 {
		p.y = 0
	}
}

func (p *Paddle) moveDown(dt float64) {
	p.y += p.speed * dt
	if p.y+p.height > ArenaHeight {
		p.y = ArenaHeight - p.height
	}
}
