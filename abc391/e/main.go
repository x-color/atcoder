package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

func main() {
	defer out.Flush()

	n := ReadInt()
	s := ReadString()

	a := make([][]int, n+1)
	a[0] = make([]int, len(s))
	for i := range a[0] {
		if s[i] == '1' {
			a[0][i] = 1
		}
	}

	for i := 0; i < n; i++ {
		a[i+1] = make([]int, len(a[i])/3)
		for j := 0; j < len(a[i]); j += 3 {
			sum := a[i][j] + a[i][j+1] + a[i][j+2]
			if sum >= 2 {
				a[i+1][j/3] = 1
			}
		}
	}

	fmt.Fprintln(out, dfs(a, a[n][0], 0, n))
}

func dfs(tree [][]int, target int, n, depth int) int {
	if depth == 0 {
		if tree[depth][n] == target {
			return 1
		}
		return 0
	}
	counts := []int{0, 0, 0}
	if tree[depth-1][n*3] == target {
		counts[0] = dfs(tree, target, n*3, depth-1)
	}
	if tree[depth-1][n*3+1] == target {
		counts[1] = dfs(tree, target, n*3+1, depth-1)
	}
	if tree[depth-1][n*3+2] == target {
		counts[2] = dfs(tree, target, n*3+2, depth-1)
	}
	sort.Ints(counts)
	return counts[0] + counts[1]
}

// --- utils ---
// --- io ---
var (
	in  = bufio.NewScanner(os.Stdin)
	out = bufio.NewWriter(os.Stdout)
)

func init() {
	in.Buffer([]byte{}, math.MaxInt64)
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

func Min[T Number](v ...T) T {
	min := v[0]
	for _, n := range v {
		if n < min {
			min = n
		}
	}
	return min
}

func Max[T Number](v ...T) T {
	max := v[0]
	for _, n := range v {
		if n > max {
			max = n
		}
	}
	return max
}

func Abs[T Number](a T) T {
	if a > 0 {
		return a
	}
	return -a
}

func IsPrime(n int) bool {
	if n == 1 {
		return false
	}

	for i := 2; i*i <= n; i++ {
		if n%i == 0 {
			return false
		}
	}

	return true
}

func PrimeFactor(n int) map[int]int {
	pf := make(map[int]int)
	for i := 2; i*i <= n; i++ {
		c := 0
		for n%i == 0 {
			c++
			n /= i
		}
		pf[i] = c
	}

	if n != 1 {
		pf[n]++
	}
	return pf
}

// --- search ---

type Queue[T any] []T

func (q *Queue[T]) Push(v T) {
	*q = append(*q, v)
}

func (q *Queue[T]) Pop() T {
	v := (*q)[0]
	*q = (*q)[1:]
	return v
}

func Bfs[T any](first []T, f func(node T) []T) {
	q := append(Queue[T]{}, first...)
	for len(q) > 0 {
		node := q.Pop()
		for _, v := range f(node) {
			q.Push(v)
		}
	}
}

// `ok` must be the max/min value that satisfies the condition.
// `ng` must be the min/max value that does not satisfy the condition.
// The range must be (ok, ng] or [ng, ok).
// The return value is the same value of `v` that satisfies the condition. (`v-1` does not satisfy the condition).
// `check` is a function that checks if the value satisfies the condition.
func BinarySearch(ok, ng int, check func(v int) bool) int {
	for Abs(ok-ng) > 1 {
		m := (ok + ng) / 2
		if check(m) {
			ok = m
		} else {
			ng = m
		}
	}

	return ok
}

// --- board ---

type Square[T any] struct {
	value   T
	visited bool
}

type Position struct {
	x int
	y int
}

type Board[T any] [][]Square[T]

func NewBoard[T any](h, w int, f func(rune) T) Board[T] {
	board := make(Board[T], h)
	for i := 0; i < h; i++ {
		s := ReadString()
		board[i] = make([]Square[T], w)
		for j := 0; j < w; j++ {
			board[i][j] = Square[T]{
				value: f(rune(s[j])),
			}
		}
	}
	return board
}

func (b Board[T]) Width() int {
	return len(b[0])
}

func (b Board[T]) Height() int {
	return len(b)
}

func (b Board[T]) In(p Position) bool {
	return !(p.x < 0 || p.x >= b.Width() || p.y < 0 || p.y >= b.Height())
}

func (b Board[T]) At(p Position) Square[T] {
	return b[p.y][p.x]
}

func (b Board[T]) Set(p Position, v Square[T]) {
	b[p.y][p.x] = v
}

func (b Board[T]) Visit(p Position) {
	b[p.y][p.x].visited = true
}

func (b Board[T]) Leave(p Position) {
	b[p.y][p.x].visited = false
}

func (b Board[T]) IsVisited(p Position) bool {
	return b[p.y][p.x].visited
}

func (b Board[T]) Neighbors4(pos Position) []Position {
	return []Position{
		{x: pos.x + 1, y: pos.y},
		{x: pos.x - 1, y: pos.y},
		{x: pos.x, y: pos.y + 1},
		{x: pos.x, y: pos.y - 1},
	}
}

func (b Board[T]) Neighbors8(pos Position) []Position {
	return []Position{
		{x: pos.x + 1, y: pos.y},
		{x: pos.x - 1, y: pos.y},
		{x: pos.x, y: pos.y + 1},
		{x: pos.x, y: pos.y - 1},
		{x: pos.x + 1, y: pos.y + 1},
		{x: pos.x - 1, y: pos.y - 1},
		{x: pos.x - 1, y: pos.y + 1},
		{x: pos.x + 1, y: pos.y - 1},
	}
}
