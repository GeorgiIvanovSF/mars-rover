package main

import (
	"github.com/GeorgiIvanovSF/mars-rover/navigation"
	"github.com/GeorgiIvanovSF/mars-rover/rover"
)

func main() {

	var pos navigation.Position
	pos.X = 0
	var r rover.Rover
	r.Position = navigation.NewPosition(0, 0, 'N')

	r.Instructions = navigation.BuildInstructions("MMRMMRMMRMMR")
	r.RunInstructions()

	r.Instructions = navigation.BuildInstructions("MMLMMLMMLMML")
	r.RunInstructions()
}
