package day10

import (
	"fmt"
	"slices"
	"testing"

	"github.com/mdw-aoc/inputs/v2/inputs"
	"github.com/mdw-go/set"
	"github.com/mdw-go/testing/should"
)

const TODO = -1

var (
	inputLines = slices.Collect(inputs.Read(2023, 10).Lines())
	sampleLines = []string{
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
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

func (this *Suite) TestPart1A() {
	this.So(this.Part1(sampleLines), should.Equal, 8)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputLines), should.Equal, 6886)
}
func (this *Suite) TestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) any {
	return CircuitLength(ParseInput(lines)) / 2
}
func (this *Suite) Part2(lines []string) any {
	return TODO
}

func CircuitLength(start Point, field map[Point]string) (result int) {
	queue := []Point{start}
	frontier := set.Of[Point](start)
	for {
		at := queue[0]
		queue = queue[1:]
		a, b := follow(at, field)
		result++
		if !frontier.Contains(a) {
			frontier.Add(a)
			queue = append(queue, a)
			continue
		}
		if !frontier.Contains(b) {
			frontier.Add(b)
			queue = append(queue, b)
			continue
		}
		if a == start || b == start {
			return result
		}

	}
}
func follow(from Point, field map[Point]string) (a, b Point) {
	n, s, e, w := neighbors(from)
	switch field[from] {
	case "|":
		return n, s
	case "-":
		return e, w
	case "L":
		return n, e
	case "F":
		return s, e
	case "7":
		return s, w
	case "J":
		return n, w
	}
	panic(fmt.Sprintln("cannot follow:", from))
}

type Point struct{ row, col int }

func ParseInput(lines []string) (start Point, field map[Point]string) {
	field = make(map[Point]string)
	for row, line := range lines {
		for col, char := range line {
			if char == '.' {
				continue
			}
			field[Point{row: row, col: col}] = string(char)
		}
	}
	for point, char := range field {
		if char == "S" {
			start = point
			field[point] = inferS(field, point)
		}
	}
	return start, field
}

func neighbors(p Point) (n, s, e, w Point) {
	n = Point{row: p.row - 1, col: p.col}
	s = Point{row: p.row + 1, col: p.col}
	e = Point{row: p.row, col: p.col + 1}
	w = Point{row: p.row, col: p.col - 1}
	return n, s, e, w
}
func inferS(field map[Point]string, p Point) string {
	n, s, e, w := neighbors(p)
	N, S, E, W := field[n], field[s], field[e], field[w]
	var pointers string
	switch N {
	case "|", "7", "F":
		pointers += "N"
	}
	switch S {
	case "|", "L", "J":
		pointers += "S"
	}
	switch E {
	case "-", "J", "7":
		pointers += "E"
	}
	switch W {
	case "-", "L", "F":
		pointers += "W"
	}
	return lookupPointers(pointers)
}
func lookupPointers(pointers string) string {
	pointerSymbols := map[string]string{
		"NS": "|",
		"EW": "-",
		"NE": "L",
		"NW": "J",
		"SW": "7",
		"SE": "F",
	}
	symbol, ok := pointerSymbols[pointers]
	if ok {
		return symbol
	}
	return "."
}
