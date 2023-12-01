package day01

import (
	"fmt"
	"unicode"

	"github.com/mdwhatcott/advent-of-code-inputs/inputs"
)

// This code was almost completely generated by chat-gpt (3.5)
// using the puzzle input (almost verbatim) as input.
// https://chat.openai.com/c/42fc3620-568f-441a-b298-7ef1bc3e3583
// It nailed part 1 on the first try but failed with part 2 (produces '0' as the result).

func extractCalibrationValue(line string) int {
	var firstDigit, lastDigit rune

	// Find the first digit
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			firstDigit = ch
			break
		}
	}

	// Find the last digit
	for i := len(line) - 1; i >= 0; i-- {
		ch := line[i]
		if unicode.IsDigit(rune(ch)) {
			lastDigit = rune(ch)
			break
		}
	}

	// Combine and convert to integer
	value := int(firstDigit-'0')*10 + int(lastDigit-'0')
	return value
}

func extractCalibrationValueWithWords(line string) int {
	var firstDigit, lastDigit string

	// Find the first digit or spelled-out digit
	for _, ch := range line {
		if unicode.IsDigit(ch) {
			firstDigit = string(ch)
			break
		} else if unicode.IsLetter(ch) {
			firstDigit += string(ch)
		}
	}

	// Find the last digit or spelled-out digit
	for i := len(line) - 1; i >= 0; i-- {
		ch := line[i]
		if unicode.IsDigit(rune(ch)) {
			lastDigit = string(ch)
			break
		} else if unicode.IsLetter(rune(ch)) {
			lastDigit = string(ch) + lastDigit
		}
	}

	// Convert spelled-out digits to numerical digits
	firstDigit = convertWordToDigit(firstDigit)
	lastDigit = convertWordToDigit(lastDigit)

	// Combine and convert to integer
	value := toNumber(firstDigit)*10 + toNumber(lastDigit)
	return value
}

func convertWordToDigit(word string) string {
	switch word {
	case "one":
		return "1"
	case "two":
		return "2"
	case "three":
		return "3"
	case "four":
		return "4"
	case "five":
		return "5"
	case "six":
		return "6"
	case "seven":
		return "7"
	case "eight":
		return "8"
	case "nine":
		return "9"
	default:
		return word
	}
}

func toNumber(s string) int {
	num, err := fmt.Sscanf(s, "%d")
	if err != nil {
		return 0
	}
	return num
}

func main() {
	// Example document
	document1 := []string{
		"1abc2",
		"pqr3stu8vwx",
		"a1b2c3d4e5f",
		"treb7uchet",
	}
	document2 := []string{
		"two1nine",
		"eightwothree",
		"abcone2threexyz",
		"xtwone3four",
		"4nineeightseven2",
		"zoneight234",
		"7pqrstsixteen",
	}

	fmt.Println("Total Sum of Calibration Values:", part1(document1))
	fmt.Println("Total Sum of Calibration Values:", part1(inputs.Read(2023, 1).Lines()))

	fmt.Println("Total Sum of Calibration Values:", part2(document2))
	fmt.Println("Total Sum of Calibration Values:", part2(inputs.Read(2023, 1).Lines()))
}

func part1(document []string) int {
	totalSum := 0
	for _, line := range document {
		calibrationValue := extractCalibrationValue(line)
		totalSum += calibrationValue
	}
	return totalSum
}
func part2(document []string) int {
	totalSum := 0
	for _, line := range document {
		calibrationValue := extractCalibrationValueWithWords(line)
		totalSum += calibrationValue
	}
	return totalSum
}
