// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

var re *regexp.Regexp

func init() {
	var err error
	re, err = regexp.Compile(`\d+`)
	if err != nil {
		panic(err)
	}
}

func main() {
	t := 1000
	lines := strings.Split(example, "\n")
	reindeers := make([]reindeer, len(lines))

	var maxdist int
	for i, line := range lines {
		m := re.FindAllString(line, -1)
		reindeers[i] = reindeer{
			speed: toInt(m[0]),
			burst: toInt(m[1]),
			rest:  toInt(m[2]),
		}
		dist := reindeers[i].distance(t)
		if dist > maxdist {
			maxdist = dist
		}
	}
	fmt.Println("part 1:", maxdist) // 2640

	scores := make([]int, len(reindeers))
	for t0 := 1; t0 <= t; t0++ {
		bestReindeers := make(map[int][]int)
		bestDist := 0
		for i, reindeer := range reindeers {
			dist := reindeer.distance(t0)
			if dist > bestDist {
				bestDist = dist
			}
			bestReindeers[dist] = append(bestReindeers[dist], i)
		}
		fmt.Println(t0, bestDist, bestReindeers)
		for _, r := range bestReindeers[bestDist] {
			scores[r]++
		}
	}
	sort.Ints(scores)
	fmt.Println("part 2:", scores[len(scores)-1]) // 1114 too high
}

type reindeer struct {
	speed int
	rest  int
	burst int
}

func (r reindeer) distance(seconds int) int {
	e := seconds / (r.rest + r.burst)
	if seconds%(r.rest+r.burst) >= r.burst {
		e++
	}
	return e * r.speed * r.burst
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `Comet can fly 14 km/s for 10 seconds, but then must rest for 127 seconds.
Dancer can fly 16 km/s for 11 seconds, but then must rest for 162 seconds.`

var input = `Dancer can fly 27 km/s for 5 seconds, but then must rest for 132 seconds.
Cupid can fly 22 km/s for 2 seconds, but then must rest for 41 seconds.
Rudolph can fly 11 km/s for 5 seconds, but then must rest for 48 seconds.
Donner can fly 28 km/s for 5 seconds, but then must rest for 134 seconds.
Dasher can fly 4 km/s for 16 seconds, but then must rest for 55 seconds.
Blitzen can fly 14 km/s for 3 seconds, but then must rest for 38 seconds.
Prancer can fly 3 km/s for 21 seconds, but then must rest for 40 seconds.
Comet can fly 18 km/s for 6 seconds, but then must rest for 103 seconds.
Vixen can fly 18 km/s for 5 seconds, but then must rest for 84 seconds.`
