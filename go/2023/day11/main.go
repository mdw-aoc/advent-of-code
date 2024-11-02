package day11

import (
	"strings"

	"github.com/mdw-go/grid"
)

func expandUniverse(lines []string) (result []string) {
	lines = expandRows(lines)
	lines = rotateRows(lines) // 90
	lines = expandRows(lines)
	lines = rotateRows(lines) // 180
	lines = rotateRows(lines) // 270
	lines = rotateRows(lines) // 360
	return lines
}
func expandRows(lines []string) (result []string) {
	for _, line := range lines {
		if !strings.Contains(line, "#") {
			result = append(result, strings.Repeat(".", len(line)))
		}
		result = append(result, line)
	}
	return result
}
func rotateRows(lines []string) (result []string) {
	for x := 0; x < len(lines[0]); x++ {
		result = append(result, "")
		for _, line := range lines {
			result[x] += line[x : x+1]
		}
	}
	return result
}
func plotGalaxies(lines []string) (result []grid.Point[int]) {
	for y, line := range lines {
		for x, char := range line {
			if char == '#' {
				result = append(result, grid.NewPoint(x, y))
			}
		}
	}
	return result
}
func galaxyPairs(galaxies []grid.Point[int]) (result []Pair[grid.Point[int]]) {
	for g, g1 := range galaxies {
		for _, g2 := range galaxies[g+1:] {
			result = append(result, Pair[grid.Point[int]]{A: g1, B: g2})
		}
	}
	return result
}

type Pair[T any] struct{ A, B T }
