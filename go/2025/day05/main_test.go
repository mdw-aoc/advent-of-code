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
	ranges := scanRanges(scanner)
	for scanner.Scan() {
		n, _ := strconv.Atoi(scanner.Text())
		for _, r := range ranges { // a binary search here would be more efficient
			if r.Lo <= n && n <= r.Hi {
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
	scanner := bufio.NewScanner(file)
	var at int
	for _, r := range scanRanges(scanner) {
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

func scanRanges(scanner *bufio.Scanner) (results []Range) {
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
		results = append(results, Range{
			Lo: min(lo, hi),
			Hi: max(lo, hi),
		})
	}
	return slices.Compact(slices.SortedFunc(slices.Values(results), func(a, b Range) int {
		return cmp.Compare(a.Lo, b.Lo)
	}))
}

type Range struct {
	Lo, Hi int
}
