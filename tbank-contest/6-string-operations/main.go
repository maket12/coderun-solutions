package main

import (
	"bufio"
	"fmt"
	"os"
)

type Operation struct {
	l, r, length int64
}

func main() {
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	var n, q int
	fmt.Fscan(reader, &n, &q)

	var s string
	fmt.Fscan(reader, &s)

	history := make([]Operation, 0)

	for k := 0; k < q; k++ {
		var t int
		fmt.Fscan(reader, &t)

		if t == 1 {
			var l, r int64
			fmt.Fscan(reader, &l, &r)
			history = append(history, Operation{l, r, r - l + 1})
		} else {
			var i int64
			fmt.Fscan(reader, &i)

			// Backtracking
			for j := len(history) - 1; j >= 0; j-- {
				_, r, subLen := history[j].l, history[j].r, history[j].length
				
				if i > r + subLen {
					i -= subLen
				} else if i > r {
					i -= subLen
				}
			}
			fmt.Fprintln(writer, string(s[i-1]))
		}
	}
}