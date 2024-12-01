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

	n, x := ReadInt2()
	v := ReadInts(n)
	p := make([]float64, n+1)
	for i := 1; i <= n; i++ {
		p[i] = float64(v[i-1]) / 100.0
	}

	// `px` is a map to calculate a probability distribution
	// that what number of rare cards are in 1 pack
	// by the dynamic programming.
	px := make([][]float64, n+1)
	for i := range px {
		px[i] = make([]float64, n+1)
	}

	px[0][0] = 1.0
	for i := 1; i <= n; i++ {
		px[i][0] = (1.0 - p[i]) * px[i-1][0]
		for j := 1; j <= n; j++ {
			px[i][j] = p[i]*px[i-1][j-1] + (1.0-p[i])*px[i-1][j]
		}
	}

	// `e` is a map to calculate the expected number of pack
	// to get `x` rare cards by the dynamic programming.
	// ---
	// e.g. This is an expect value what number of packs need to open
	// 		to get 3 rare cards by packs including 3+ cards.
	// e[3] = (e[0] * px[3])	(Get 3 rare cards by 1 pack)
	// 		+ (e[1] * px[2])	(Get 2 rare cards by 1 pack)
	// 		+ (e[2] * px[1])	(Get 1 rare cards by 1 pack)
	// 		+ (e[3] * px[0])	(Don't get any rare cards by 1 pack)
	//		+ 1					(Count this pack)
	// => e[3] - (e[3] * px[0]) = (e[0] * px[3]) + (e[1] * px[2]) + (e[2] * px[1]) + 1
	// => e[3] = {(e[0] * px[3]) + (e[1] * px[2]) + (e[2] * px[1]) + 1} / (1 - px[0])
	e := make([]float64, x+1)
	for i := 1; i <= x; i++ {
		for j := 1; j <= n && j < i; j++ {
			e[i] += px[n][j] * e[i-j]
		}
		e[i] += 1.0
		e[i] /= 1.0 - px[n][0]
	}

	fmt.Fprintf(out, "%.10f\n", e[x])
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
	return p.x < 0 || p.x >= b.Width() || p.y < 0 || p.y >= b.Height()
}

func (b Board[T]) At(p Position) Square[T] {
	return b[p.y][p.x]
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
