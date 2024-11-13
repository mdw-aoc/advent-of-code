package starter

import (
	"log"
	"slices"
	"strings"
	"testing"

	"github.com/mdw-aoc/inputs"
	_ "github.com/mdw-go/funcy"
	_ "github.com/mdw-go/must"
	_ "github.com/mdw-go/set"
	"github.com/mdw-go/testing/should"
)

var (
	samplePatternA = []string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
	}
	samplePatternB = []string{
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}
	fullSample = strings.Join([]string{
		"#.##..##.",
		"..#.##.#.",
		"##......#",
		"##......#",
		"..#.##.#.",
		"..##..##.",
		"#.#.##.#.",
		"", ////////
		"#...##..#",
		"#....#..#",
		"..##..###",
		"#####.##.",
		"#####.##.",
		"..##..###",
		"#....#..#",
	}, "\n")
	samplePatternC = []string{
		"####..#.#####",
		".#.#.#.#.##..", // <- Smudge is second to last '.' (should be '#')
		"..#.##....#.#",
		"#...#..##....",
		"###.####.#.#.",
		"..#....#.#..#",
		"..#....#.#..#",
		"###.####.#.#.",
		"#...#..##....",
		"..#.##....#.#",
		".#.#.#.#.###.",
		"####..#.#####",
		"####..#.#####",
	}
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) TestRotate() {
	this.So(Rotate(Rotate(Rotate(Rotate(samplePatternA)))), should.Equal, samplePatternA)
	this.So(Rotate(Rotate(Rotate(Rotate(samplePatternB)))), should.Equal, samplePatternB)
}
func (this *Suite) TestPart1Samples() {
	this.So(Reflect(samplePatternB), should.Equal, 4)
	this.So(Reflect(Rotate(samplePatternA)), should.Equal, 5)
	this.So(Reflect(samplePatternC), should.Equal, 12)
}
func (this *Suite) TestPart1Full() {
	this.So(Part1(strings.TrimSpace(fullSample)), should.Equal, 405)
	this.So(Part1(inputs.Read(2023, 13).String()), should.Equal, 31265)
}
func (this *Suite) TestPart2Samples() {
}
func (this *Suite) TestPart2Full() {
}

func Part1(input string) int {
	var ABOVE int
	var LEFT int
	patterns := strings.Split(input, "\n\n")
	for p, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		lines := strings.Split(pattern, "\n")
		above := Reflect(lines)
		if above > 0 {
			ABOVE += above
			continue
		}
		left := Reflect(Rotate(lines))
		if left > 0 {
			LEFT += left
			continue
		}
		log.Panicf("wat %d", p)
	}
	return ABOVE*100 + LEFT
}
func Rotate(lines []string) (columns []string) {
	for x := 0; x < len(lines[0]); x++ {
		columns = append(columns, "")
		for y := 0; y < len(lines); y++ {
			columns[x] += string(lines[y][x])
		}
	}
	return columns
}
func Reflect(lines []string) int {
	for x := 1; x < len(lines); x++ {
		before := make([]string, x)
		copy(before, lines[:x])
		slices.Reverse(before)
		after := lines[x:]
		if len(before) > len(after) {
			before = before[:len(after)]
		}
		if len(after) > len(before) {
			after = after[:len(before)]
		}
		if slices.Equal(before, after) {
			return x
		}
	}

	return 0
}
func Part2(lines []string) any {
	return -1
}
