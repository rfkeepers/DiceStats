package main

import (
	"fmt"
	"sort"
)

// how many of M dice succeed
var sides = 12
var instabilityDice = []int{6, 8, 10}

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
	
	for _, order := range []string{"advantage", "disadvantage"} {
		r := getStats(instabilityDie, order)
		fmt.Printf(
`
|%v|%v%%|%v%%|%v%%|`,
			order,
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

func getStats(instabilityDie int, order string) stat {
	counts := map[string]int{"total": 0}
	initCounts(counts, instabilityDie, order)
	return stat{
		strong: counts["strong"],
		weak: counts["weak"],
		miss: counts["miss"],
		total: counts["total"],
	}
}

func initCounts(counts map[string]int, instabilityDie int, order string) {
	for i := 1; i <= instabilityDie; i++ {
		for d1 := 1; d1 <= sides; d1++ {
			for d2 := 1; d2 <= sides; d2 ++ {
				for d3 := 1; d3 <= sides; d3++ {
					addCount(counts, i, d1, d2, d3, order)
				}	
			}
		}
	}
}

func addCount(counts map[string]int, dInst, d1, d2, d3 int, order string) {
	counts["total"] = counts["total"] + 1
	
	dice := []int{d1, d2, d3}
	sort.Ints(dice)
	if order == "disadvantage" {
		dice[0], dice[2] = dice[2], dice[0]
	}
	
	result := "weak"
	if dice[0] <= dInst && dice[1] <= dInst {
		result = "strong"
	} else if dice[0] > dInst && dice[1] > dInst {
		result = "miss"
	}

	counts[result] = counts[result]+1
}
