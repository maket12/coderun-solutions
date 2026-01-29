package main

import (
	"bufio"
	"fmt"
	"os"
)


type Operation struct {
	l, r      int64
	oldLength int64
}

func main() {
	var n, q int
	var s string
	
	reader := bufio.NewReader(os.Stdin)
	writer := bufio.NewWriter(os.Stdout)
	defer writer.Flush()

	fmt.Fscan(reader, &n, &q)
	fmt.Fscan(reader, &s)

	history := make([]Operation, 0)
	currentLength := int64(n)

	for k := 0; k < q; k++ {
		var opType int
		fmt.Fscan(reader, &opType)

		if opType == 1 {
			var l, r int64
			fmt.Fscan(reader, &l, &r)
			// Archive operation
			history = append(history, Operation{
				l:         l,
				r:         r,
				oldLength: currentLength,
			})
			currentLength += (r - l + 1)
		} else {
			var i int64
			fmt.Fscan(reader, &i)

			// Backtracking
			for j := len(history) - 1; j >= 0; j-- {
				op := history[j]
				// If i is in the new part
				if i > op.oldLength {
					offset := i - op.oldLength
					i = op.l + offset - 1
				}
			}
			fmt.Fprintln(writer, string(s[i-1]))
		}
	}
}