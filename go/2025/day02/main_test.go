package main

import (
	"os"
	"strconv"
	"strings"
	"testing"
)

func Test(t *testing.T) {
	assertEqual(t, tallyInvalidNumbersInRanges("sample-input.txt", isValidPart1), 1227775554)
	assertEqual(t, tallyInvalidNumbersInRanges("input.txt", isValidPart1), 23534117921)
	assertEqual(t, tallyInvalidNumbersInRanges("sample-input.txt", isValidPart2), 4174379265)
	assertEqual(t, tallyInvalidNumbersInRanges("input.txt", isValidPart2), 31755323497)
}
func assertEqual(t *testing.T, got, want any) {
	t.Log(got)
	if got != want {
		t.Helper()
		t.Errorf("got %v, want %v", got, want)
	}
}

func tallyInvalidNumbersInRanges(filename string, isValid func(string) bool) (result int) {
	input, _ := os.ReadFile(filename)
	for lohi := range strings.SplitSeq(string(input), ",") {
		if lo, hi, ok := strings.Cut(lohi, "-"); ok {
			LO, _ := strconv.Atoi(lo)
			HI, _ := strconv.Atoi(hi)
			for i := LO; i <= HI; i++ {
				if !isValid(strconv.Itoa(i)) {
					result += i
				}
			}
		}
	}
	return result
}
func isValidPart1(number string) bool {
	if len(number)%2 != 0 {
		return true
	}
	split := len(number) / 2
	half1 := number[:split]
	half2 := number[split:]
	return half1 != half2
}
func isValidPart2(number string) bool {
	for x := 1; x <= len(number)/2; x++ {
		if len(number)%x != 0 {
			continue
		}
		prefix := number[:x]
		if number == strings.Repeat(prefix, len(number)/x) {
			return false
		}
	}
	return true
}
