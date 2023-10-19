package rover

import "github.com/GeorgiIvanovSF/mars-rover/navigation"

type Rover struct {
	X            int
	Y            int
	Position     navigation.Position
	Instructions []navigation.Instruction
}

func (r *Rover) Move() {
	r.Position.Change()
}

func (r *Rover) RunInstructions() {
	for _, ins := range r.Instructions {
		if ins.IsRotation() {
			r.Position.ChangeDirection(ins)
		} else {
			r.Move()
		}
	}
}
