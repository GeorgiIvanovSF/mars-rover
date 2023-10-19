package main

import (
	"fmt"

	"github.com/GeorgiIvanovSF/mars-rover/inputparser"
	"github.com/GeorgiIvanovSF/mars-rover/navigation"
	"github.com/GeorgiIvanovSF/mars-rover/rover"
)

func main() {

	parser := inputparser.Parser{
		FilePath: "resources/example-commands.txt",
	}

	inputLines, err := parser.GetFileLines()

	if err != nil || inputLines == nil {
		fmt.Printf("There was an Error in parsing the instructions file: %s\n", err.Error())
	}

	gridLine := inputLines[0]

	var grid navigation.Grid
	if parser.ValidateGridLine(gridLine) {
		grid = parser.GetGridDefinition(gridLine)
	}

	var rovers []rover.Rover

	if grid.IsValid() && len(inputLines) >= 3 {
		for i := 1; i < len(inputLines); i = i + 2 {
			if parser.ValidatePositionLine(inputLines[i]) && parser.ValidateInstructionsLine(inputLines[i+1]) {
				rovers = append(rovers, rover.Rover{
					Position:     parser.GetPositionDefinition(inputLines[i]),
					Instructions: navigation.BuildInstructions(inputLines[i+1]),
				})
			}
		}
	}

	for _, rover := range rovers {
		rover.RunInstructions()
		fmt.Printf("%d %d %s\n", rover.Position.X, rover.Position.Y, string(rover.Position.Direction))
	}

}
