package starter

import (
	"strings"
	"testing"
	"unicode"

	"github.com/mdw-aoc/inputs"
	. "github.com/mdw-go/funcy/ranger"
	"github.com/mdw-go/must/strconvmust"
	"github.com/mdw-go/testing/should"
)

var (
	inputElements  = strings.Split(inputs.Read(2023, 15).String(), ",")
	sampleElements = strings.Split("rn=1,cm-,qp=3,cm=2,qp-,pc=4,ot=9,ab=5,pc-,pc=6,ot=7", ",")
)

func TestSuite(t *testing.T) {
	should.Run(&Suite{T: should.New(t)}, should.Options.UnitTests())
}

type Lens struct {
	BoxID       int
	LensID      int
	FocalLength int
	Label       string
}

type Box struct {
	ID     int
	Lenses []Lens
}

type Suite struct {
	*should.T
	boxes [256]Box
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
	this.So(this.Part2(sampleElements), should.Equal, 145)
}
func (this *Suite) TestPart2Full() {
	this.So(this.Part2(inputElements), should.Equal, 212449)
}
func (this *Suite) Part1(elements []string) any {
	return Sum(Map(this.Hash, Iterator(elements)))
}
func (this *Suite) Part2(elements []string) any {
	DoAll(this.PlaceLens, Iterator(elements))
	return Sum(Map(this.BoxFocusingPower, Iterator(this.boxes[:])))
}
func (this *Suite) Hash(s string) (result int) {
	for _, c := range s {
		result += int(c)
		result *= 17
		result %= 256
	}
	return result
}
func (this *Suite) PlaceLens(rawLens string) {
	chars := Iterator([]rune(rawLens))
	label := string(Slice(TakeWhile(unicode.IsLetter, chars)))
	boxID := this.Hash(label)
	action := First(DropWhile(unicode.IsLetter, chars))
	if action == '-' {
		matchingLabel := func(lens Lens) bool { return lens.Label == label }
		this.boxes[boxID].Lenses = Slice(Remove(matchingLabel, Iterator(this.boxes[boxID].Lenses)))
		for b := boxID; b < len(this.boxes); b++ {
			for l := range this.boxes[b].Lenses {
				this.boxes[b].Lenses[l].LensID = l
			}
		}
	} else if action == '=' {
		focalLength := strconvmust.Atoi(string(Slice(Filter(unicode.IsDigit, chars))))
		found := false
		for l, lens := range this.boxes[boxID].Lenses {
			if lens.Label == label {
				found = true
				this.boxes[boxID].Lenses[l].FocalLength = focalLength
			}
		}
		if !found {
			this.boxes[boxID].ID = boxID
			this.boxes[boxID].Lenses = append(this.boxes[boxID].Lenses, Lens{
				BoxID:       boxID,
				LensID:      len(this.boxes[boxID].Lenses),
				Label:       label,
				FocalLength: focalLength,
			})
		}
	} else {
		this.Fatal("expected '-' or '=', got:", action)
	}
}
func (this Lens) MatchesName(label string) bool {
	return this.Label == label
}
func (this *Suite) BoxFocusingPower(box Box) (result int) {
	return Sum(Map(this.LensFocusingPower, Iterator(box.Lenses)))
}
func (this *Suite) LensFocusingPower(lens Lens) (result int) {
	result = lens.BoxID + 1
	result *= lens.LensID + 1
	result *= lens.FocalLength
	return result
}
