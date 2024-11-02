package day11

import (
	"slices"
	"testing"

	"github.com/mdw-aoc/inputs/v2/inputs"
	_ "github.com/mdw-go/funcy"
	"github.com/mdw-go/grid"
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
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) Setup() {
}

func (this *Suite) TestPart1A() {
	this.So(expandUniverse(sampleLines), should.Equal, expandedUniverse)
	this.So(plotGalaxies(expandedUniverse), should.Equal, []grid.Point[int]{
		grid.NewPoint(4, 0),
		grid.NewPoint(9, 1),
		grid.NewPoint(0, 2),
		grid.NewPoint(8, 5),
		grid.NewPoint(1, 6),
		grid.NewPoint(12, 7),
		grid.NewPoint(9, 10),
		grid.NewPoint(0, 11),
		grid.NewPoint(5, 11),
	})
	this.So(
		galaxyPairs([]grid.Point[int]{
			grid.NewPoint(4, 0),
			grid.NewPoint(9, 1),
			grid.NewPoint(0, 2),
		}),
		should.Equal,
		[]Pair[grid.Point[int]]{
			{A: grid.NewPoint(4, 0), B: grid.NewPoint(9, 1)},
			{A: grid.NewPoint(4, 0), B: grid.NewPoint(0, 2)},
			{A: grid.NewPoint(9, 1), B: grid.NewPoint(0, 2)},
		},
	)
	this.So(this.Part1(sampleLines), should.Equal, 374)
}

func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputLines), should.Equal, 9799681)
}
func (this *Suite) TestPart2A() {
	this.So(this.Part2(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) (result int) {
	for _, pair := range galaxyPairs(plotGalaxies(expandUniverse(lines))) {
		result += grid.CityBlockDistance(pair.A, pair.B)
	}
	return result
}

func (this *Suite) Part2(lines []string) any {
	return TODO
}
