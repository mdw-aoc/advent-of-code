package advent

import (
	"strings"
	"testing"

	"github.com/mdwhatcott/testing/should"
)

func TestParseFieldDefinition(t *testing.T) {
	parsed := ParseFieldDefinition("departure location: 25-568 or 594-957")

	should.So(t, parsed.Name, should.Equal, "departure location")
	should.So(t, parsed.AllWithinRange(25, 100, 568, 594, 799, 957), should.BeTrue)
	should.So(t, parsed.WithinRange(24), should.BeFalse)
}

func TestParseAllFieldDefinitions(t *testing.T) {
	all := ParseAllFieldDefinitions(exampleInput1)
	should.So(t, len(all), should.Equal, 3)
	should.So(t, all[0].Name, should.Equal, "class")
}

func TestParseAllTickets(t *testing.T) {
	all := ParseAllTickets(exampleInput1)
	should.So(t, len(all), should.Equal, 5)
	should.So(t, all[0], should.Equal, []int{7, 1, 14})
	should.So(t, all[1], should.Equal, []int{7, 3, 47})
}

func TestCalculateErrorRate(t *testing.T) {
	rate := CalculateErrorRate(ParseAllFieldDefinitions(exampleInput1), ParseAllTickets(exampleInput1))
	should.So(t, rate, should.Equal, 71)
}

func TestFilterValid(t *testing.T) {
	valid := FilterValidTickets(ParseAllFieldDefinitions(exampleInput1), ParseAllTickets(exampleInput1))
	should.So(t, len(valid), should.Equal, 2) // your ticket and one of the nearby tickets

	valid2 := FilterValidTickets(ParseAllFieldDefinitions(exampleInput2), ParseAllTickets(exampleInput2))
	should.So(t, len(valid2), should.Equal, 4) // your ticket and one of the nearby tickets
}

var exampleInput1 = strings.TrimSpace(`
class: 1-3 or 5-7
row: 6-11 or 33-44
seat: 13-40 or 45-50

your ticket:
7,1,14

nearby tickets:
7,3,47
40,4,50
55,2,20
38,6,12
`)

var exampleInput2 = strings.TrimSpace(`
class: 0-1 or 4-19
row: 0-5 or 8-19
seat: 0-13 or 16-19

your ticket:
11,12,13

nearby tickets:
3,9,18
15,1,5
5,14,9
`)

func TestCandidateFields(t *testing.T) {
	definitions := ParseAllFieldDefinitions(exampleInput1)
	tickets := FilterValidTickets(definitions, ParseAllTickets(exampleInput1))
	placements := IdentifyFieldPlacementCandidates(definitions, tickets)
	should.So(t, placements, should.Equal, map[string][]int{
		"class": {0, 1},
		"row":   {0},
		"seat":  {2},
	})
}

func TestFinalizeFieldPlacements(t *testing.T) {
	candidates := map[string][]int{
		"class": {0, 1},
		"row":   {0},
		"seat":  {2},
	}
	finalized := FinalizeFieldPlacements(candidates)
	should.So(t, finalized, should.Equal, map[string]int{
		"class": 1,
		"row":   0,
		"seat":  2,
	})
}

func TestCalculateDepartureProduct(t *testing.T) {
	definitions := ParseAllFieldDefinitions(exampleInput2)
	tickets := FilterValidTickets(definitions, ParseAllTickets(exampleInput2))
	candidates := IdentifyFieldPlacementCandidates(definitions, tickets)
	finalized := FinalizeFieldPlacements(candidates)
	finalized["departure row"] = finalized["row"]
	finalized["departure seat"] = finalized["seat"]
	product := CalculateDepartureProduct(finalized, tickets[0])
	should.So(t, product, should.Equal, 11*13)
}
