package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	s := wordScanner()
	n, a := scanNAndIntSlice(s)

	var l, r [100001]int
	for i := 0; i < n-1; i++ {
		l[i+1] = gcd(l[i], a[i])
	}
	for i := n - 1; i > 0; i-- {
		r[i] = gcd(r[i+1], a[i])
	}

	var ans, m int
	for i := 0; i < n; i++ {
		m = gcd(l[i], r[i+1])
		ans = max(ans, m)
	}
	fmt.Println(ans)
}

func wordScanner() *bufio.Scanner {
	s := bufio.NewScanner(os.Stdin)
	s.Split(bufio.ScanWords)
	return s
}

func scanNAndIntSlice(s *bufio.Scanner) (int, []int) {
	s.Scan()
	n, _ := strconv.Atoi(s.Text())
	return n, scanIntSlice(s, n)
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

func gcd(x, y int) int {
	if x > y {
		x, y = y, x
	}
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
}

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}
