package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

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

func main() {
	var n int
	s := wordScanner(100)
	scanInts(s, &n)

	a := scanIntSlice(s, n)
	b := scanIntSlice(s, n)
	c := scanIntSlice(s, n-1)

	ans := 0
	j := -1
	for _, i := range a {
		ans += b[i-1]
		if i == j+1 {
			ans += c[j-1]
		}
		j = i
	}
	fmt.Println(ans)
}
