package day19

import (
	"bufio"
	"strings"

	"advent/lib/util"
)

func ParseScannerReports(reports string) (results [][]Point) {
	scanner := bufio.NewScanner(strings.NewReader(strings.TrimSpace(reports) + "\n"))
	var beacon []Point
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			results = append(results, beacon)
			beacon = nil
			continue
		}
		if strings.HasPrefix(line, "---") {
			continue
		}
		fields := strings.Split(line, ",")
		x := util.ParseInt(fields[0])
		y := util.ParseInt(fields[1])
		z := util.ParseInt(fields[2])
		beacon = append(beacon, NewPoint(x, y, z))
	}
	return append(results, beacon)
}

func AreOverlapping(group1 []Point, group2 []Point) bool {
	return false // TODO
	//g2 := set.From[Point](group2...)
	//for x := 0; x < 24; x++ {
	//	g1 := set.From[Point](RotateAll(x, group1...)...)
	//	for p1 := range g1 {
	//		for _, p2 := range group2 {
	//			diff := Diff(p1, p2)
	//
	//		}
	//	}
	//}
}
