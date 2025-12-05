package main

import (
	"bufio"
	"os"
	"strconv"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, 357, calculateTotalJoltage("sample-input.txt", 2))
	assertEqual(t, 17332, calculateTotalJoltage("input.txt", 2))
	assertEqual(t, 3121910778619, calculateTotalJoltage("sample-input.txt", 12))
	assertEqual(t, 172516781546707, calculateTotalJoltage("input.txt", 12))
}
func assertEqual(t *testing.T, expected, actual any) {
	t.Log(actual)
	if expected != actual {
		t.Helper()
		t.Errorf("Expected %v, got %v", expected, actual)
	}
}
func calculateTotalJoltage(filename string, batteryCount int) (result int) {
	file, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer func() { _ = file.Close() }()
	for scanner := bufio.NewScanner(file); scanner.Scan(); {
		result += calculateJoltage(scanner.Text(), batteryCount)
	}
	return result
}
func calculateJoltage(line string, batteryCount int) int {
	start, powered := 0, ""
	for remaining := batteryCount; remaining > 0; remaining-- {
		maxI, maxC := 0, '0'
		for i, c := range line[start:(len(line) - remaining + 1)] {
			if c > maxC {
				maxC, maxI = c, i
			}
		}
		start += maxI + 1
		powered += string(maxC)
	}
	return parseInt(powered)
}
func parseInt(s string) int {
	n, _ := strconv.Atoi(s)
	return n
}
