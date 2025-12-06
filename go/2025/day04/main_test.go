package main

import (
	"bufio"
	"iter"
	"os"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, 13, len(accessibleRolls(parseWorld("sample-input.txt"))))
	assertEqual(t, 1533, len(accessibleRolls(parseWorld("input.txt"))))
}
func assertEqual(t *testing.T, expected, actual any) {
	t.Log(actual)
	if expected != actual {
		t.Helper()
		t.Errorf("Expected %v (type %v) - Got %v (type %v)", expected, expected, actual, actual)
	}
}

func accessibleRolls(world World) (result []Location) {
	for loc := range world {
		if !world.isOccupied(loc) {
			continue
		}
		count := 0
		for neighbor := range loc.neighbors8() {
			if world.isOccupied(neighbor) {
				count++
			}
		}
		if count < 4 {
			result = append(result, loc)
		}
	}
	return result
}

func parseWorld(filename string) World {
	result := make(World)
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for row := 0; scanner.Scan(); row++ {
		line := scanner.Text()
		for col, c := range line {
			result[Location{row: row, col: col}] = c == '@'
		}
	}
	return result
}

type Location struct {
	row, col int
}

func (this Location) neighbors8() iter.Seq[Location] {
	return func(yield func(Location) bool) {
		for r := this.row - 1; r <= this.row+1; r++ {
			for c := this.col - 1; c <= this.col+1; c++ {
				if r == this.row && c == this.col {
					continue
				}
				if !yield(Location{row: r, col: c}) {
					return
				}
			}
		}
	}
}

type World map[Location]bool

func (this World) isInside(loc Location) bool {
	_, ok := this[loc]
	return ok
}
func (this World) isOccupied(loc Location) bool {
	return this.isInside(loc) && this[loc]
}
