package navigation

import "fmt"

var compass = map[Direction]map[Instruction]Direction{
	North: {RotateLeft: West,
		RotateRight: East},
	East: {RotateLeft: North,
		RotateRight: South},
	South: {RotateLeft: East,
		RotateRight: West},
	West: {RotateLeft: South,
		RotateRight: North},
}

func (p *Position) ChangeDirection(i Instruction) {
	p.Direction = compass[p.Direction][i]
}

func (p *Position) ChangeCoordinates() {
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

func BuildInstructions(line string) []Instruction {
	var instructions []Instruction

	for _, ins := range line {
		if Instruction(ins).IsValid() {
			instructions = append(instructions, Instruction(ins))
		}
	}
	return instructions
}

func (p *Position) String() string {
	return fmt.Sprintf("%d %d %s", p.X, p.Y, string(p.Direction))
}
