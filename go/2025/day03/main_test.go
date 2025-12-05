package main

import (
	"bufio"
	"os"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, 357, calculateTotalJoltage("sample-input.txt"))
	assertEqual(t, 17332, calculateTotalJoltage("input.txt"))
}
func assertEqual(t *testing.T, expected, actual any) {
	t.Log(actual)
	if expected != actual {
		t.Helper()
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}

func calculateTotalJoltage(filename string) (result int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		var hi1, hi2 rune
		var i int
		line := scanner.Text()
		for c := range line[:len(line)-1] {
			r := rune(line[c])
			if r > hi1 {
				hi1 = r
				i = c
			}
		}
		for c := range line {
			if c <= i {
				continue
			}
			r := rune(line[c])
			if r > hi2 {
				hi2 = r
			}
		}
		result += joltage(hi1, hi2)
	}
	return result
}

func joltage(a, b rune) int {
	return (10 * (int(a) - '0')) + (int(b) - '0')
}
