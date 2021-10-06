package main

import (
	"fmt"
	"sort"
)

// roll N dice, bucket the results, take only the highest, outermost, and lowest results.

const (
	dice   = 3
	offset = 0
)

var sides = []int{6, 6, 6}
var buckets = []int{9, 6, 1}
var take = []string{"h", "o", "l"}

type stat struct {
	bucket int
	count  int
	chance float64
}

func main() {
	for _, v := range take {
		dieProbs(v)
	}
}

func dieProbs(take string) {
	fmt.Printf("take\t\t---\t%s\n", take)
	results, total := getProbs(take)
	for _, r := range results {
		fmt.Printf("bucket\t%d\n", r.bucket)
		fmt.Printf("\t%d/%d\n", r.count, total)
		fmt.Printf("\t%.2f\n", r.chance)
	}
}

func getProbs(take string) ([]stat, int) {
	bucketCount := map[int]int{-1: 0}
	for _, b := range buckets {
		bucketCount[b] = 0
	}
	initCounts(bucketCount, take)
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

func initCounts(counts map[int]int, take string) {
	for i := 1; i <= sides[0]; i++ {
		addCount(counts, []int{i}, 1, take)
	}
}

func addCount(counts map[int]int, set []int, die int, take string) {
	if die == dice {
		counts[-1] = counts[-1] + 1
		sort.Ints(set)
		sum := offset
		l := len(set)
		switch take {
		case "h":
			sum = sum + set[l-1] + set[l-2]
		case "o":
			sum = sum + set[0] + set[l-1]
		case "l":
			sum = sum + set[0] + set[1]
		}
		for _, b := range buckets {
			if sum >= b {
				counts[b] = counts[b] + 1
				break
			}
		}
		return
	}
	for i := 1; i <= sides[die]; i++ {
		addCount(counts, append(set, i), die+1, take)
	}
}

func highest(rolls []int, without int) (int, int) {
	var high, idx int
	for i, v := range rolls {
		if i == without {
			continue
		}
		if v > high {
			high = v
			idx = i
		}
	}
	return high, idx
}

func sumSlice(sl []int) int {
	sum := 0
	for _, v := range sl {
		sum = sum + v
	}
	return sum
}
