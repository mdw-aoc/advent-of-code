package day11

import (
	"log"
	"slices"
	"testing"

	"github.com/mdw-aoc/inputs/v2/inputs"
	"github.com/mdw-go/testing/should"
)

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
	expandedUniverse = []string{
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
	should.Run(&Suite{T: should.New(t)}, should.Options.IntegrationTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) Setup() {
	log.SetOutput(this)
}
func (this *Suite) TestPart1A() {
	this.So(Part1(sampleLines), should.Equal, 374)
}
func (this *Suite) TestPart1Full() {
	this.So(Part1(inputLines), should.Equal, 9799681)
}
func (this *Suite) TestPart2A() {
	this.So(Part2(sampleLines, 10), should.Equal, 1030)
	this.So(Part2(sampleLines, 100), should.Equal, 8410)
}
func (this *Suite) TestPart2Full() {
	this.So(Part2(inputLines, 1_000_000), should.Equal, 513171773355)
}
