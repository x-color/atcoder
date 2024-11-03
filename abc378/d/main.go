package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	defer out.Flush()

	h, w, k := ReadInt3()
	board := make([][]bool, h)
	for i := 0; i < h; i++ {
		s := ReadString()
		board[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			board[i][j] = s[j] == '.'
		}
	}

	routes := 0
	for y := 0; y < h; y++ {
		for x := 0; x < w; x++ {
			routes += dfs(board, x, y, k)
		}
	}

	fmt.Fprintln(out, routes)
}

func dfs(board [][]bool, x, y, k int) int {
	if x < 0 || x >= len(board[0]) || y < 0 || y >= len(board) {
		return 0
	}
	if !board[y][x] {
		return 0
	}
	if k == 0 {
		return 1
	}

	board[y][x] = false
	routes := 0
	routes += dfs(board, x+1, y, k-1)
	routes += dfs(board, x-1, y, k-1)
	routes += dfs(board, x, y+1, k-1)
	routes += dfs(board, x, y-1, k-1)
	board[y][x] = true
	return routes
}

// --- utils ---

// --- utils ---
// --- io ---
var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func init() {
	in.Split(bufio.ScanWords)
}

func ReadInt() int {
	in.Scan()
	n, _ := strconv.Atoi(in.Text())
	return n
}

func ReadInt2() (int, int) {
	return ReadInt(), ReadInt()
}

func ReadInt3() (int, int, int) {
	return ReadInt(), ReadInt(), ReadInt()
}

func ReadInts(n int) []int {
	l := make([]int, n)
	for i := 0; i < n; i++ {
		l[i] = ReadInt()
	}
	return l
}

func ReadString() string {
	in.Scan()
	return in.Text()
}

func ReadStrings(n int) []string {
	l := make([]string, n)
	for i := 0; i < n; i++ {
		l[i] = ReadString()
	}
	return l
}

func ReadStringInt() (string, int) {
	return ReadString(), ReadInt()
}

func PrintSlice[T any](l []T) {
	v := make([]any, len(l))
	for i, n := range l {
		v[i] = n
	}
	fmt.Fprintln(out, v...)
}

// --- math ---

type Number interface {
	int | int32 | int64 | float32 | float64
}

func Min[T Number](a, b T) T {
	if a < b {
		return a
	}
	return b
}

func Max[T Number](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func Abs[T Number](a T) T {
	if a > 0 {
		return a
	}
	return -a
}
