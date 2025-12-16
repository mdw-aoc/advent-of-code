package main

import (
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	samplePart1, samplePart2 := solve("sample-input.txt")
	realPart1, realPart2 := solve("input.txt")

	assertEqual(t, 4277556, samplePart1)
	assertEqual(t, 3261038365331, realPart1)

	assertEqual(t, 3263827, samplePart2)
	assertEqual(t, 8342588849093, realPart2)
}
func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	} else {
		t.Log(actual)
	}
}

func solve(filename string) (part1, part2 int) {
	worksheet1, worksheet2, operations := scanWorksheet(filename)
	part1 = completeWorksheet(worksheet1, operations)
	part2 = completeWorksheet(worksheet2, operations)
	return part1, part2
}
func scanWorksheet(filename string) (worksheet1, worksheet2 [][]int, operations []string) {
	content, err := os.ReadFile(filename)
	if err != nil {
		panic(err)
	}
	lines := strings.Split(string(content), "\n")
	for l := range lines {
		lines[l] = lines[l] + "    " // pad lines so we consistently detect whitespace at the end of each line.
	}
	rawOps := lines[len(lines)-1]
	lines = lines[:len(lines)-1]
	var verticals = make([]string, len(lines))
	var horizontals = make([]string, 4)
	var verticalOffset int
	for c := range rawOps {
		if allDigitsBlank(c, lines) { // + or *
			var vertical []int
			var horizontal []int
			for _, raw := range verticals {
				parsed, _ := strconv.Atoi(strings.TrimSpace(raw))
				if parsed != 0 {
					vertical = append(vertical, parsed)
				}
			}
			for _, raw := range horizontals {
				parsed, _ := strconv.Atoi(strings.TrimSpace(raw))
				if parsed != 0 {
					horizontal = append(horizontal, parsed)
				}
			}
			worksheet1 = append(worksheet1, horizontal)
			worksheet2 = append(worksheet2, vertical)
			verticals = make([]string, len(lines))
			horizontals = make([]string, 4)
			verticalOffset = 0
			continue
		}
		for l, line := range lines {
			verticals[verticalOffset] += string(line[c])
			horizontals[l] += string(line[c])
		}
		verticalOffset++
	}
	operations = strings.Fields(rawOps)
	return worksheet1, worksheet2, operations
}
func allDigitsBlank(c int, lines []string) bool {
	for _, line := range lines {
		if line[c] != ' ' {
			return false
		}
	}
	return true
}
func completeWorksheet(worksheet [][]int, operations []string) (total int) {
	for r, row := range worksheet {
		if row != nil {
			rowTotal := reduce(operations[r], row...)
			total += rowTotal
		}
	}
	return total
}
func reduce(op string, values ...int) (result int) {
	var operation func(a, b int) int
	if op == "*" {
		operation, result = mul, 1
	} else {
		operation, result = add, 0
	}
	for _, value := range values {
		result = operation(result, value)
	}
	return result
}
func add(a, b int) int { return a + b }
func mul(a, b int) int { return a * b }
