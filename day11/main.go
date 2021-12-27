// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	input := "hxbxwxba"
	// input = "abcdefgh" // abcdffaa
	// input = "ghijklmn" // ghjaabcc
	valid := false
	for !valid {
		input, valid = newPassword(input)
	}
	fmt.Println("part 1:", input)
	
	valid = false
	for !valid {
		input, valid = newPassword(input)
	}
	fmt.Println("part 2:", input)
}

func newPassword(input string) (string, bool) {
	nums := make([]int, len(input))
	for i, ch := range input {
		nums[i] = ord(ch)
	}
	deny := map[int]bool{
		ord('i'): true,
		ord('o'): true,
		ord('l'): true,
	}

	// No i, o and l
	for i, n := range nums {
		if deny[n] {
			nums[i]++
			for j := i + 1; j < len(nums); j++ {
				nums[j] = ord('a')
			}
			break
		}
	}

	// Increment last
	i := len(nums) - 1
	for {
		nums[i] += 1
		if nums[i]/26 > 0 {
			nums[i] %= 26
			i--
			continue
		}
		break
	}

	// Has increasing straight
	hasIncreasingStraight := false
	for i, n := range nums {
		if i < len(nums)-2 {
			if n+1 == nums[i+1] && n+2 == nums[i+2] {
				hasIncreasingStraight = true
			}
		}
	}

	// Has at least two different, non-overlapping pairs.
	nonOverlappingPairs := 0
	lastIndex := 0
	pairs := make(map[string]bool)
	for i, n := range nums {
		if i < len(nums)-1 {
			l, r := string(char(n)), string(char(nums[i+1]))
			if l != r {
				continue
			}
			pair := l + r
			if pairs[pair] {
				if lastIndex+2 <= i {
					lastIndex = i
					continue
				}
				lastIndex = i
				delete(pairs, pair)
				nonOverlappingPairs--
				continue
			}
			pairs[pair] = true
			if lastIndex+2 <= i {
				nonOverlappingPairs++
			}
			lastIndex = i
		}
	}

	var res string
	for _, n := range nums {
		res += string(char(n))
	}
	return res, nonOverlappingPairs >= 2 && hasIncreasingStraight
}

func ord(r rune) int {
	return int(r - 'a')
}

func char(n int) rune {
	return rune(n + 'a')
}

func toInt(password string) int {
	var total int
	for i := 0; i < len(password); i++ {
		ord := rune(password[i]) - 'a' + 1
		total *= 26
		total += int(ord)
	}
	return total
}

func toPassword(n int) string {
	pwd := ""
	for n/26 > 0 {
		pwd = string('a'+n%26-1) + pwd
		n /= 26
	}
	if n > 0 {
		pwd = string('a'+n%26-1) + pwd
	}
	return pwd
}
