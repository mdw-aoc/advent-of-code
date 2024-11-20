package starter

import (
	"strings"
	"testing"

	"github.com/mdw-aoc/inputs"
	_ "github.com/mdw-go/funcy"
	. "github.com/mdw-go/funcy/ranger"
	_ "github.com/mdw-go/must"
	_ "github.com/mdw-go/set"
	"github.com/mdw-go/testing/should"
)

const TODO = -1

var (
	inputElements  = strings.Split(inputs.Read(2023, 15).String(), ",")
	sampleElements = strings.Split("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ",")
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Suite struct {
	*should.T
}

func (this *Suite) Setup() {
}
func (this *Suite) TestPart1Samples() {
	this.So(this.Part1(sampleElements), should.Equal, 1320)
}
func (this *Suite) TestPart1Full() {
	this.So(this.Part1(inputElements), should.Equal, 510273)
}
func (this *Suite) TestPart2Samples() {
	this.So(this.Part2(sampleElements), should.Equal, TODO)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputElements), should.Equal, TODO)
}
func (this *Suite) Part1(elements []string) any {
	return Sum(Map(this.HASH, Iterator(elements)))
}
func (this *Suite) Part2(elements []string) any {
	return TODO
}

func (this *Suite) HASH(s string) (result int) {
	for _, c := range s {
		result += int(c)
		result *= 17
		result %= 256
	}
	return result
}
