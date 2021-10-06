package main

import (
	"fmt"
)

const (
	dice   = 5
	sides  = 6
	offset = 0
)

var buckets = []int{6, 4, 1}

type stat struct {
	bucket int
	count  int
	crits  float64
	chance float64
}

func main() {
	for i := 1; i <= dice; i++ {
		dieProbs(i)
	}
}

func dieProbs(nDice int) {
	fmt.Printf("Dice count\t---\t%d\n", nDice)
	results, total, crits := getProbs(nDice)
	fmt.Printf("Crit chance\t%.2f\n", crits)
	for _, r := range results {
		fmt.Printf("bucket\t%d\n", r.bucket)
		fmt.Printf("\t%d/%d\n", r.count, total)
		fmt.Printf("\t%.2f\n", r.chance)
	}
}

func getProbs(nDice int) ([]stat, int, float64) {
	bucketCount := map[int]int{-1: 0}
	for _, b := range buckets {
		bucketCount[b] = 0
	}
	initCounts(bucketCount, nDice)
	stats := []stat{}
	total := bucketCount[-1]
	crits := bucketCount[-2]
	for _, b := range buckets {
		s := stat{
			bucket: b,
			count:  bucketCount[b],
			chance: float64(bucketCount[b]) / float64(total),
		}
		stats = append(stats, s)
	}
	return stats, total, float64(crits) / float64(total)
}

func initCounts(counts map[int]int, nDice int) {
	for i := 1; i <= sides; i++ {
		addCount(counts, []int{i}, i, 1, nDice)
	}
}

func addCount(counts map[int]int, set []int, setMax, die, nDice int) {
	if die == nDice {
		counts[-1] = counts[-1] + 1
		for _, b := range buckets {
			if (setMax + offset) >= b {
				counts[b] = counts[b] + 1
				break
			}
		}
		sixes := 0
		for _, d := range set {
			if d == 6 {
				sixes = sixes + 1
			}
		}
		if sixes > 1 {
			counts[-2] = counts[-2] + 1
		}
		return
	}
	for i := 1; i <= sides; i++ {
		max := setMax
		if i > max {
			max = i
		}
		addCount(counts, append(set, i), max, die+1, nDice)
	}
}
