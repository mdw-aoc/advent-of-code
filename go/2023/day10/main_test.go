package day10

import (
	"fmt"
	"slices"
	"testing"

	"github.com/mdw-aoc/inputs"
	"github.com/mdw-go/funcy/ranger/is"
	"github.com/mdw-go/set"
	"github.com/mdw-go/testing/should"
)

var (
	inputLines = slices.Collect(inputs.Read(2023, 10).Lines())
	sampleA    = []string{ // 1 enclosed tile
		"..F7.",
		".FJ|.",
		"SJ.L7",
		"|F--J",
		"LJ...",
	}
	sampleB = []string{ // 4 enclosed tiles
		"...........",
		".S-------7.",
		".|F-----7|.",
		".||.....||.",
		".||.....||.",
		".|L-7.F-J|.",
		".|..|.|..|.",
		".L--J.L--J.",
		"...........",
	}
	sampleC = []string{ // 4 enclosed tiles
		"..........",
		".S------7.",
		".|F----7|.",
		".||....||.",
		".||....||.",
		".|L-7F-J|.",
		".|II||II|.",
		".L--JL--J.",
		"..........",
	}
	sampleD = []string{ // 8 enclosed tiles
		".F----7F7F7F7F-7....",
		".|F--7||||||||FJ....",
		".||.FJ||||||||L7....",
		"FJL7L7LJLJ||LJ.L-7..",
		"L--J.L7...LJS7F-7L7.",
		"....F-J..F7FJ|L7L7L7",
		"....L7.F7||L7|.L7L7|",
		".....|FJLJ|FJ|F7|.LJ",
		"....FJL-7.||.||||...",
		"....L---J.LJ.LJLJ...",
	}
	sampleE = []string{ // 10 enclosed tiles
		"FF7FSF7F7F7F7F7F---7",
		"L|LJ||||||||||||F--J",
		"FL-7LJLJ||||||LJL-77",
		"F--JF--7||LJLJ7F7FJ-",
		"L---JF-JLJ.||-FJLJJ7",
		"|F|F-JF---7F7-L7L|7|",
		"|FFJF7L7F-JF7|JL---7",
		"7-L-JL7||F7|L7F-7F7|",
		"L.L7LFJ|||||FJL7||LJ",
		"L7JLJL-JLJLJL--JLJ.L",
	}
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct{ *should.T }

func (this *Suite) TestPart1Samples() {
	this.So(this.Part1(sampleA), should.Equal, 8)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputLines), should.Equal, 6886)
}
func (this *Suite) TestPart2Samples() {
	this.So(this.Part2(sampleA), should.Equal, 1)
	this.So(this.Part2(sampleB), should.Equal, 4)
	this.So(this.Part2(sampleC), should.Equal, 4)
	this.So(this.Part2(sampleD), should.Equal, 8)
	this.So(this.Part2(sampleE), should.Equal, 10)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, 371)
}
func (this *Suite) Part1(lines []string) any {
	return len(ParseInput(lines).circuit) / 2
}
func (this *Suite) Part2(lines []string) int {
	world := ParseInput(lines)
	return len(EnclosedPoints(world.field, world.circuit))
}

type Point struct{ row, col int }

type World struct {
	start   Point
	circuit map[Point]string
	field   map[Point]string
}

func ParseInput(lines []string) World {
	result := World{
		circuit: make(map[Point]string),
		field:   make(map[Point]string),
	}
	for row, line := range lines {
		for col, char := range line {
			result.field[Point{row: row, col: col}] = string(char)
		}
	}
	for point, char := range result.field {
		if char == "S" {
			result.start = point
			result.field[point] = inferStartingS(result.field, point)
		}
	}
	result.circuit = Circuit(result.start, result.field)
	for point := range result.field {
		if _, ok := result.circuit[point]; !ok {
			result.field[point] = "."
		}
	}

	return result
}

func neighbors(p Point) (n, s, e, w Point) {
	n = Point{row: p.row - 1, col: p.col}
	s = Point{row: p.row + 1, col: p.col}
	e = Point{row: p.row, col: p.col + 1}
	w = Point{row: p.row, col: p.col - 1}
	return n, s, e, w
}
func inferStartingS(field map[Point]string, p Point) string {
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

func Circuit(start Point, field map[Point]string) map[Point]string {
	queue := []Point{start}
	frontier := map[Point]string{start: field[start]}
	for {
		at := queue[0]
		queue = queue[1:]
		a, b := follow(at, field)
		if _, ok := frontier[a]; !ok {
			frontier[a] = field[a]
			queue = append(queue, a)
			continue
		}
		if _, ok := frontier[b]; !ok {
			frontier[b] = field[b]
			queue = append(queue, b)
			continue
		}
		if a == start || b == start {
			return frontier
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

func EnclosedPoints(field, circuit map[Point]string) (result []Point) {
	for point := range field {
		if _, ok := circuit[point]; ok {
			continue
		}
		if isEnclosed(point, circuit) {
			result = append(result, point)
		}
	}
	return result
}

// isEnclosed walks east from the starting point, counting how many times it crosses the circuit.
// An odd number of crossings (as defined here) means that the starting point is actually
// outside the circuit.
// An even number of crossings indicates that the starting point is inside the circuit.
// 'Crossing' means we have encountered a '|', or a 'L' or a 'J'.
// If both an 'L' or 'J' are encountered we end up with an even number of crossings because
// those corner joints equate to a 'U'-bend that can be circumvented: not a crossing.
// If, say, an 'L' and an '7' are encountered, those corner joints connect pipes above and
// below the starting point. We'll end up only count the 'L' by this code, so resulting in
// an odd number and therefore a crossing!
// Thanks to this video for the tips: https://youtu.be/edVSG8Y_qf8?t=610
func isEnclosed(point Point, circuit map[Point]string) bool {
	crossings := 0
	for col := point.col; col >= 0; col-- {
		step := Point{point.row, col}
		if c, ok := circuit[step]; ok {
			if crossChecks.Contains(c) {
				crossings++
			}
		}
	}
	return is.Odd(crossings)
}

var crossChecks = set.Of("|", "L", "J")
