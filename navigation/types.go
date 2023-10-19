package navigation

type Direction rune
type Instruction rune

const (
	North       Direction   = 'N'
	South       Direction   = 'S'
	East        Direction   = 'E'
	West        Direction   = 'W'
	Move        Instruction = 'M' // x y +-1
	RotateRight Instruction = 'R'
	RotateLeft  Instruction = 'L'
)

type Position struct {
	X         int
	Y         int
	Direction Direction
}

type Grid struct {
	Width  int
	Height int
}

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
