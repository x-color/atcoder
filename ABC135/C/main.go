package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	var n int
	s := wordScanner()
	scanInts(s, &n)
	a := scanIntSlice(s, n+1)
	b := scanIntSlice(s, n)

	ans := 0
	for i, bi := range b {
		if d := bi - a[i]; d > 0 {
			if d > a[i+1] {
				ans += a[i] + a[i+1]
				a[i+1] = 0
				continue
			} else {
				a[i+1] -= d
			}
		}
		ans += bi
	}
	fmt.Println(ans)
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

func scanIntSlice(s *bufio.Scanner, n int) []int {
	vals := make([]int, n)
	for i := range vals {
		s.Scan()
		m, _ := strconv.Atoi(s.Text())
		vals[i] = m
	}
	return vals
}
