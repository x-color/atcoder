package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
)

type Home struct {
	x       int
	y       int
	visited bool
}

type Point struct {
	v    int
	home *Home
}

func main() {
	defer out.Flush()

	n, m := ReadInt2()
	sx, sy := ReadInt2()
	homes := make([]Home, n)
	for i := 0; i < n; i++ {
		x, y := ReadInt2()
		homes[i] = Home{x, y, false}
	}

	d := make([]string, m)
	c := make([]int, m)
	for i := 0; i < m; i++ {
		d[i], c[i] = ReadStringInt()
	}

	x2y := make(map[int][]Point)
	y2x := make(map[int][]Point)
	for i, h := range homes {
		if _, ok := x2y[h.x]; !ok {
			x2y[h.x] = []Point{}
		}
		if _, ok := y2x[h.y]; !ok {
			y2x[h.y] = []Point{}
		}
		x2y[h.x] = append(x2y[h.x], Point{
			v:    h.y,
			home: &homes[i],
		})
		y2x[h.y] = append(y2x[h.y], Point{
			v:    h.x,
			home: &homes[i],
		})
	}
	for i := range x2y {
		sort.Slice(x2y[i], func(j, k int) bool {
			return x2y[i][j].v < x2y[i][k].v
		})
	}
	for i := range y2x {
		sort.Slice(y2x[i], func(j, k int) bool {
			return y2x[i][j].v < y2x[i][k].v
		})
	}

	ans := 0
	for i, op := range d {
		visits := 0
		switch op {
		case "U":
			visits, x2y[sx] = visitedHomes(x2y[sx], sy, c[i])
			sy += c[i]
		case "D":
			visits, x2y[sx] = visitedHomes(x2y[sx], sy, -c[i])
			sy -= c[i]
		case "L":
			visits, y2x[sy] = visitedHomes(y2x[sy], sx, -c[i])
			sx -= c[i]
		case "R":
			visits, y2x[sy] = visitedHomes(y2x[sy], sx, c[i])
			sx += c[i]
		}
		ans += visits
	}

	fmt.Fprintln(out, sx, sy, ans)
}

func visitedHomes(homes []Point, p, d int) (int, []Point) {
	from, to := p, p+d
	if from > to {
		from, to = to, from
	}
	s := BinarySearch(len(homes), -1, func(i int) bool {
		return homes[i].v >= from
	})
	g := BinarySearch(len(homes), -1, func(i int) bool {
		return homes[i].v > to
	})

	c := 0
	if s > g {
		s, g = g, s
	}
	for i := s; i < g; i++ {
		if !homes[i].home.visited {
			c++
			homes[i].home.visited = true
		}
	}

	return c, append(homes[:s], homes[g:]...)
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

// `ok` must be the maximum value that satisfies the condition.
// `ng` must be the minmum value that does not satisfy the condition.
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