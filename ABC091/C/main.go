package main

import (
	"bufio"
	"fmt"
	"os"
	"sort"
	"strconv"
)

type Point struct {
	x int
	y int
}

type Points []Point

func (l Points) Len() int {
	return len(l)
}

func (l Points) Less(i, j int) bool {
	return l[i].x < l[j].x
}

func (l Points) Swap(i, j int) {
	l[i], l[j] = l[j], l[i]
}

func main() {
	s := wordScanner()

	var n int
	scanInts(s, &n)

	reds := make(Points, n)
	for i := range reds {
		scanInts(s, &reds[i].x, &reds[i].y)
	}

	blues := make(Points, n)
	for i := range blues {
		scanInts(s, &blues[i].x, &blues[i].y)
	}

	sort.Sort(reds)
	sort.Sort(blues)

	c := 0
	for _, b := range blues {
		i := -1
		for j, r := range reds {
			if isInner(r, b) {
				if i == -1 || reds[i].y < reds[j].y {
					i = j
				}
			}
		}
		if i != -1 {
			c++
			reds = append(reds[:i], reds[i+1:]...)
		}
	}
	fmt.Println(c)
}

func isInner(inner, outer Point) bool {
	return outer.x > inner.x && outer.y > inner.y
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
