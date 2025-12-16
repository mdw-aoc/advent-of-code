package main

import (
	"os"
	"reflect"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	samplePart1, samplePart2 := solve("sample-input.txt")
	realPart1, realPart2 := solve("input.txt")

	assertEqual(t, 4277556, samplePart1)
	assertEqual(t, 3261038365331, realPart1)

	assertEqual(t, 0, samplePart2)
	assertEqual(t, 0, realPart2)
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
	rawOps := lines[len(lines)-1]
	_ = rawOps
	for c := range len(lines[0]) {
		_ = c // TODO build both vertical and horizontal numbers
	}
	operations = strings.Fields(lines[len(lines)-1])
	return worksheet1, worksheet2, operations
}
func completeWorksheet(worksheet [][]int, operations []string) (total int) {
	for c := range len(worksheet[0]) {
		op := operations[c]
		var values []int
		for r := range worksheet {
			values = append(values, worksheet[r][c])
		}
		total += reduce(op, values...)
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
