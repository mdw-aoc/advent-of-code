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
	this.So(Reflect(samplePatternB, Equal), should.Equal, 4)
	this.So(Reflect(Rotate(samplePatternA), Equal), should.Equal, 5)
	this.So(Reflect(samplePatternC, Equal), should.Equal, 12)
}
func (this *Suite) TestPart1Full() {
	this.So(Part1(strings.TrimSpace(fullSample)), should.Equal, 405)
	this.So(Part1(inputs.Read(2023, 13).String()), should.Equal, 31265)
}
func (this *Suite) TestPart2Samples() {
	this.So(Reflect(samplePatternA, EqualSmudged), should.Equal, 3)
	this.So(Reflect(samplePatternB, EqualSmudged), should.Equal, 1)
}
func (this *Suite) TestPart2Full() {
	this.So(Part2(strings.TrimSpace(fullSample)), should.Equal, 400)
	this.So(Part2(inputs.Read(2023, 13).String()), should.Equal, 39359)
}

func Part1(input string) int {
	return SummarizePatterns(input, Equal)
}
func Part2(input string) any {
	return SummarizePatterns(input, EqualSmudged)
}
func SummarizePatterns(input string, equal func(a, b []string) bool) int {
	var ABOVE int
	var LEFT int
	patterns := strings.Split(input, "\n\n")
	for p, pattern := range patterns {
		pattern = strings.TrimSpace(pattern)
		lines := strings.Split(pattern, "\n")
		above := Reflect(lines, equal)
		if above > 0 {
			ABOVE += above
			continue
		}
		left := Reflect(Rotate(lines), equal)
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
func Reflect(lines []string, equal func(a, b []string) bool) int {
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
		if equal(before, after) {
			return x
		}
	}
	return 0
}
func Equal(a, b []string) bool                 { return slices.Equal(a, b) }
func EqualSmudged(before, after []string) bool { return Diff(before, after) == 1 }
func Diff(s1, s2 []string) (result int) {
	if len(s1) != len(s2) {
		panic("slices must have equal length")
	}
	for i := range s1 {
		a, b := s1[i], s2[i]
		for c := range a {
			if a[c] != b[c] {
				result++
			}
		}
	}
	return result
}
