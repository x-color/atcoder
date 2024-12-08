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

	h, w, d := ReadInt3()
	s := make([]string, h)
	for i := 0; i < h; i++ {
		s[i] = ReadString()
	}

	office := make([][]bool, h)
	for i := 0; i < h; i++ {
		office[i] = make([]bool, w)
		for j := 0; j < w; j++ {
			office[i][j] = s[i][j] == '#'
		}
	}

	max := 0
	for i := 0; i < h*w-1; i++ {
		if office[i/w][i%w] {
			continue
		}
		for j := i + 1; j < h*w; j++ {
			if office[j/w][j%w] {
				continue
			}
			c := 0
			for k := 0; k < h*w; k++ {
				if office[k/w][k%w] {
					continue
				}
				if Abs(k/w-i/w)+Abs(k%w-i%w) <= d || Abs(k/w-j/w)+Abs(k%w-j%w) <= d {
					c++
				}
			}
			if c > max {
				max = c
			}
		}
	}

	fmt.Fprintln(out, max)
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
