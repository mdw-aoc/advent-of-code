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
	this.So(this.Part1(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputLines), should.Equal, TODO)
}
func (this *Suite) TestPart2Samples() {
	this.So(this.Part2(sampleLines), should.Equal, TODO)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputLines), should.Equal, TODO)
}
func (this *Suite) Part1(lines []string) any {
	energized := set.Of[grid.Point[int]]()
	queue := []grid.Point[int]{grid.NewPoint(0, 0)}
	direction := grid.NewDirection[int](1, 0)
	for len(queue) > 0 {

	}
	return TODO
}
func (this *Suite) Part2(lines []string) any {
	return TODO
}
