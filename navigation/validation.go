package navigation

import (
	"golang.org/x/exp/slices"
)

func (g *Grid) IsValid() bool {
	if g.Height < 0 || g.Width < 0 {
		return false
	}
	return true
}

func (p *Position) IsValid() bool {
	if p.X < 0 || p.Y < 0 {
		return false
	}
	return true
}

func (p *Position) IsWithinGrid(g Grid) bool {
	if p.X > g.Width || p.Y > g.Height {
		return false
	}
	return true
}

func (i Instruction) IsValid() bool {
	return slices.Contains([]Instruction{RotateLeft, RotateRight, Move}, i)
}

func (i Instruction) IsRotation() bool {
	return slices.Contains([]Instruction{RotateLeft, RotateRight}, i)
}

func (d Direction) IsValid() bool {
	return slices.Contains([]Direction{North, South, East, West}, d)
}
