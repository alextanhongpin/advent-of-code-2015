// You can edit this code!
// Click here and start typing.
package main

import (
	"crypto/md5"
	"fmt"
	"io"
	"strings"
)

func main() {
	fmt.Println(hash("bgvyzdsv", 5))
	fmt.Println(hash("bgvyzdsv", 6))
}

func hash(password string, zeros int) int {
	zero := strings.Repeat("0", zeros)
	var n int
	for {
		h := md5.New()
		io.WriteString(h, fmt.Sprintf("%s%d", password, n))
		hash := fmt.Sprintf("%x", h.Sum(nil))
		if strings.HasPrefix(hash, zero) {
			return n
		}
		n++
	}
}
