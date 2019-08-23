package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n, k int
	s := wordScanner(100)
	scanInts(s, &n, &k)
	x := scanIntSlice(s, n)

	ans := 1 << 30
	for l := 0; l < n-k+1; l++ {
		r := l + k - 1
		ans = min(ans, min(abs(x[r]), abs(x[l]))+abs(x[r]-x[l]))
	}
	fmt.Println(ans)
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
