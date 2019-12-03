package advent

import (
	"advent/lib/util"
	_ "advent/lib/util"
)

func Part1() interface{} {
	program := util.InputInts(",")
	RunTweakedProgram(program, 12, 2)
	return program[0]
}

func Part2() interface{} {
	original := util.InputInts(",")
	program := make([]int, len(original))
	for x := 0; x < 100; x++ {
		for y := 0; y < 100; y++ {
			copy(program, original)
			RunTweakedProgram(program, x, y)
			if program[0] == 19690720 {
				return 100*x + y
			}
		}
	}
	panic("solution unknown")
}
