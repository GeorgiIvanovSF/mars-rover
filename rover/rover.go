package rover

import "github.com/GeorgiIvanovSF/mars-rover/navigation"

type Rover struct {
	Position     navigation.Position
	Instructions []navigation.Instruction
}

func (r *Rover) Move(g navigation.Grid) {

	nextPosition := r.Position
	nextPosition.ChangeCoordinates()

	if nextPosition.IsWithinGrid(g) {
		r.Position.ChangeCoordinates()
	}

}

func (r *Rover) Turn(instruction navigation.Instruction) {
	r.Position.ChangeDirection(instruction)
}

func (r *Rover) RunInstructions(g navigation.Grid) {
	for _, instruction := range r.Instructions {
		if instruction.IsRotation() {
			r.Turn(instruction)
		} else {
			r.Move(g)
		}
	}
}
