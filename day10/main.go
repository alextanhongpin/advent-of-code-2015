// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
	"strconv"
	"strings"
)

func main() {
	input := "1113222113"
	n := 40
	for i := 0; i < n; i++ {
		input = lookAndSay(input)
	}
	// 40 iter: 252594
	// 50 iter: 3579328
	fmt.Println(len(input))
}

func lookAndSay(input string) string {
	var sb strings.Builder
	var n string
	var count int
	for _, ch := range input {
		r := string(ch)
		if n == r {
			count++
		} else {
			if count != 0 {
				sb.WriteString(strconv.Itoa(count) + n)
			}
			count = 1
			n = r
		}
	}
	if count > 0 {
		sb.WriteString(strconv.Itoa(count) + n)
	}
	return sb.String()
}
