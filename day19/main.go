// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"math"
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
	for i, kv := range m {
		invm[i] = [2]string{kv[1], kv[0]}
	}
	dp := make(map[string]int)
	var replace func(string, int) int
	replace = func(medicine string, steps int) int {
		if len(medicine) > len(target) {
			return 0
		}
		if medicine == "e" {
			fmt.Println("steps", steps)
			return steps
		}
		if steps, ok := dp[medicine]; ok {
			return steps
		}
		minSteps := math.MaxInt
		for molecule := range distinctMolecules(medicine, invm) {
			if _, ok := dp[molecule]; !ok {
				dp[molecule] = replace(molecule, steps+1)
			}

			v := dp[molecule]
			if v < minSteps {
				minSteps = v
			}
		}
		dp[medicine] = minSteps
		return minSteps
	}

	return replace(target, 0)
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
