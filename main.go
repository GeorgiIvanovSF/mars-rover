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

	fmt.Println(string(test.Direction))
}
