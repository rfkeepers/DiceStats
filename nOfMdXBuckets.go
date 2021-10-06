package main

import (
	"fmt"
)

// N of MdX pass given success range
var dice = []int{12, 12, 12}
var successes = []int{9, 10, 11}

type stat struct {
	bucket int
	count  int
	chance float64
}

func main() {
	for _, s := range successes {
		dieProbs(s)
	}
}

func dieProbs(success int) {
	fmt.Printf("Matching\t---\td%d min %d\n", len(dice), success)
	results, total := getProbs(len(dice), success)
	for _, r := range results {
		fmt.Printf("count\t%d\n", r.bucket)
		fmt.Printf("\t%d/%d\n", r.count, total)
		fmt.Printf("\t%.2f\n", r.chance)
	}
}

func getProbs(nDice, success int) ([]stat, int) {
	bucketCount := map[int]int{-1: 0}
	initCounts(bucketCount, nDice, success)
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

func initCounts(counts map[int]int, nDice, success int) {
	for i := 1; i <= dice[0]; i++ {
		addCount(counts, []int{i}, 1, nDice, success)
	}
}

func addCount(counts map[int]int, set []int, die, nDice, success int) {
	if die == len(dice) {
		counts[-1] = counts[-1] + 1
		nSuccess := numberOfSuccesses(set, success)
		counts[nSuccess] = counts[nSuccess] + 1
	} else {
		for i := 1; i <= dice[die]; i++ {
			addCount(counts, append(set, i), die+1, nDice, success)
		}
	}
}

func numberOfSuccesses(rolls []int, success int) int {
	passes := 0
	for _, r := range rolls {
		if r >= success {
			passes = passes + 1
		}
	}
	return passes
}
