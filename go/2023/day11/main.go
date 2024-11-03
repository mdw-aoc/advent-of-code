package day11

import (
	"slices"

	. "github.com/mdw-go/funcy/ranger"
	"github.com/mdw-go/grid"
	"github.com/mdw-go/set"
)

func Part1(lines []string) (result int) {
	return Part2(lines, 2)
}
func Part2(lines []string, timeFactor int) int {
	return sumDistances(galaxyPairs(expandGalaxies(plotGalaxies(lines), timeFactor)))
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
func galaxyPairs(galaxies []grid.Point[int]) (result []pair[grid.Point[int]]) {
	for g, g1 := range galaxies {
		for _, g2 := range galaxies[g+1:] {
			result = append(result, pair[grid.Point[int]]{A: g1, B: g2})
		}
	}
	return result
}
func sumDistances(pairs []pair[grid.Point[int]]) (result int) {
	for _, pair := range pairs {
		result += grid.CityBlockDistance(pair.A, pair.B)
	}
	return result
}
func expandGalaxies(galaxies []grid.Point[int], timeFactor int) (result []grid.Point[int]) {
	G := Iterator(galaxies)
	X := Map(grid.Point[int].X, G)
	Y := Map(grid.Point[int].Y, G)
	gapsX := slices.Sorted[int](set.FromSeq(Range(0, Max(X))).Difference(set.FromSeq(X)).All())
	gapsY := slices.Sorted[int](set.FromSeq(Range(0, Max(Y))).Difference(set.FromSeq(Y)).All())
	for galaxy := range G {
		dx := 0
		dy := 0
		for _, gapX := range gapsX {
			if gapX < galaxy.X() {
				dx += timeFactor - 1
			}
		}
		for _, gapY := range gapsY {
			if gapY < galaxy.Y() {
				dy += timeFactor - 1
			}
		}
		result = append(result, galaxy.Move(grid.NewDirection(dx, dy)))
	}
	return result
}

type pair[T any] struct{ A, B T }
