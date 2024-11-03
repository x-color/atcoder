package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type QR struct {
	q int
	r int
}

type TD struct {
	t int
	d int
}

func main() {
	n := ReadInt()
	qr := make([]QR, n)
	for i := 0; i < n; i++ {
		q, r := ReadInt2()
		qr[i] = QR{q, r}
	}

	q := ReadInt()
	td := make([]TD, q)
	for i := 0; i < q; i++ {
		t, d := ReadInt2()
		td[i] = TD{t, d}
	}

	ans := make([]string, q)
	for i, v := range td {
		m := v.d % qr[v.t-1].q
		if m <= qr[v.t-1].r {
			ans[i] = fmt.Sprintf("%d", v.d+(qr[v.t-1].r-m))
		} else {
			ans[i] = fmt.Sprintf("%d", v.d+qr[v.t-1].q-(m-qr[v.t-1].r))
		}
	}

	fmt.Println(strings.Join(ans, "\n"))
}

// --- utils ---
var (
	in = bufio.NewScanner(os.Stdin)
)

func ReadInt() int {
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	return n
}

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
