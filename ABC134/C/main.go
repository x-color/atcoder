package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Set struct {
	before int
	after  int
}

func main() {
	s := wordScanner()
	n, a := scanNAndInts(s)

	var bmax, amax int
	l := make([]Set, n)
	for i := 1; i < n; i++ {
		bmax = max(a[i-1], bmax)
		amax = max(a[n-i], amax)
		l[i].before = bmax
		l[n-i-1].after = amax
	}

	for i := range a {
		fmt.Println(max(l[i].before, l[i].after))
	}
}

func wordScanner() *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	return s
}

func scanNAndInts(s *bufio.Scanner) (int, []int) {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	vals := make([]int, n)
	for i := 0; i < n; i++ {
		s.Scan()
		m, _ := strconv.Atoi(s.Text())
		vals[i] = m
	}
	return n, vals
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
