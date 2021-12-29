// You can edit this code!
// Click here and start typing.
package main

import "fmt"

func main() {
	target := 33100000

	houses := make([]int, target/10)
	for i := 1; i < target/10; i++ {
		for j := i; j < len(houses); j += i {
			houses[j] += i * 10
		}
	}
	for i, h := range houses {
		if h >= target {
			fmt.Println("part 1:", i)
			break
		}
	}

	houses = make([]int, target/11)
	for i := 1; i < target/11; i++ {
		for j, n := i, 0; j < len(houses) && n < 50; j, n = j+i, n+1 {
			houses[j] += i * 11
		}
	}
	for i, h := range houses {
		if h >= target {
			fmt.Println("part 2:", i)
			break
		}
	}
}
