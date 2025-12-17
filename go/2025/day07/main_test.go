package main

import (
	"bufio"
	"os"
	"reflect"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, 21, part1("sample-input.txt"))
	assertEqual(t, 1717, part1("input.txt"))
}
func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if reflect.DeepEqual(expected, actual) {
		t.Log(actual)
	} else {
		t.Errorf("expected %v, got %v", expected, actual)
	}
}

type Set[T comparable] struct{ m map[T]struct{} }

func NewSet[T comparable]() *Set[T]    { return &Set[T]{m: make(map[T]struct{})} }
func (this *Set[T]) Add(x T)           { this.m[x] = struct{}{} }
func (this *Set[T]) Remove(x T)        { delete(this.m, x) }
func (this *Set[T]) Contains(x T) bool { _, ok := this.m[x]; return ok }
func (this *Set[T]) Len() int          { return len(this.m) }

type Point struct{ x, y int }

func (this Point) Left() Point  { return Point{x: this.x - 1, y: this.y} }
func (this Point) Right() Point { return Point{x: this.x + 1, y: this.y} }
func (this Point) Down() Point  { return Point{x: this.x, y: this.y + 1} }

func part1(filename string) int {
	return countSplits(parseField(filename))
}
func parseField(filename string) (start Point, splitters *Set[Point]) {
	splitters = NewSet[Point]()
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for row := 1; scanner.Scan(); row++ {
		line := scanner.Text()
		for col, char := range line {
			if char == 'S' {
				start = Point{x: col, y: row}
			}
			if char == '^' {
				splitters.Add(Point{x: col, y: row})
			}
		}
	}
	return start, splitters
}
func countSplits(start Point, splitters *Set[Point]) (result int) {
	streams := NewSet[Point]()
	streams.Add(start.Down())
	for row := 0; row < 150; row++ {
		next := NewSet[Point]()
		for stream := range streams.m {
			if splitters.Contains(stream) {
				result++
				next.Add(stream.Left().Down())
				next.Add(stream.Right().Down())
			} else {
				next.Add(stream.Down())
			}
		}
		streams = next
	}
	return result
}
