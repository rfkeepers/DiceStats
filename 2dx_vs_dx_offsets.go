package main

import (
	"fmt"
)

// how many of M dice succeed
var sides = 12
var instabilityDice = []int{6, 8, 10}
var minOffset = -3
var maxOffset = 3

type stat struct {
	strong int
	weak int
	miss int
	total  int
}

func main() {
	for _, iDie := range instabilityDice {
		dieProbs(iDie)
	}
}

func dieProbs(instabilityDie int) {		
	fmt.Printf(
`
| d%d | strong | weak | miss |
|:----|:---:|:---:|:---:|`,
instabilityDie)
	
	for i := minOffset; i <= maxOffset; i++ {
	r := getStats(instabilityDie, i)
	fmt.Printf(
`
|+%d|%v%%|%v%%|%v%%|`,
i,
getPercent(r.strong, r.total),
getPercent(r.weak, r.total),
getPercent(r.miss, r.total))
	}
}

func getPercent(n, d int) float64 {
	p := float64(n)/float64(d)
	i := int(p * 10000)
	return float64(i)/100.0
}

func getStats(instabilityDie, offset int) stat {
	counts := map[string]int{"total": 0}
	initCounts(counts, instabilityDie, offset)
	return stat{
		strong: counts["strong"],
		weak: counts["weak"],
		miss: counts["miss"],
		total: counts["total"],
	}
}

func initCounts(counts map[string]int, instabilityDie, offset int) {
	for i := 1; i <= instabilityDie; i++ {
		for d1 := 1; d1 <= sides; d1++ {
			for d2 := 1; d2 <= sides; d2 ++ {
				addCount(counts, i+offset, d1, d2)	
			}
		}
	}
}

func addCount(counts map[string]int, dInst, d1, d2 int) {
	counts["total"] = counts["total"] + 1
	
	result := "weak"
	if d1 <= dInst && d2 <= dInst {
		result = "strong"
	} else if d1 > dInst && d2 > dInst {
		result = "miss"
	}

	counts[result] = counts[result]+1
}
