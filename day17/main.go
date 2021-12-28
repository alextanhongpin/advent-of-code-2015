// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

func main() {
	buckets := []int{20, 15, 10, 5, 5}
	fmt.Println(countCombinations(buckets, len(buckets)-1, 25, ""))

	lines := strings.Split(input, "\n")
	buckets = make([]int, len(lines))
	for i, line := range lines {
		n, _ := strconv.Atoi(line)
		buckets[i] = n
	}
	fmt.Println("part 1:", countCombinations(buckets, len(buckets)-1, 150, ""))
	fmt.Println(countMinCombinations(buckets, len(buckets)-1, 150, 0, ""))
}

func countCombinations(buckets []int, bucket int, liter int, display string) int {
	// One successful step.
	if liter == 0 {
		// fmt.Println(display)
		return 1
	}

	// No successful combination.
	if liter < 0 || bucket < 0 {
		return 0
	}
	nextDisplay := display + "," + fmt.Sprint(buckets[bucket])
	return countCombinations(buckets, bucket-1, liter-buckets[bucket], strings.Trim(nextDisplay, ",")) +
		countCombinations(buckets, bucket-1, liter, display)
}

func countMinCombinations(buckets []int, bucket int, liter, n int, display string) (count, min int) {
	// One successful step.
	if liter == 0 {
		count = 1
		min = n
		// fmt.Println(display)
		return
	}

	// No successful combination.
	if liter < 0 || bucket < 0 {
		min = math.MaxInt
		return
	}
	nextDisplay := display + "," + fmt.Sprint(buckets[bucket])
	l, lm := countMinCombinations(buckets, bucket-1, liter-buckets[bucket], n+1, strings.Trim(nextDisplay, ","))
	r, rm := countMinCombinations(buckets, bucket-1, liter, n, display)
	if lm < rm {
		count = l
		min = lm
	} else if lm > rm {
		count = r
		min = rm
	} else {
		count = l + r
		min = rm
	}
	return
}

var input = `50
44
11
49
42
46
18
32
26
40
21
7
18
43
10
47
36
24
22
40`
