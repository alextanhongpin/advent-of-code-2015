// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("part 1:", maxHappiness(example, false))
	fmt.Println("part 1:", maxHappiness(input, false))

	fmt.Println("part 2:", maxHappiness(input, true))
}

func maxHappiness(input string, includeYourself bool) int {
	input = strings.ReplaceAll(input, ".", "")
	lines := strings.Split(input, "\n")
	happiness := make(map[string]int)
	individuals := make(map[string]bool)
	if includeYourself {
		individuals["You"] = true
	}
	for _, line := range lines {
		parts := strings.Fields(line)
		a, b, delta, unit := parts[0], parts[len(parts)-1], parts[2], toInt(parts[3])
		if delta != "gain" {
			unit *= -1
		}
		happiness[a+":"+b] = unit
		individuals[a] = true
		individuals[b] = true
	}
	dp := make(map[string]int)
	var maxHappiness func(arrangements []string) int
	maxHappiness = func(arrangements []string) int {
		key := strings.Join(arrangements, ",")
		if v, ok := dp[key]; ok {
			return v
		}
		n := len(arrangements)
		unit := 0
		for i, ind := range arrangements {
			li, ri := (i-1)%n, (i+1)%n
			if li < 0 {
				li += n
			}
			left := ind + ":" + arrangements[li]
			right := ind + ":" + arrangements[ri]
			unit += happiness[left] + happiness[right]
		}
		dp[key] = unit
		var units []int
		for i := 0; i < len(arrangements); i++ {
			units = append(units, maxHappiness(swap(arrangements, 0, i)))
		}
		dp[key] = max(unit, units...)
		return max(unit, units...)
	}

	var arrangements []string
	for ind := range individuals {
		arrangements = append(arrangements, ind)
	}
	return maxHappiness(arrangements)
}

func swap(positions []string, i, j int) []string {
	res := make([]string, len(positions))
	copy(res, positions)
	res[i], res[j] = res[j], res[i]
	return res
}

func max(head int, rest ...int) int {
	for _, r := range rest {
		if r > head {
			head = r
		}
	}
	return head
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `Alice would gain 54 happiness units by sitting next to Bob.
Alice would lose 79 happiness units by sitting next to Carol.
Alice would lose 2 happiness units by sitting next to David.
Bob would gain 83 happiness units by sitting next to Alice.
Bob would lose 7 happiness units by sitting next to Carol.
Bob would lose 63 happiness units by sitting next to David.
Carol would lose 62 happiness units by sitting next to Alice.
Carol would gain 60 happiness units by sitting next to Bob.
Carol would gain 55 happiness units by sitting next to David.
David would gain 46 happiness units by sitting next to Alice.
David would lose 7 happiness units by sitting next to Bob.
David would gain 41 happiness units by sitting next to Carol.`

var input = `Alice would lose 2 happiness units by sitting next to Bob.
Alice would lose 62 happiness units by sitting next to Carol.
Alice would gain 65 happiness units by sitting next to David.
Alice would gain 21 happiness units by sitting next to Eric.
Alice would lose 81 happiness units by sitting next to Frank.
Alice would lose 4 happiness units by sitting next to George.
Alice would lose 80 happiness units by sitting next to Mallory.
Bob would gain 93 happiness units by sitting next to Alice.
Bob would gain 19 happiness units by sitting next to Carol.
Bob would gain 5 happiness units by sitting next to David.
Bob would gain 49 happiness units by sitting next to Eric.
Bob would gain 68 happiness units by sitting next to Frank.
Bob would gain 23 happiness units by sitting next to George.
Bob would gain 29 happiness units by sitting next to Mallory.
Carol would lose 54 happiness units by sitting next to Alice.
Carol would lose 70 happiness units by sitting next to Bob.
Carol would lose 37 happiness units by sitting next to David.
Carol would lose 46 happiness units by sitting next to Eric.
Carol would gain 33 happiness units by sitting next to Frank.
Carol would lose 35 happiness units by sitting next to George.
Carol would gain 10 happiness units by sitting next to Mallory.
David would gain 43 happiness units by sitting next to Alice.
David would lose 96 happiness units by sitting next to Bob.
David would lose 53 happiness units by sitting next to Carol.
David would lose 30 happiness units by sitting next to Eric.
David would lose 12 happiness units by sitting next to Frank.
David would gain 75 happiness units by sitting next to George.
David would lose 20 happiness units by sitting next to Mallory.
Eric would gain 8 happiness units by sitting next to Alice.
Eric would lose 89 happiness units by sitting next to Bob.
Eric would lose 69 happiness units by sitting next to Carol.
Eric would lose 34 happiness units by sitting next to David.
Eric would gain 95 happiness units by sitting next to Frank.
Eric would gain 34 happiness units by sitting next to George.
Eric would lose 99 happiness units by sitting next to Mallory.
Frank would lose 97 happiness units by sitting next to Alice.
Frank would gain 6 happiness units by sitting next to Bob.
Frank would lose 9 happiness units by sitting next to Carol.
Frank would gain 56 happiness units by sitting next to David.
Frank would lose 17 happiness units by sitting next to Eric.
Frank would gain 18 happiness units by sitting next to George.
Frank would lose 56 happiness units by sitting next to Mallory.
George would gain 45 happiness units by sitting next to Alice.
George would gain 76 happiness units by sitting next to Bob.
George would gain 63 happiness units by sitting next to Carol.
George would gain 54 happiness units by sitting next to David.
George would gain 54 happiness units by sitting next to Eric.
George would gain 30 happiness units by sitting next to Frank.
George would gain 7 happiness units by sitting next to Mallory.
Mallory would gain 31 happiness units by sitting next to Alice.
Mallory would lose 32 happiness units by sitting next to Bob.
Mallory would gain 95 happiness units by sitting next to Carol.
Mallory would gain 91 happiness units by sitting next to David.
Mallory would lose 66 happiness units by sitting next to Eric.
Mallory would lose 75 happiness units by sitting next to Frank.
Mallory would lose 99 happiness units by sitting next to George.`
