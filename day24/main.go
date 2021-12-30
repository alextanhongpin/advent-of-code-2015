// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"sort"
	"strconv"
	"strings"
)

func main() {

	boxes := []int{1, 2, 3, 4, 5, 7, 8, 9, 10, 11}
	fmt.Println("example 1:", minQuantumEntanglement(3, boxes)) // 99
	fmt.Println("example 2:", minQuantumEntanglement(4, boxes)) // 44

	boxes = parse(input)
	fmt.Println("part 1:", minQuantumEntanglement(3, boxes)) // 10439961859
	fmt.Println("part 2:", minQuantumEntanglement(4, boxes)) // 72050269
}

func minQuantumEntanglement(n int, boxes []int) int {
	all := make([][]int, 0)
	_ = combinations(boxes, len(boxes)-1, sum(boxes)/n, nil, &all)

	sort.Slice(all, func(i, j int) bool {
		lhs, rhs := all[i], all[j]

		byLen := len(lhs) - len(rhs)
		byQuantumEntanglement := product(lhs) - product(rhs)
		return sortBy(byLen, byQuantumEntanglement)
	})

	return product(all[0])
}

func combinations(boxes []int, i, capacity int, selected []int, all *[][]int) int {
	if capacity == 0 {
		sort.Ints(selected)
		*all = append(*all, selected)
		return 1
	}
	if i < 0 || capacity < 0 {
		return 0
	}
	newSelected := copyInt(selected)
	newSelected = append(newSelected, boxes[i])
	return combinations(boxes, i-1, capacity-boxes[i], newSelected, all) +
		combinations(boxes, i-1, capacity, selected, all)
}

type Set []int

func (s Set) Map() map[int]bool {
	m := make(map[int]bool)
	for _, v := range s {
		m[v] = true
	}
	return m
}

func (s Set) Difference(other []int) []int {
	m := s.Map()

	for _, v := range other {
		delete(m, v)
	}

	var res []int
	for k := range m {
		res = append(res, k)
	}
	return res
}

func (s Set) Union(other []int) []int {
	m := s.Map()
	for _, v := range other {
		m[v] = true
	}
	var res []int
	for k := range m {
		res = append(res, k)
	}
	return res
}

func (s Set) Intersect(other []int) []int {
	m := s.Map()
	var res []int
	for _, v := range other {
		if m[v] {
			res = append(res, v)
		}
	}
	return res
}

func copyInt(src []int) []int {
	tgt := make([]int, len(src))
	copy(tgt, src)
	return tgt
}

func sum(n []int) int {
	res := 0
	for _, v := range n {
		res += v
	}
	return res
}

func product(n []int) int {
	res := 1
	for _, v := range n {
		res *= v
	}
	return res
}

func sortBy(sc ...int) bool {
	for _, c := range sc {
		if c != 0 {
			return c < 0
		}
	}
	return true
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

func parse(input string) []int {
	lines := strings.Split(input, "\n")
	res := make([]int, len(lines))
	for i, s := range lines {
		res[i] = toInt(s)
	}
	return res
}

var input = `1
3
5
11
13
17
19
23
29
31
37
41
43
47
53
59
67
71
73
79
83
89
97
101
103
107
109
113`
