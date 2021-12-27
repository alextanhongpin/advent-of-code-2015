// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"regexp"
	"strconv"
	"strings"
)

var re *regexp.Regexp

func init() {
	var err error
	re, err = regexp.Compile(`-?\d+`)
	if err != nil {
		panic(err)
	}
}

func combinations(inputs [][]string) []string {
	for len(inputs) > 1 {
		var l, r []string
		l, r, inputs = inputs[0], inputs[1], inputs[2:]
		var res []string
		for _, k := range l {
			for _, v := range r {
				res = append(res, k+":"+v)
			}
		}
		inputs = append(inputs, res)
	}
	return inputs[0]
}

func main() {
	fmt.Println("part 1:", bestScore(input, false))
	fmt.Println("part 2:", bestScore(input, true))
}

type recipe struct {
	capacity, durability, flavor, texture, calories int
}

func bestScore(input string, withFixedCalories bool) int {
	lines := strings.Split(input, "\n")

	recipes := make([]recipe, len(lines))
	for i, line := range lines {
		m := re.FindAllString(line, -1)
		recipes[i] = recipe{
			capacity:   toInt(m[0]),
			durability: toInt(m[1]),
			flavor:     toInt(m[2]),
			texture:    toInt(m[3]),
			calories:   toInt(m[4]),
		}
	}
	n := len(recipes)
	deltas := make([][]string, n)
	for i := 0; i < n; i++ {
		var rng []string
		for j := -25; j < 25; j++ {
			rng = append(rng, strconv.Itoa(j))
		}
		deltas[i] = rng

	}
	combs := combinations(deltas)

	type score struct {
		distr []int
		delta []int
	}

	var stack []score
	for _, comb := range combs {
		distr := make([]int, n)
		for i := 0; i < n; i++ {
			distr[i] = 100 / n
		}

		parts := strings.Split(comb, ":")
		delta := make([]int, n)
		var sum int
		for i, s := range parts {
			delta[i] = toInt(s)
			sum += delta[i]
		}
		if sum != 0 {
			continue
		}
		stack = append(stack, score{distr, delta})
	}

	bestScore := 0
	for len(stack) > 0 {
		var s score
		s, stack = stack[0], stack[1:]
		sum := 0
		for i := 0; i < len(s.distr); i++ {
			s.distr[i] += s.delta[i]
			sum += s.distr[i]
		}
		mix := recipe{}
		for i, recipe := range recipes {
			mix.capacity += recipe.capacity * s.distr[i]
			mix.durability += recipe.durability * s.distr[i]
			mix.flavor += recipe.flavor * s.distr[i]
			mix.texture += recipe.texture * s.distr[i]
			mix.calories += recipe.calories * s.distr[i]
		}
		if withFixedCalories && mix.calories != 500 {
			continue
		}
		newScore := max(mix.capacity, 0) *
			max(mix.durability, 0) *
			max(mix.flavor, 0) *
			max(mix.texture, 0)

		if newScore > bestScore {
			bestScore = newScore
			stack = append(stack, s)
		}
	}
	return bestScore
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

type tuple struct{ i, j int }

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `Butterscotch: capacity -1, durability -2, flavor 6, texture 3, calories 8
Cinnamon: capacity 2, durability 3, flavor -2, texture -1, calories 3`

var input = `Sprinkles: capacity 5, durability -1, flavor 0, texture 0, calories 5
PeanutButter: capacity -1, durability 3, flavor 0, texture 0, calories 1
Frosting: capacity 0, durability -1, flavor 4, texture 0, calories 6
Sugar: capacity -1, durability 0, flavor 0, texture 2, calories 8`
