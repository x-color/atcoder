package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n, x int
	s := wordScanner(100)
	scanInts(s, &n, &x)
	l := scanIntSlice(s, n)

	sort.Ints(l)
	d := make([]int, n)
	for i, xi := range l {
		d[i] = abs(x - xi)
	}

	ans := d[0]
	for _, di := range d {
		ans = gcd(ans, di)
	}

	fmt.Println(ans)
}

func gcd(x, y int) int {
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
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
