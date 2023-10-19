package navigation

func (p *Position) ChangeDirection(i Instruction) {
	// 1 1 N
	p.Direction = Compass[p.Direction][i]
}
