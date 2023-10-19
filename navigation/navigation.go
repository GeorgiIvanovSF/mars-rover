package navigation

var Compass = map[Direction]map[Instruction]Direction{
	North: {RotateLeft: West,
		RotateRight: East},
	South: {RotateLeft: East,
		RotateRight: West},
	East: {RotateLeft: North,
		RotateRight: South},
	West: {RotateLeft: South,
		RotateRight: North},
}

func (p *Position) ChangeDirection(i Instruction) {
	p.Direction = Compass[p.Direction][i]
}

func (p *Position) Move() {
	switch p.Direction {
	case North:
		p.Y++
	case East:
		p.X++
	case South:
		p.Y--
	case West:
		p.X--
	}

}
