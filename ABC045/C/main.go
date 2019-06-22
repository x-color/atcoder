package main

import (
	"fmt"
	"strconv"
)

func main() {
	var s string
	fmt.Scan(&s)

	fmt.Println(dfs(s, 0, uint(len(s))-1))
}

func dfs(s string, f int, d uint) int {
	if d == 0 {
		return calc(s, f)
	}
	return dfs(s, f, d-1) + dfs(s, f|(1<<(d-1)), d-1)
}

func calc(s string, f int) int {
	v, j := 0, 0
	for i := range s {
		if f&1 == 1 {
			n, _ := strconv.Atoi(s[j : i+1])
			v += n
			j = i + 1
		}
		f >>= 1
	}
	n, _ := strconv.Atoi(s[j:])
	v += n
	return v
}
