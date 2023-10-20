package navigation

type Direction rune
type Instruction rune

const (
	North       Direction   = 'N'
	East        Direction   = 'E'
	South       Direction   = 'S'
	West        Direction   = 'W'
	Move        Instruction = 'M'
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

type compass struct{}
