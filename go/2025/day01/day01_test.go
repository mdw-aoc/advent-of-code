package main

import (
	"bufio"
	"iter"
	"os"
	"strconv"
	"testing"
)

func YieldInstructions(filename string) iter.Seq[int] {
	return func(yield func(int) bool) {
		file, err := os.Open(filename)
		if err != nil {
			panic(err)
		}
		defer func() { _ = file.Close() }()
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			line := scanner.Text()
			if len(line) < 2 {
				continue
			}
			n, err := strconv.Atoi(line[1:])
			if err != nil {
				panic(err)
			}
			if line[0] == 'L' {
				n = -n
			}
			if !yield(n) {
				return
			}
		}
	}
}

func TestPart1Sample(t *testing.T) {
	expected := 3
	actual := countZerosAfterEachInstruction("day01-sample.txt")
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
func TestPart1(t *testing.T) {
	expected := 1102
	actual := countZerosAfterEachInstruction("day01.txt")
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func countZerosAfterEachInstruction(filename string) int {
	zeros := 0
	at := 50
	for n := range YieldInstructions(filename) {
		at += n
		at %= 100
		if at == 0 {
			zeros++
		}
	}
	return zeros
}

func TestPart2Sample(t *testing.T) {
	expected := 6
	actual := countAllZeros("day01-sample.txt")
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}
func TestPart2(t *testing.T) {
	expected := 6175
	actual := countAllZeros("day01.txt")
	if actual != expected {
		t.Errorf("expected %d, got %d", expected, actual)
	}
}

func countAllZeros(filename string) int {
	zeros := 0
	at := 50
	for n := range YieldInstructions(filename) {
		for n > 0 {
			n--
			at++
			if at > 99 {
				at = 0
				zeros++
			}
		}
		for n < 0 {
			n++
			at--
			if at == 0 {
				zeros++
			}
			if at < 0 {
				at = 99
			}
		}
	}
	return zeros
}
