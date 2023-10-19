package inputparser

import (
	"regexp"
	"strings"
)

func (p Parser) ValidateGridLine(line string) bool {
	if !regexp.MustCompile(`^[0-9]+ [0-9]+$`).MatchString(line) {
		return false
	}

	words := strings.Fields(line)
	if len(words) != 2 {
		return false
	}

	return true
}

func (p Parser) ValidatePositionLine(line string) bool {
	if !regexp.MustCompile(`^[0-9]+ [0-9]+ [N|E|S|W]$`).MatchString(line) {
		return false
	}

	words := strings.Fields(line)
	if len(words) != 3 {
		return false
	}

	return true
}

func (p Parser) ValidateInstructionsLine(line string) bool {
	if !regexp.MustCompile(`^[MLR]+$`).MatchString(line) {
		return false
	}

	return true
}
