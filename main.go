package main

import (
	"errors"
	"flag"
	"fmt"

	"github.com/GeorgiIvanovSF/mars-rover/inputparser"
	"github.com/GeorgiIvanovSF/mars-rover/navigation"
	"github.com/GeorgiIvanovSF/mars-rover/rover"
)

func main() {
	cmdFilePathInput := flag.String("inputFile", "resources/example-commands.txt", "full or relative path of the file containing the input definitions")
	flag.Parse()

	parser := inputparser.Parser{
		FilePath: *cmdFilePathInput,
	}

	inputLines, err := parser.GetFileLines()

	if err != nil {
		panic(err)
	}

	if inputLines == nil || len(inputLines) < 3 || (len(inputLines)-1)%2 != 0 {
		panic(errors.New("Not enough lines in input file"))
	}

	gridLine := inputLines[0]

	var theGrid navigation.Grid
	if parser.ValidateGridLine(gridLine) {
		theGrid = parser.GetGridDefinition(gridLine)
	} else {
		panic(errors.New("Invalid Grid input: " + gridLine))
	}

	// Somewhere once upon a time I've stumbled upon one line - "You should be able to read your code like a sentence".
	// "If the grid is not valid panic with new error" sounds good to me.
	if !theGrid.IsValid() {
		panic(errors.New("Grid was built with invalid arguments: " + gridLine))
	}

	var rovers []rover.Rover

	for i := 1; i < len(inputLines); i = i + 2 {
		if parser.ValidatePositionLine(inputLines[i]) && parser.ValidateInstructionsLine(inputLines[i+1]) {
			position := parser.GetPositionDefinition(inputLines[i])
			if position.IsValid() {
				rovers = append(rovers, rover.Rover{
					Position:     position,
					Instructions: navigation.BuildInstructions(inputLines[i+1]),
				})
			}
		}
	}

	for _, rover := range rovers {
		if len(rover.Instructions) < 1 {
			panic(errors.New("There is a rover without Instructions at position " + rover.Position.String()))
		}
		rover.RunInstructions(theGrid)
		fmt.Println(rover.Position.String())
	}
}
