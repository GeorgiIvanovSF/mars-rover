package rover

import "github.com/GeorgiIvanovSF/mars-rover/navigation"

type RoverDefinition struct {
	PositionLine     string
	InstructionsLine string
}
type Rover struct {
	Position     navigation.Position
	Instructions []navigation.Instruction
}

func (r *Rover) Move() {
	r.Position.ChangeCoordinates()
}

func (r *Rover) Turn(instruction navigation.Instruction) {
	r.Position.ChangeDirection(instruction)
}

func (r *Rover) RunInstructions() {
	for _, instruction := range r.Instructions {
		if instruction.IsRotation() {
			r.Turn(instruction)
		} else {
			r.Move()
		}
	}
}
