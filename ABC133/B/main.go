package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	s := wordScanner()

	var n, d int
	scanInts(s, &n, &d)

	l := make([][]int, n)
	for i := 0; i < n; i++ {
		l[i] = make([]int, d)
		for j := 0; j < d; j++ {
			scanInts(s, &l[i][j])
		}
	}

	c := 0
	for i, a := range l[:len(l)] {
		for _, b := range l[i+1:] {
			if dist(a, b) {
				c++
			}
		}
	}

	fmt.Println(c)
}

func dist(x, y []int) bool {
	d2 := 0.0
	for i := range x {
		c := x[i] - y[i]
		d2 += float64(c) * float64(c)
	}

	tmp := int(math.Sqrt(d2))
	if d2 == float64(tmp)*float64(tmp) {
		return true
	}
	return false
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
