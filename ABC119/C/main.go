package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

var L []int
var N, A, B, C int

func main() {
	s := wordScanner(100)
	scanInts(s, &N, &A, &B, &C)
	L = scanIntSlice(s, N)

	fmt.Println(dfs(N-1, 0, 0, 0))
}

func dfs(i, a, b, c int) int {
	if i < 0 {
		if min(a, b, c) == 0 {
			return 9999
		}
		return abs(A-a) + abs(B-b) + abs(C-c) - 30
	}
	r1 := dfs(i-1, a, b, c)
	r2 := dfs(i-1, a+L[i], b, c) + 10
	r3 := dfs(i-1, a, b+L[i], c) + 10
	r4 := dfs(i-1, a, b, c+L[i]) + 10
	return min(r1, r2, r3, r4)
}

func min(l ...int) int {
	r := l[0]
	for _, n := range l {
		if n < r {
			r = n
		}
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func wordScanner(n int) *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	b := make([]byte, n)
	s.Buffer(b, n)
	return s
}

func scanInts(s *bufio.Scanner, vals ...*int) {
	for i := range vals {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		*vals[i] = n
	}
}

func scanIntSlice(s *bufio.Scanner, n int) []int {
	vals := make([]int, n)
	for i := range vals {
		s.Scan()
		m, _ := strconv.Atoi(s.Text())
		vals[i] = m
	}
	return vals
}
