// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	reg := make(map[string]int)
	fmt.Println("part 1:", eval(input, reg))

	reg = make(map[string]int)
	reg["a"] = 1
	fmt.Println("part 2:", eval(input, reg))
}

func eval(input string, reg map[string]int) int {
	lines := strings.Split(input, "\n")
	i := 0
	for {
		if i >= len(lines) {
			break
		}

		inst := lines[i]
		op, rest := inst[:3], inst[4:]
		switch op {
		case "hlf":
			reg[rest] /= 2
		case "tpl":
			reg[rest] *= 3
		case "inc":
			reg[rest]++
		case "jmp":
			offset := toInt(rest)
			i += offset
			continue
		case "jie":
			parts := strings.Split(rest, ", ")
			vars, offset := parts[0], toInt(parts[1])
			if reg[vars]%2 == 0 {
				i += offset
				continue
			}
		case "jio":
			parts := strings.Split(rest, ", ")
			vars, offset := parts[0], toInt(parts[1])
			if reg[vars] == 1 {
				i += offset
				continue
			}
		}
		i++
	}

	return reg["b"]
}

func toInt(s string) int {
	n, err := strconv.Atoi(s)
	if err != nil {
		panic(err)
	}
	return n
}

var example = `inc a
jio a, +2
tpl a
inc a`

var input = `jio a, +18
inc a
tpl a
inc a
tpl a
tpl a
tpl a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
tpl a
tpl a
inc a
jmp +22
tpl a
inc a
tpl a
inc a
inc a
tpl a
inc a
tpl a
inc a
inc a
tpl a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
inc a
inc a
tpl a
jio a, +8
inc b
jie a, +4
tpl a
inc a
jmp +2
hlf a
jmp -7`
