package starter

import (
	"slices"
	"testing"

	"github.com/mdw-aoc/inputs"
	_ "github.com/mdw-go/funcy"
	_ "github.com/mdw-go/must"
	"github.com/mdw-go/set"
	_ "github.com/mdw-go/set"
	"github.com/mdw-go/testing/should"
)

const TODO = -1

var (
	inputLines  = slices.Collect(inputs.Read(2023, 14).Lines())
	sampleLines = []string{
		"O....#....",
		"O.OO#....#",
		".....##...",
		"OO.#O....O",
		".O.....O#.",
		"O.#..O.#.#",
		"..O..#O..O",
		".......O..",
		"#....###..",
		"#OO..#....",
	}
	sampleLinesAfterNorthTilt = []string{
		"OOOO.#.O..",
		"OO..#....#",
		"OO..O##..O",
		"O..#.OO...",
		"........#.",
		"..#....#.#",
		"..O..#.O.O",
		"..O.......",
		"#....###..",
		"#....#....",
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
	this.So(this.Part1(sampleLines), should.Equal, 136)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputLines), should.Equal, 106648)
}
func (this *Suite) TestPart2Samples() {
	this.So(this.Part2(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) any {
	return ParseInput(lines).Tilt(north).NorthLoad()
}
func (this *Suite) Part2(lines []string) any {
	return TODO
}

type Point struct{ row, col int }
type Direction struct{ dRow, dCol int }

var (
	north = Direction{dRow: -1}
	south = Direction{dRow: 1}
	east  = Direction{dCol: -1}
	west  = Direction{dCol: 1}
)

type World struct {
	rowCount int
	colCount int
	field    map[Point]string
	spheres  set.Set[Point]
	cubes    set.Set[Point]
}

func (this World) Tilt(direction Direction) World {
	floors := map[int]int{}
	for x := 0; x < this.colCount; x++ {
		floors[x] = -1
	}
	for row := 0; row < this.rowCount; row++ {
		for col := 0; col < this.colCount; col++ {
			p := Point{row, col}
			if this.cubes.Contains(p) {
				floors[col] = row
			} else if this.spheres.Contains(p) {
				this.spheres.Remove(p)
				floors[col]++
				p.row = floors[col]
				this.spheres.Add(p)
			}
		}
	}
	return this
}

func (this World) NorthLoad() (result int) {
	for point := range this.spheres {
		result += this.rowCount - point.row
	}
	return result
}

func ParseInput(lines []string) (result World) {
	result.field = make(map[Point]string)
	result.spheres = set.Of[Point]()
	result.cubes = set.Of[Point]()
	result.colCount = len(lines[0])
	for row, line := range lines {
		result.rowCount++
		for col, char := range line {
			point := Point{row, col}
			result.field[point] = string(char)
			if char == 'O' {
				result.spheres.Add(point)
			} else if char == '#' {
				result.cubes.Add(point)
			}
		}
	}
	return result
}
