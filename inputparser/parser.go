package inputparser

import (
	"bufio"
	"os"
	"strconv"
	"strings"

	"github.com/GeorgiIvanovSF/mars-rover/navigation"
)

type Parser struct {
	FilePath string
}

func (p *Parser) GetFileLines() ([]string, error) {
	var lines []string

	file, err := os.Open(p.FilePath)
	if err != nil {
		return nil, err
	}
	defer file.Close()

	scanner := bufio.NewScanner(file)
	scanner.Split(bufio.ScanLines)

	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}

	return lines, nil
}

func (p *Parser) GetGridDefinition(line string) navigation.Grid {

	words := strings.Fields(line)
	width, _ := strconv.Atoi(words[0])
	height, _ := strconv.Atoi(words[1])

	return navigation.Grid{Width: width, Height: height}
}

func (p *Parser) GetPositionDefinition(line string) navigation.Position {
	words := strings.Fields(line)
	x, _ := strconv.Atoi(words[0])
	y, _ := strconv.Atoi(words[1])
	direction := words[2][0]

	return navigation.Position{X: x, Y: y, Direction: navigation.Direction(direction)}
}
