package main

import (
	"fmt"
)

// N of MdX pass a success threshold
var dice = []int{6, 6, 6}
var success = 5

type stat struct {
	bucket int
	count  int
	chance float64
}

func main() {
	dieProbs()
}

func dieProbs() {
	fmt.Printf("Matching\t---\t%d\n", len(dice))
	results, total := getProbs(len(dice))
	for _, r := range results {
		fmt.Printf("count\t%d\n", r.bucket)
		fmt.Printf("\t%d/%d\n", r.count, total)
		fmt.Printf("\t%.2f\n", r.chance)
	}
}

func getProbs(nDice int) ([]stat, int) {
	bucketCount := map[int]int{-1: 0}
	initCounts(bucketCount, nDice)
	stats := []stat{}
	total := bucketCount[-1]
	for i := 0; i <= len(dice); i++ {
		s := stat{
			bucket: i,
			count:  bucketCount[i],
			chance: float64(bucketCount[i]) / float64(total),
		}
		stats = append(stats, s)
	}
	return stats, total
}

func initCounts(counts map[int]int, nDice int) {
	for i := 1; i <= dice[0]; i++ {
		addCount(counts, []int{i}, 1, nDice)
	}
}

func addCount(counts map[int]int, set []int, die, nDice int) {
	if die == len(dice) {
		counts[-1] = counts[-1] + 1
		nSuccess := numberOfSuccesses(set)
		counts[nSuccess] = counts[nSuccess] + 1
	} else {
		for i := 1; i <= dice[die]; i++ {
			addCount(counts, append(set, i), die+1, nDice)
		}
	}
}

func numberOfSuccesses(rolls []int) int {
	passes := 0
	for _, r := range rolls {
		if r >= success {
			passes = passes + 1
		}
	}
	return passes
}
