package starter

import (
	"slices"
	"strings"
	"testing"

	"github.com/mdw-aoc/inputs"
	_ "github.com/mdw-go/funcy"
	"github.com/mdw-go/grid"
	_ "github.com/mdw-go/must"
	"github.com/mdw-go/set"
	_ "github.com/mdw-go/set"
	"github.com/mdw-go/testing/should"
)

const TODO = -1

var (
	inputLines  = slices.Collect(inputs.Read(2023, 16).Lines())
	sampleLines = []string{
		strings.TrimSpace(`  .|...\....  `),
		strings.TrimSpace(`  |.-.\.....  `),
		strings.TrimSpace(`  .....|-...  `),
		strings.TrimSpace(`  ........|.  `),
		strings.TrimSpace(`  ..........  `),
		strings.TrimSpace(`  .........\  `),
		strings.TrimSpace(`  ..../.\\..  `),
		strings.TrimSpace(`  .-.-/..|..  `),
		strings.TrimSpace(`  .|....-|.\  `),
		strings.TrimSpace(`  ..//.|....  `),
	}
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) Setup() {
}
func (this *Suite) TestPart1Samples() {
	this.So(this.Part1(sampleLines), should.Equal, 46)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputLines), should.Equal, 6795)
}
func (this *Suite) TestPart2Samples() {
	this.So(this.Part2(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) int {
	world := make(map[grid.Point[int]]rune)
	for y, line := range lines {
		for x, char := range line {
			world[grid.NewPoint(x, y)] = char
		}
	}
	energized := set.Of[Turtle]()
	turtle := Turtle{At: grid.Point[int]{}, To: right}
	queue := []Turtle{turtle}
	for len(queue) > 0 {
		turtle, queue = queue[0], queue[1:]
		if energized.Contains(turtle) {
			continue
		}
		at, ok := world[turtle.At]
		if !ok {
			continue
		}
		energized.Add(turtle)
		for _, direction := range point(turtle.To, at) {
			queue = append(queue, Turtle{At: turtle.At.Move(direction), To: direction})
		}
	}
	energizedPoints := set.Of[grid.Point[int]]()
	for turtle := range energized {
		energizedPoints.Add(turtle.At)
	}
	this.emitEnergizedPoints(lines, energizedPoints)
	return energizedPoints.Len()
}

func (this *Suite) emitEnergizedPoints(lines []string, energizedPoints set.Set[grid.Point[int]]) {
	var buffer strings.Builder
	for y, line := range lines {
		for x, _ := range line {
			if energizedPoints.Contains(grid.NewPoint(x, y)) {
				buffer.WriteString("#")
			} else {
				buffer.WriteString(".")
			}
		}
		buffer.WriteString("\n")
	}
	this.Println("\n" + buffer.String())
}
func point(d grid.Direction[int], at rune) (result []grid.Direction[int]) {
	switch at {
	case '-':
		if d == up || d == down {
			return append(result, right, left)
		}
	case '|':
		if d == left || d == right {
			return append(result, up, down)
		}
	case '/':
		switch d {
		case up:
			return append(result, right)
		case down:
			return append(result, left)
		case right:
			return append(result, up)
		case left:
			return append(result, down)
		}
	case '\\':
		switch d {
		case up:
			return append(result, left)
		case down:
			return append(result, right)
		case right:
			return append(result, down)
		case left:
			return append(result, up)
		}
	}
	return append(result, d)
}
func (this *Suite) Part2(lines []string) any {
	return TODO
}

type Turtle struct {
	At grid.Point[int]
	To grid.Direction[int]
}

var (
	up    = grid.NewDirection(0, -1)
	down  = grid.NewDirection(0, 1)
	right = grid.NewDirection(1, 0)
	left  = grid.NewDirection(-1, 0)
)
