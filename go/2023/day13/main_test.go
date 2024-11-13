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

const TODO = -1

var (
	inputLines     = slices.Collect(inputs.Read(2023, 13).Lines())
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
	samplePatternC = []string{
		"####..#.#####",
		".#.#.#.#.##..",
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

func (this *Suite) TestPart1Samples() {
	this.So(Rotate(Rotate(Rotate(Rotate(samplePatternA)))), should.Equal, samplePatternA)
	this.So(Rotate(Rotate(Rotate(Rotate(samplePatternB)))), should.Equal, samplePatternB)

	this.So(Reflect(samplePatternB), should.Equal, 4)
	this.So(Reflect(Rotate(samplePatternA)), should.Equal, 5)
	this.So(Reflect(samplePatternC), should.Equal, 12)

	var sample strings.Builder
	for _, line := range samplePatternA {
		sample.WriteString(line)
		sample.WriteString("\n")
	}
	sample.WriteString("\n")
	for _, line := range samplePatternB {
		sample.WriteString(line)
		sample.WriteString("\n")
	}
	this.So(this.Part1(strings.TrimSpace(sample.String())), should.Equal, 405)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputs.Read(2023, 13).String()), should.Equal, 31265)
}
func (this *Suite) TestPart2Samples() {
}
func (this *Suite) TestPart2Full() {
}
func (this *Suite) Part1(input string) int {
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
func (this *Suite) Part2(lines []string) any {
	return TODO
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
