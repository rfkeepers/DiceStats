package main

import (
	"fmt"
)

// how many of M dice succeed
var dice = []int{6, 6, 6, 6}
var buckets = []int{4, 1}

type stat struct {
	bucket int
	count  int
	chance float64
}

func main() {
	for i := 2; i <= len(dice); i++ {
		dieProbs(i)
	}
}

func dieProbs(nDice int) {
	fmt.Printf("Matching\t---\t%d of %d\n", nDice, len(dice))
	results, total := getProbs(nDice)
	for _, r := range results {
		fmt.Printf("bucket\t%d\n", r.bucket)
		fmt.Printf("\t%d/%d\n", r.count, total)
		fmt.Printf("\t%.2f\n", r.chance)
	}
}

func getProbs(nDice int) ([]stat, int) {
	bucketCount := map[int]int{-1: 0}
	for _, b := range buckets {
		bucketCount[b] = 0
	}
	initCounts(bucketCount, nDice)
	stats := []stat{}
	total := bucketCount[-1]
	for _, b := range buckets {
		s := stat{
			bucket: b,
			count:  bucketCount[b],
			chance: float64(bucketCount[b]) / float64(total),
		}
		stats = append(stats, s)
	}
	return stats, total
}

func initCounts(counts map[int]int, nDice int) {
	for i := 1; i <= dice[0]; i++ {
		addCount(counts, []int{i}, i, 1, nDice)
	}
}

func addCount(counts map[int]int, set []int, setMax, die, nDice int) {
	if die == len(dice) {
		counts[-1] = counts[-1] + 1
		setCounts := map[int]int{}
		for _, die := range set {
			for _, bkt := range buckets {
				if die >= bkt {
					setCounts[bkt] = setCounts[bkt] + 1
					break
				}
			}
		}
		for _, b := range buckets {
			if setCounts[b] >= nDice {
				counts[b] = counts[b] + 1
				break
			}
			if b == 1 {
				counts[b] = counts[b] + 1
			}
		}
	} else {
		for i := 1; i <= dice[die]; i++ {
			max := setMax
			if i > max {
				max = i
			}
			addCount(counts, append(set, i), max, die+1, nDice)
		}
	}
}
