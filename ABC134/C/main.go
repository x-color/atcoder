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
	var n int
	s := wordScanner()
	scanInts(s, &n)
	a := make([]int, n)
	for i := 0; i < n; i++ {
		scanInts(s, &a[i])
	}

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

func scanInts(s *bufio.Scanner, vals ...*int) {
	for i := range vals {
		s.Scan()
		n, _ := strconv.Atoi(s.Text())
		*vals[i] = n
	}
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
