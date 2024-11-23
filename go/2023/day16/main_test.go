package starter

import (
	"slices"
	"strings"
	"testing"

	"github.com/mdw-aoc/inputs"
	"github.com/mdw-go/grid"
	"github.com/mdw-go/set"
	"github.com/mdw-go/testing/should"
)

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
	this.So(this.Part2(sampleLines), should.Equal, 51)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, 7154)
}
func (this *Suite) Part1(lines []string) int {
	energizedPoints := this.flood(this.parseWorld(lines), Turtle{At: grid.Point[int]{}, To: right})
	this.emitEnergizedPoints(len(lines), energizedPoints)
	return energizedPoints.Len()
}
func (this *Suite) Part2(lines []string) any {
	world := this.parseWorld(lines)
	maxEnergy := 0
	for n := range len(lines) {
		// Starting on top row facing down
		turtle := Turtle{At: grid.NewPoint(n, 0), To: down}
		energy := this.flood(world, turtle).Len()
		if energy > maxEnergy {
			maxEnergy = energy
		}

		// Starting on bottom row facing up
		turtle = Turtle{At: grid.NewPoint(n, len(lines)-1), To: up}
		energy = this.flood(world, turtle).Len()
		if energy > maxEnergy {
			maxEnergy = energy
		}

		// Starting on left column facing right
		turtle = Turtle{At: grid.NewPoint(0, n), To: right}
		energy = this.flood(world, turtle).Len()
		if energy > maxEnergy {
			maxEnergy = energy
		}

		// Starting on right column facing left
		turtle = Turtle{At: grid.NewPoint(len(lines)-1, n), To: left}
		energy = this.flood(world, turtle).Len()
		if energy > maxEnergy {
			maxEnergy = energy
		}
	}
	return maxEnergy
}
func (this *Suite) parseWorld(lines []string) map[grid.Point[int]]rune {
	world := make(map[grid.Point[int]]rune)
	for y, line := range lines {
		for x, char := range line {
			world[grid.NewPoint(x, y)] = char
		}
	}
	return world
}
func (this *Suite) flood(world map[grid.Point[int]]rune, turtle Turtle) set.Set[grid.Point[int]] {
	energized := set.Of[Turtle]()
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
	return energizedPoints
}
func (this *Suite) emitEnergizedPoints(width int, energizedPoints set.Set[grid.Point[int]]) {
	var buffer strings.Builder
	for y := range width {
		for x := range width {
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
