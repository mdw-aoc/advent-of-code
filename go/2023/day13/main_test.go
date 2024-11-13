package starter

import (
	"slices"
	"strings"
	"testing"

	"github.com/mdw-aoc/inputs"
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
	this.So(Reflect(samplePatternB, 0), should.Equal, 4)
	this.So(Reflect(Rotate(samplePatternA), 0), should.Equal, 5)
	this.So(Reflect(samplePatternC, 0), should.Equal, 12)
}
func (this *Suite) TestPart1Full() {
	this.So(Part1(strings.TrimSpace(fullSample)), should.Equal, 405)
	this.So(Part1(inputs.Read(2023, 13).String()), should.Equal, 31265)
}
func (this *Suite) TestPart2Samples() {
	this.So(Reflect(samplePatternA, 1), should.Equal, 3)
	this.So(Reflect(samplePatternB, 1), should.Equal, 1)
}
func (this *Suite) TestPart2Full() {
	this.So(Part2(strings.TrimSpace(fullSample)), should.Equal, 400)
	this.So(Part2(inputs.Read(2023, 13).String()), should.Equal, 39359)
}

func Part1(input string) int { return SummarizePatterns(input, 0) }
func Part2(input string) int { return SummarizePatterns(input, 1) }
func SummarizePatterns(input string, diffTarget int) int {
	var above, left int
	for _, pattern := range strings.Split(input, "\n\n") {
		lines := strings.Split(pattern, "\n")
		above += Reflect(lines, diffTarget)
		left += Reflect(Rotate(lines), diffTarget)
	}
	return above*100 + left
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
func Reflect(lines []string, diffTarget int) int {
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
		if Diff(before, after) == diffTarget {
			return x
		}
	}
	return 0
}
func Diff(s1, s2 []string) (result int) {
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
