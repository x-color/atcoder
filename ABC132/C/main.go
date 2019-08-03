package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

func main() {
	var n int
	s := wordScanner()
	scanInts(s, &n)

	l := make([]int, n)
	for i := 0; i < n; i++ {
		scanInts(s, &l[i])
	}

	sort.Ints(l)
	fmt.Println(l[n/2] - l[n/2-1])
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
