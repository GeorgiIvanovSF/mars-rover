package main

import (
	"fmt"

	"github.com/GeorgiIvanovSF/mars-rover/navigation"
)

func main() {

	test := navigation.Position{
		X:         1,
		Y:         2,
		Direction: 'N',
	}

	test.ChangeDirection(navigation.RotateLeft)
	test.Move()

	fmt.Printf("X: %d, Y: %d, Direction: %s", test.X, test.Y, string(test.Direction))
}
