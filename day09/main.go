// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

var dp map[string]int
var distance map[string]int

func main() {
	dp = make(map[string]int)
	distance = make(map[string]int)
	lines := strings.Split(input, "\n")
	cities := make(map[string]bool)
	for _, line := range lines {
		parts := strings.Split(line, " = ")
		fromTo := strings.Split(parts[0], " to ")
		from, to := fromTo[0], fromTo[1]
		dist, err := strconv.Atoi(parts[1])
		if err != nil {
			panic(err)
		}
		distance[from+":"+to] = dist
		distance[to+":"+from] = dist
		cities[from] = true
		cities[to] = true
	}
	fmt.Println(minDistance(cities, "", []string{})) // 141
	fmt.Println(maxDistance(cities, "", []string{})) // 736
}

func maxDistance(cities map[string]bool, from string, visited []string) int {
	var toVisit []string
	for city := range cities {
		if !Contains(visited, city) {
			toVisit = append(toVisit, city)
		}
	}
	if len(toVisit) == 0 {
		key := strings.Join(visited, ",")
		return dp[key]
	}
	distances := make([]int, len(toVisit))
	for i, city := range toVisit {
		lst := ""
		if len(visited) > 0 {
			lst = visited[len(visited)-1]
		}
		distances[i] = maxDistance(cities, lst, append(visited, []string{city}...)) + distance[lst+":"+city]
	}
	visited = append(visited, from)
	key := strings.Join(visited, ",")
	dp[key] = max(distances...)
	return dp[key]
}

func minDistance(cities map[string]bool, from string, visited []string) int {
	var toVisit []string
	for city := range cities {
		if !Contains(visited, city) {
			toVisit = append(toVisit, city)
		}
	}
	if len(toVisit) == 0 {
		key := strings.Join(visited, ",")
		return dp[key]
	}
	distances := make([]int, len(toVisit))
	for i, city := range toVisit {
		lst := ""
		if len(visited) > 0 {
			lst = visited[len(visited)-1]
		}
		distances[i] = minDistance(cities, lst, append(visited, []string{city}...)) + distance[lst+":"+city]
	}
	visited = append(visited, from)
	key := strings.Join(visited, ",")
	dp[key] = min(distances...)
	return dp[key]
}

func min(nums ...int) int {
	switch len(nums) {
	case 0:
		return math.MinInt
	case 1:
		return nums[0]
	case 2:
		a, b := nums[0], nums[1]
		if a < b {
			return a
		}
		return b
	default:
		a, b, rest := nums[0], nums[1], nums[2:]
		return min(min(a, b), min(rest...))
	}
}
func max(nums ...int) int {
	switch len(nums) {
	case 0:
		return math.MaxInt
	case 1:
		return nums[0]
	case 2:
		a, b := nums[0], nums[1]
		if a > b {
			return a
		}
		return b
	default:
		a, b, rest := nums[0], nums[1], nums[2:]
		return max(max(a, b), max(rest...))
	}
}

func Contains(strs []string, v string) bool {
	for _, str := range strs {
		if str == v {
			return true
		}
	}
	return false
}

var example = `London to Dublin = 464
London to Belfast = 518
Dublin to Belfast = 141`

var input = `AlphaCentauri to Snowdin = 66
AlphaCentauri to Tambi = 28
AlphaCentauri to Faerun = 60
AlphaCentauri to Norrath = 34
AlphaCentauri to Straylight = 34
AlphaCentauri to Tristram = 3
AlphaCentauri to Arbre = 108
Snowdin to Tambi = 22
Snowdin to Faerun = 12
Snowdin to Norrath = 91
Snowdin to Straylight = 121
Snowdin to Tristram = 111
Snowdin to Arbre = 71
Tambi to Faerun = 39
Tambi to Norrath = 113
Tambi to Straylight = 130
Tambi to Tristram = 35
Tambi to Arbre = 40
Faerun to Norrath = 63
Faerun to Straylight = 21
Faerun to Tristram = 57
Faerun to Arbre = 83
Norrath to Straylight = 9
Norrath to Tristram = 50
Norrath to Arbre = 60
Straylight to Tristram = 27
Straylight to Arbre = 81
Tristram to Arbre = 90`
