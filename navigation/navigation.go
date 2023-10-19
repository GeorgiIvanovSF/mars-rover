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

func NewPosition(x int, y int, d rune) Position {
	return Position{
		X:         x,
		Y:         y,
		Direction: Direction(d),
	}
}

func (p *Position) ChangeDirection(i Instruction) {
	fmt.Printf("Change Direction from %s to %s\n", string(p.Direction), string(compass[p.Direction][i]))
	p.Direction = compass[p.Direction][i]
}

func (p *Position) ChangeCoordinates() {
	fmt.Printf("Move from X: %d, Y: %d to ", p.X, p.Y)
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

	fmt.Printf("X: %d, Y: %d\n", p.X, p.Y)

}

func BuildInstructions(row string) []Instruction {
	var instructions []Instruction

	for _, ins := range row {
		if Instruction(ins).IsValid() {
			instructions = append(instructions, Instruction(ins))
		}
	}
	return instructions
}
