package main

import (
	"bufio"
	"cmp"
	"os"
	"slices"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, 3, countFreshIngredients("sample-input.txt"))
	assertEqual(t, 598, countFreshIngredients("input.txt"))
	assertEqual(t, 14, countFreshInventory("sample-input.txt"))
	assertEqual(t, 360341832208407, countFreshInventory("input.txt"))
}

func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if expected == actual {
		t.Log(actual)
	} else {
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
func countFreshInventory(filename string) (total int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	var ranges []Range
	scanner := bufio.NewScanner(file)
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
		ranges = append(ranges, Range{
			Lo: min(lo, hi),
			Hi: max(lo, hi),
		})
	}
	ranges = slices.Compact(slices.SortedFunc(slices.Values(ranges), func(a, b Range) int {
		return cmp.Compare(a.Lo, b.Lo)
	}))
	var at int
	for _, r := range ranges {
		var lo int
		if at <= r.Lo { // jump over gap
			lo = r.Lo
		} else if r.Hi <= at { // already counted
			continue
		} else { // account for overlap
			lo = at
		}
		total += r.Hi - lo + 1
		at = r.Hi + 1
	}
	return total
}

type Range struct {
	Lo, Hi int
}
