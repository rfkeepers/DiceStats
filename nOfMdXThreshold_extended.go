package main

import (
	"fmt"
)

// N of MdX pass a success threshold

// count of dice rolled (len), and sides of each dice (val)
var dice = []int{6, 6, 6, 6, 6, 6}

// min face value needed for success
// array allowed for quality of life if you're checking
// multiple possible values
var successes = []int{5}

// number of successful rolls required for each bucket
var buckets = []int{0, 1, 2, 3}
var andUp = 3

// modifiers to roll results
var plusMod = 0

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
	results, total := getProbs(success)
	for _, r := range results {
		fmt.Printf("successes:\t%d\n", r.bucket)
		fmt.Printf("\t\t%d/%d\n", r.count, total)
		fmt.Printf("\t\t%.2f\n", r.chance)
	}
}

func getProbs(success int) ([]stat, int) {
	bucketCount := map[int]int{-1: 0}
	stats := []stat{}

	runCounts(bucketCount, success)
	total := bucketCount[-1]

	// for i := 0; i <= len(dice); i++ {

	for _, bkt := range buckets {
		s := stat{
			bucket: bkt,
			count:  bucketCount[bkt],
			chance: float64(bucketCount[bkt]) / float64(total),
		}
		stats = append(stats, s)
	}

	return stats, total
}

func runCounts(counts map[int]int, success int) {
	for i := 1; i <= dice[0]; i++ {
		addCount(counts, []int{i}, 1, success)
	}
}

func addCount(counts map[int]int, set []int, rolled, success int) {
	if rolled == len(dice) {
		counts[-1] = counts[-1] + 1
		nSuccess := numberOfSuccesses(set, success)
		for _, bkt := range buckets {
			if nSuccess == bkt {
				counts[bkt] = counts[bkt] + 1
			}
			if bkt == andUp && nSuccess > bkt {
				counts[bkt] = counts[bkt] + 1
			}
		}
	} else {
		for i := 1; i <= dice[rolled]; i++ {
			addCount(counts, append(set, i), rolled+1, success)
		}
	}
}

func numberOfSuccesses(rolls []int, success int) int {
	passes := 0
	for _, r := range rolls {
		if r+plusMod >= success {
			passes = passes + 1
		}
	}
	return passes
}
