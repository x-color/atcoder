package main

import (
	"bufio"
	"fmt"
	"math"
	"os"
	"strconv"
)

const MOD = 1000000007

func main() {
	defer out.Flush()

	n, b, k := ReadInt3()
	c := ReadInts(k)

	// N = 10^18
	// log2(N) = 62
	log2 := 62

	// mod[i] = 10^(2^i) % b
	mod := make([]int, log2)
	mod[0] = 10 % b
	for i := 1; i < len(mod); i++ {
		mod[i] = (mod[i-1] * mod[i-1]) % b
	}

	// ddp is a table for Doubling DP
	// このテーブルは2^i桁で余りがr個の数の個数を保持する
	// 本来は1~N桁の余りの数を保持する必要がある。しかし、全ての数は２の累乗で表現できるため、
	// 全ての桁の代わりに2の累乗のデータだけを保持する。
	// これにより、データ数と計算量をlog2(N)まで削減可能
	ddp := make([][]int, log2)
	for i := range ddp {
		ddp[i] = make([]int, b)
	}

	for _, m := range c {
		ddp[0][m%b]++
	}

	// dp[i + j]は数値を前後半に分けた際に、前半が2^i桁、後半が2^j桁である数に対するデータを保持する。
	// このデータは、2^i桁の数の個数と2^j桁の数の個数を掛け合わせることで算出可能。
	// dp[i + j][(p * (10j % B) + q) % B] += dp[i][p] * dp[j][q]
	for i := 1; i < log2; i++ {
		ddp[i] = mul(ddp[i-1], ddp[i-1], mod[i-1], b)
	}

	// resはn桁の数の個数を保持する。
	// 例） n = 12桁の場合
	// 12 = 8 + 4 = 2^3 + 2^2
	// 2^3桁の数の個数と2^2桁の数の個数を掛け合わせることで、12桁の数の個数を求める。
	res := make([]int, b)
	res[0] = 1
	for i := 0; i < log2; i++ {
		if n&(1<<i) > 0 {
			res = mul(res, ddp[i], mod[i], b)
		}
	}

	// bの倍数が求めるべき答えなので、余りが0の個数を出力する
	fmt.Fprintln(out, res[0])
}

func mul(dpi []int, dpj []int, tj int, b int) []int {
	res := make([]int, b)
	for p := 0; p < b; p++ {
		for q := 0; q < b; q++ {
			res[(p*tj+q)%b] += dpi[p] * dpj[q]
			res[(p*tj+q)%b] %= MOD
		}
	}
	return res
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
