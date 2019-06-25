package main

import (
	"fmt"
)

func main() {
	var n, k int
	fmt.Scan(&n, &k)

	var s string
	fmt.Scan(&s)

	fmt.Println(betterSort(s, s, n, k))
}

func betterSort(s, t string, n, k int) string {
	for i := 0; i < n; i++ {
		a := i
		for j := i + 1; j < n; j++ {
			if t[j] < t[a] && count(s, swap(t, i, j)) <= k {
				a = j
			}
		}
		if a > i {
			t = swap(t, i, a)
		}
	}

	return t
}

func swap(s string, i, j int) string {
	return s[:i] + s[j:j+1] + s[i+1:j] + s[i:i+1] + s[j+1:]
}

func count(s, t string) int {
	c := 0
	for i := range s {
		if s[i] != t[i] {
			c++
		}
	}
	return c
}
