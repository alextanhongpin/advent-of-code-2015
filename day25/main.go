// You can edit this code!
// Click here and start typing.
package main

import (
	"fmt"
)

func main() {
	row, col := 2981, 3075
	dp := make(map[int]map[int]int)
	dp[1] = make(map[int]int)
	for i := 1; i <= col*2; i++ {
		dp[1][i] = dp[1][i-1] + i
	}
	for j := 2; j <= row; j++ {
		if _, ok := dp[j]; !ok {
			dp[j] = make(map[int]int)
		}
		for i := 1; i <= col*2; i++ {
			dp[j][i] = dp[j-1][i+1] - 1
		}
	}

	code := 20151125
	n := dp[row][col]
	for i := 1; i < n; i++ {
		code *= 252533
		code %= 33554393
	}

	fmt.Println(code) // 9132360
}
