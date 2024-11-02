package day11

import (
	"slices"
	"testing"

	"github.com/mdw-aoc/inputs/v2/inputs"
	_ "github.com/mdw-go/funcy"
	_ "github.com/mdw-go/must/must"
	_ "github.com/mdw-go/set/v2/set"
	"github.com/mdw-go/testing/should"
)

const TODO = -1

var (
	inputLines  = slices.Collect(inputs.Read(2023, 11).Lines())
	sampleLines = []string{
		"...#......",
		".......#..",
		"#.........",
		"..........",
		"......#...",
		".#........",
		".........#",
		"..........",
		".......#..",
		"#...#.....",
	}
	expendedSampleLines = []string{
		"....#........",
		".........#...",
		"#............",
		".............",
		".............",
		"........#....",
		".#...........",
		"............#",
		".............",
		".............",
		".........#...",
		"#....#.......",
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
	this.So(this.Part1(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputLines), should.Equal, TODO)
}
func (this *Suite) TestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) any {
	return TODO
}

func (this *Suite) Part2(lines []string) any {
	return TODO
}
