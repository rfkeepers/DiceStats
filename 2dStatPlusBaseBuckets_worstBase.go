package main

import (
	"fmt"
)

// bucketed results for rolling 2 dice, one static, one acting as a graded stat
// this version rolls the stat dice twice and takes the lower result

const die1 = 6

var die2 = []int{4, 6, 8, 10, 12}
var buckets = []int{10, 7, 1}

type stat struct {
	bucket int
	count  int
	chance float64
}

func main() {
	fmt.Printf("Base die:\t---\td%d\n\n", die1)
	for _, die := range die2 {
		dieProbs(die1, die)
	}
}

func dieProbs(d1sides, d2sides int) {
	fmt.Printf("Upgrade die\t---\td%d\n", d2sides)
	results, total := getProbs(d1sides, d2sides)
	for _, r := range results {
		fmt.Printf("bucket\t%d\n", r.bucket)
		fmt.Printf("\t%d/%d\n", r.count, total)
		fmt.Printf("\t%.2f\n", r.chance)
	}
	fmt.Println("")
}

func getProbs(d1sides, d2sides int) ([]stat, int) {
	bucketCount := map[int]int{-1: 0}
	for _, b := range buckets {
		bucketCount[b] = 0
	}
	initCounts(bucketCount, d1sides, d2sides)
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

func initCounts(counts map[int]int, d1sides, d2sides int) {
	for i := 1; i <= d2sides; i++ {
		addCount(counts, i, -1, -1, d1sides)
	}
}

func addCount(counts map[int]int, d1, d2a, d2b, d1sides int) {
	if d2a > 0 {
		counts[-1] = counts[-1] + 1
		nsum := sum(d1, d2a, d2b)
		for _, b := range buckets {
			if nsum >= b {
				counts[b] = counts[b] + 1
				break
			}
		}
		return
	}
	for i := 1; i <= d1sides; i++ {
		for j := 1; j <= d1sides; j++ {
			addCount(counts, d1, i, j, d1sides)
		}
	}
}

func sum(i, j, k int) int {
	if j <= k {
		return i + j
	}
	return i + k
}
