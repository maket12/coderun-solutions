package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	reader := bufio.NewReader(os.Stdin)
	var t int
	fmt.Fscan(reader, &t)

	for ; t > 0; t-- {
		var s string
		fmt.Fscan(reader, &s)
		n := int64(len(s))

		maxK := int64(0)
		currentK := int64(0)

		if !strings.Contains(s, "0") {
			fmt.Println(n * n)
			continue
		}

		s2 := s + s
		for _, char := range s2 {
			if char == '1' {
				currentK++
				if currentK > maxK {
					maxK = currentK
				}
			} else {
				currentK = 0
			}
		}
		
		if maxK > n {
			maxK = n
		}

		w := (maxK + 1) / 2
		h := maxK - w + 1
		fmt.Println(w * h)
	}
}