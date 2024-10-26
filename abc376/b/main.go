package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, q := ReadInt2()

	h := make([]string, q)
	t := make([]int, q)
	for i := 0; i < q; i++ {
		h[i], t[i] = ReadStringInt()
	}

	c := 0
	l, r := 1, 2
	for i := 0; i < q; i++ {
		from, to, mid := 0, t[i], 0
		if h[i] == "R" {
			from, mid = r, l
			r = t[i]
		} else {
			from, mid = l, r
			l = t[i]
		}
		if from > to {
			from, to = to, from
		}

		if from < mid && mid < to {
			c += n - (to - from)
		} else {
			c += to - from
		}
	}
	fmt.Println(c)
}

// --- utils ---
var (
	in = bufio.NewScanner(os.Stdin)
)

func ReadInt2() (int, int) {
	in.Scan()
	nm := strings.Split(in.Text(), " ")
	if len(nm) != 2 {
		panic("invalid input")
	}
	n, _ := strconv.Atoi(nm[0])
	m, _ := strconv.Atoi(nm[1])
	return n, m
}

func ReadInts() []int {
	in.Scan()
	l := strings.Split(in.Text(), " ")
	r := make([]int, len(l))
	for i, v := range l {
		n, _ := strconv.Atoi(v)
		r[i] = n
	}
	return r
}

func ReadStringInt() (string, int) {
	in.Scan()
	sn := strings.Split(in.Text(), " ")
	if len(sn) != 2 {
		panic("invalid input")
	}
	s := sn[0]
	m, _ := strconv.Atoi(sn[1])
	return s, m
}
