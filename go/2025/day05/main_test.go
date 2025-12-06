package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, 3, countFreshIngredients("sample-input.txt"))
	assertEqual(t, 598, countFreshIngredients("input.txt"))
}
func assertEqual(t *testing.T, expected, actual any) {
	t.Log(actual)
	if expected != actual {
		t.Helper()
		t.Errorf("Got: %v (Want: %v)", actual, expected)
	}
}
func countFreshIngredients(filename string) (result int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	var fresh []int
	for scanner.Scan() {
		line := scanner.Text()
		if line == "" {
			break
		}
		rawLo, rawHi, ok := strings.Cut(line, "-")
		if !ok {
			panic("invalid line: " + line)
		}
		lo, _ := strconv.Atoi(rawLo)
		hi, _ := strconv.Atoi(rawHi)
		fresh = append(fresh, lo, hi)
	}
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		for x := 0; x < len(fresh); x += 2 {
			if fresh[x] <= n && n <= fresh[x+1] {
				result++
				break
			}
		}
	}
	return result
}
