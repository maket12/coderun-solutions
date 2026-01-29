package main

import (
	"fmt"
	"sort"
	"strings"
)

// Optimized solution (using more bult-in methods)
func main() {
	var s string
	_, _ = fmt.Scan(&s)

	arr := strings.Split(s, "")
	sort.Strings(arr)

	// If the first digit is zero, it will search for the next non-zero digit
	if arr[0] == "0" {
		for i := 1; i < len(arr); i++ {
			if arr[i] != "0" {
				arr[0], arr[i] = arr[i], arr[0]
				break
			}
		}
	}

	fmt.Println(strings.Join(arr, ""))
}