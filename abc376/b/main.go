package main

import (
	"bufio"
	"os"
	"strconv"
	"strings"
)

func main() {

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
