package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

func main() {
	defer out.Flush()

	n := ReadInt()
	loads := make(map[int][]int)
	for i := 0; i < n-1; i++ {
		a, b := ReadInt2()
		if _, ok := loads[a]; !ok {
			loads[a] = []int{}
		}
		if _, ok := loads[b]; !ok {
			loads[b] = []int{}
		}
		loads[a] = append(loads[a], b)
		loads[b] = append(loads[b], a)
	}

	// u is the farthest node from 1
	u, _ := dfs(loads, make(map[int]bool), 1, 0)
	// d is a distance from u to the farthest node
	// d is a diameter of the tree
	_, d := dfs(loads, make(map[int]bool), u, 0)

	fmt.Fprintln(out, d+1)
}

// dfs returns node and distance
func dfs(loads map[int][]int, visited map[int]bool, cur int, dis int) (int, int) {
	if visited[cur] {
		return 0, 0
	}

	visited[cur] = true

	node := cur
	max := dis
	for _, next := range loads[cur] {
		n, d := dfs(loads, visited, next, dis+1)
		if d > max {
			node = n
			max = d
		}
	}

	return node, max
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
