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
