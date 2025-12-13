package main

import (
	"bufio"
	"os"
	"reflect"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, 4277556, completeWorksheet(scanWorksheet("sample-input.txt")))
	assertEqual(t, 3261038365331, completeWorksheet(scanWorksheet("input.txt")))
}
func assertEqual(t *testing.T, expected, actual any) {
	t.Helper()
	if !reflect.DeepEqual(expected, actual) {
		t.Errorf("Expected %v, got %v", expected, actual)
	} else {
		t.Log(actual)
	}
}

func scanWorksheet(filename string) (worksheet [][]int, operations []string) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()
		fields := strings.Fields(line)
		if fields[0] == "+" || fields[0] == "*" {
			operations = fields[:]
		} else if line == "" {
			continue
		} else {
			var record []int
			for _, field := range fields {
				n, _ := strconv.Atoi(field)
				record = append(record, n)
			}
			worksheet = append(worksheet, record)
		}
	}
	return worksheet, operations
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
