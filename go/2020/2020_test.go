package advent

import (
	"testing"

	day01 "github.com/mdwhatcott/advent-of-code/go/2020/day01"
	day02 "github.com/mdwhatcott/advent-of-code/go/2020/day02"
	day03 "github.com/mdwhatcott/advent-of-code/go/2020/day03"
	day04 "github.com/mdwhatcott/advent-of-code/go/2020/day04"
	day05 "github.com/mdwhatcott/advent-of-code/go/2020/day05"
	day06 "github.com/mdwhatcott/advent-of-code/go/2020/day06"
	day07 "github.com/mdwhatcott/advent-of-code/go/2020/day07"
	day08 "github.com/mdwhatcott/advent-of-code/go/2020/day08"
	day09 "github.com/mdwhatcott/advent-of-code/go/2020/day09"
	day10 "github.com/mdwhatcott/advent-of-code/go/2020/day10"
	day11 "github.com/mdwhatcott/advent-of-code/go/2020/day11"
	day12 "github.com/mdwhatcott/advent-of-code/go/2020/day12"
	day13 "github.com/mdwhatcott/advent-of-code/go/2020/day13"
	day14 "github.com/mdwhatcott/advent-of-code/go/2020/day14"
	day15 "github.com/mdwhatcott/advent-of-code/go/2020/day15"
	day16 "github.com/mdwhatcott/advent-of-code/go/2020/day16"
	day17 "github.com/mdwhatcott/advent-of-code/go/2020/day17"
	day18 "github.com/mdwhatcott/advent-of-code/go/2020/day18"
	day19 "github.com/mdwhatcott/advent-of-code/go/2020/day19"
	day20 "github.com/mdwhatcott/advent-of-code/go/2020/day20"
	day21 "github.com/mdwhatcott/advent-of-code/go/2020/day21"
	day22 "github.com/mdwhatcott/advent-of-code/go/2020/day22"
	day23 "github.com/mdwhatcott/advent-of-code/go/2020/day23"
	day24 "github.com/mdwhatcott/advent-of-code/go/2020/day24"
	day25 "github.com/mdwhatcott/advent-of-code/go/2020/day25"

	"github.com/mdwhatcott/testing/should"
)

func TestFixture2020(t *testing.T) {
	should.Run(&Fixture2020{T: should.New(t)}, should.Options.UnitTests())
}

type Fixture2020 struct {
	*should.T
}

func (this *Fixture2020) LongTestDay01() {
	this.So(day01.Part1(), should.Equal, 1015476)
	this.So(day01.Part2(), should.Equal, 200878544)
}
func (this *Fixture2020) LongTestDay02() {
	this.So(day02.Part1(), should.Equal, 643)
	this.So(day02.Part2(), should.Equal, 388)
}
func (this *Fixture2020) LongTestDay03() {
	this.So(day03.Part1(), should.Equal, 225)
	this.So(day03.Part2(), should.Equal, 1115775000)
}
func (this *Fixture2020) LongTestDay04() {
	this.So(day04.Part1(), should.Equal, 264)
	this.So(day04.Part2(), should.Equal, 224)
}
func (this *Fixture2020) LongTestDay05() {
	this.So(day05.Part1(), should.Equal, 959)
	this.So(day05.Part2(), should.Equal, 527)
}
func (this *Fixture2020) LongTestDay06() {
	this.So(day06.Part1(), should.Equal, 6775)
	this.So(day06.Part2(), should.Equal, 3356)
}
func (this *Fixture2020) LongTestDay07() {
	this.So(day07.Part1(), should.Equal, 119)
	this.So(day07.Part2(), should.Equal, 155802)
}
func (this *Fixture2020) LongTestDay08() {
	this.So(day08.Part1(), should.Equal, 1832)
	this.So(day08.Part2(), should.Equal, 662)
}
func (this *Fixture2020) LongTestDay09() {
	this.So(day09.Part1(), should.Equal, 23278925)
	this.So(day09.Part2(), should.Equal, 4011064)
}
func (this *Fixture2020) LongTestDay10() {
	this.So(day10.Part1(), should.Equal, 2592)
	this.So(day10.Part2(), should.Equal, 198428693313536)
}
func (this *Fixture2020) LongTestDay11() {
	this.So(day11.Part1(), should.Equal, 2438)
	this.So(day11.Part2(), should.Equal, 2174)
}
func (this *Fixture2020) LongTestDay12() {
	this.So(day12.Part1(), should.Equal, 2847)
	this.So(day12.Part2(), should.Equal, 29839)
}
func (this *Fixture2020) LongTestDay13() {
	this.So(day13.Part1(), should.Equal, 4938)
	this.So(day13.Part2(), should.Equal, 230903629977901)
}
func (this *Fixture2020) LongTestDay14() {
	this.So(day14.Part1(), should.Equal, 18630548206046)
	this.So(day14.Part2(), should.Equal, 4254673508445)
}
func (this *Fixture2020) LongTestDay15() {
	this.So(day15.Part1(), should.Equal, 1025)
	this.So(day15.Part2(), should.Equal, 129262)
}
func (this *Fixture2020) LongTestDay16() {
	this.So(day16.Part1(), should.Equal, 25984)
	this.So(day16.Part2(), should.Equal, 1265347500049)
}
func (this *Fixture2020) LongTestDay17() {
	this.So(day17.Part1(), should.Equal, 265)
	this.So(day17.Part2(), should.Equal, 1936)
}
func (this *Fixture2020) LongTestDay18() {
	this.So(day18.Part1(), should.Equal, 6923486965641)
	this.So(day18.Part2(), should.Equal, 70722650566361)
}
func (this *Fixture2020) LongTestDay19() {
	this.So(day19.Part1(), should.Equal, nil)
	this.So(day19.Part2(), should.Equal, nil)
}
func (this *Fixture2020) LongTestDay20() {
	this.So(day20.Part1(), should.Equal, nil)
	this.So(day20.Part2(), should.Equal, nil)
}
func (this *Fixture2020) LongTestDay21() {
	this.So(day21.Part1(), should.Equal, nil)
	this.So(day21.Part2(), should.Equal, nil)
}
func (this *Fixture2020) LongTestDay22() {
	this.So(day22.Part1(), should.Equal, nil)
	this.So(day22.Part2(), should.Equal, nil)
}
func (this *Fixture2020) LongTestDay23() {
	this.So(day23.Part1(), should.Equal, nil)
	this.So(day23.Part2(), should.Equal, nil)
}
func (this *Fixture2020) LongTestDay24() {
	this.So(day24.Part1(), should.Equal, nil)
	this.So(day24.Part2(), should.Equal, nil)
}
func (this *Fixture2020) LongTestDay25() {
	this.So(day25.Part1(), should.Equal, nil)
	this.So(day25.Part2(), should.Equal, nil)
}
