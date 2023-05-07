package main

import (
	"bufio"
	"bytes"

	"github.com/mdwhatcott/advent-of-code/go/2016/util/assembunny"
	"github.com/mdwhatcott/advent-of-code/go/2016/util/lcd"
	"github.com/mdwhatcott/advent-of-code/go/lib/util"
)

func main() {
	interpreter := assembunny.NewInterpreter(util.InputLines())
	interpreter.ExecuteProgram()
	lcdInstructions := interpreter.Out()
	scanner := bufio.NewScanner(bytes.NewReader(lcdInstructions))
	lcd.Display(scanner, true)
}
