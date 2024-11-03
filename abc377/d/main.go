package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	n, m := ReadInt2()

	d := make([]int, m)
	for i := 0; i < n; i++ {
		l, r := ReadInt2()
		l--
		r--
		d[r] = Max(d[r], l+1)
	}

	for r := 1; r < m; r++ {
		d[r] = Max(d[r], d[r-1])
	}

	c := 0
	for r := 0; r < m; r++ {
		c += r - d[r] + 1
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

func ReadInt3() (int, int, int) {
	in.Scan()
	abc := strings.Split(in.Text(), " ")
	if len(abc) != 3 {
		panic("invalid input")
	}
	a, _ := strconv.Atoi(abc[0])
	b, _ := strconv.Atoi(abc[1])
	c, _ := strconv.Atoi(abc[2])
	return a, b, c
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

func ReadString() string {
	in.Scan()
	return in.Text()
}

func ReadStrings() []string {
	in.Scan()
	return strings.Split(in.Text(), " ")
}

func Min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func Max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

func Abs(a int) int {
	if a > 0 {
		return a
	}
	return -a
}
