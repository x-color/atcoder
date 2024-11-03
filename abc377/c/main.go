package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type Pos struct {
	y int
	x int
}

func main() {
	board := map[int64]bool{}

	n, m := ReadInt2()
	pos := make([]Pos, m)
	for i := 0; i < m; i++ {
		a, b := ReadInt2()
		pos[i] = Pos{a - 1, b - 1}
		board[int64(pos[i].x+pos[i].y*n)] = true
	}

	for _, p := range pos {
		attack(board, Pos{p.y + 2, p.x + 1}, n)
		attack(board, Pos{p.y + 1, p.x + 2}, n)
		attack(board, Pos{p.y - 1, p.x + 2}, n)
		attack(board, Pos{p.y - 2, p.x + 1}, n)
		attack(board, Pos{p.y - 2, p.x - 1}, n)
		attack(board, Pos{p.y - 1, p.x - 2}, n)
		attack(board, Pos{p.y + 1, p.x - 2}, n)
		attack(board, Pos{p.y + 2, p.x - 1}, n)
	}

	fmt.Println(n*n - len(board))
}

func attack(board map[int64]bool, pos Pos, n int) {
	if pos.x < 0 || pos.y < 0 || pos.x >= n || pos.y >= n {
		return
	}
	board[int64(pos.x+pos.y*n)] = true
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
