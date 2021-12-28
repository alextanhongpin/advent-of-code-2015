// You can edit this code!
// Click here and start typing.
package main

import (
	"container/heap"
	"fmt"
	"strings"
)

func main() {
	start, m := parse(input)
	fmt.Println("part 1:", len(distinctMolecules(start, m)))
	fmt.Println("part 2:", fabricateMedicine(input))
}

func fabricateMedicine(input string) int {
	target, m := parse(input)
	invm := make([][2]string, len(m))
	// Inverted map: Instead of making the strings longer, we focus on the opposite.
	for i := range m {
		invm[i] = [2]string{m[i][1], m[i][0]}
	}
	pq := make(PriorityQueue, 1, 1e3)
	pq[0] = &Item{value: target}
	heap.Init(&pq)

	for pq.Len() > 0 {
		item := heap.Pop(&pq).(*Item)
		molecules := distinctMolecules(item.value, invm)
		for m := range molecules {
			if m == "e" {
				return item.steps + 1
			}
			heap.Push(&pq, &Item{value: m, steps: item.steps + 1})
		}
	}
	return -1
}

func distinctMolecules(start string, m [][2]string) map[string]bool {
	unq := make(map[string]bool)
	for i := range start {
		for _, kv := range m {
			k, v := kv[0], kv[1]
			if strings.HasPrefix(start[i:], k) {
				offset := len(k)
				key := strings.Join([]string{start[:i], v, start[i+offset:]}, "")
				unq[key] = true
			}
		}
	}
	return unq
}

// An Item is something we manage in a priority queue.
type Item struct {
	value string // The value of the item; arbitrary.
	steps int
}

// A PriorityQueue implements heap.Interface and holds Items.
type PriorityQueue []*Item

func (pq PriorityQueue) Len() int { return len(pq) }

func (pq PriorityQueue) Less(i, j int) bool {
	// We want Pop to give us the highest, not lowest, priority so we use greater than here.
	return len(pq[i].value) < len(pq[j].value)
}

func (pq PriorityQueue) Swap(i, j int) {
	pq[i], pq[j] = pq[j], pq[i]
}

func (pq *PriorityQueue) Push(x interface{}) {
	item := x.(*Item)
	*pq = append(*pq, item)
}

func (pq *PriorityQueue) Pop() interface{} {
	old := *pq
	n := len(old)
	item := old[n-1]
	old[n-1] = nil // avoid memory leak
	*pq = old[0 : n-1]
	return item
}

var example = `H => HO
H => OH
O => HH

HOHOHO`

var example2 = `e => H
e => O
H => HO
H => OH
O => HH

HOHOHO`

func parse(input string) (string, [][2]string) {
	lines := strings.Split(input, "\n")
	var res [][2]string
	start := ""
	next := false
	for _, line := range lines {
		line = strings.TrimSpace(line)
		if line == "" {
			next = true
			continue
		}
		if next {
			start = line
		} else {
			parts := strings.Split(line, " => ")
			res = append(res, [2]string{
				strings.TrimSpace(parts[0]),
				strings.TrimSpace(parts[1]),
			})
		}
	}
	return start, res
}

var input = `Al => ThF
Al => ThRnFAr
B => BCa
B => TiB
B => TiRnFAr
Ca => CaCa
Ca => PB
Ca => PRnFAr
Ca => SiRnFYFAr
Ca => SiRnMgAr
Ca => SiTh
F => CaF
F => PMg
F => SiAl
H => CRnAlAr
H => CRnFYFYFAr
H => CRnFYMgAr
H => CRnMgYFAr
H => HCa
H => NRnFYFAr
H => NRnMgAr
H => NTh
H => OB
H => ORnFAr
Mg => BF
Mg => TiMg
N => CRnFAr
N => HSi
O => CRnFYFAr
O => CRnMgAr
O => HP
O => NRnFAr
O => OTi
P => CaP
P => PTi
P => SiRnFAr
Si => CaSi
Th => ThCa
Ti => BP
Ti => TiTi
e => HF
e => NAl
e => OMg

CRnSiRnCaPTiMgYCaPTiRnFArSiThFArCaSiThSiThPBCaCaSiRnSiRnTiTiMgArPBCaPMgYPTiRnFArFArCaSiRnBPMgArPRnCaPTiRnFArCaSiThCaCaFArPBCaCaPTiTiRnFArCaSiRnSiAlYSiThRnFArArCaSiRnBFArCaCaSiRnSiThCaCaCaFYCaPTiBCaSiThCaSiThPMgArSiRnCaPBFYCaCaFArCaCaCaCaSiThCaSiRnPRnFArPBSiThPRnFArSiRnMgArCaFYFArCaSiRnSiAlArTiTiTiTiTiTiTiRnPMgArPTiTiTiBSiRnSiAlArTiTiRnPMgArCaFYBPBPTiRnSiRnMgArSiThCaFArCaSiThFArPRnFArCaSiRnTiBSiThSiRnSiAlYCaFArPRnFArSiThCaFArCaCaSiThCaCaCaSiRnPRnCaFArFYPMgArCaPBCaPBSiRnFYPBCaFArCaSiAl`
