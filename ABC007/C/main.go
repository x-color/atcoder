package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

type Pos struct {
	x int
	y int
	d int
}

var field map[int]map[int]bool
var arrived map[int]map[int]bool

var nexts []*Pos
var gx, gy int

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	r, c := getNums(sc)
	sy, sx := getNums(sc)
	gy, gx = getNums(sc)

	field, arrived = make(map[int]map[int]bool, r), make(map[int]map[int]bool, r)
	for i := 1; i <= r; i++ {
		field[i], arrived[i] = make(map[int]bool, c), make(map[int]bool, c)
		sc.Scan()
		p := sc.Text()
		for j := 1; j <= c; j++ {
			field[i][j] = p[j-1] == byte('.')
		}
	}

	nexts = []*Pos{&Pos{sx, sy, 0}}
	arrived[sy][sx] = true
	for {
		if d := distance(nexts[0]); d > 0 {
			fmt.Println(d)
			return
		}
		nexts = nexts[1:]
	}
}

func getNums(sc *bufio.Scanner) (int, int) {
	sc.Scan()
	a, _ := strconv.Atoi(sc.Text())
	sc.Scan()
	b, _ := strconv.Atoi(sc.Text())
	return a, b
}

func distance(p *Pos) int {
	if p.x == gx && p.y == gy {
		return p.d
	}
	for _, v := range []struct{ x, y int }{{1, 0}, {0, 1}, {-1, 0}, {0, -1}} {
		nx, ny := p.x+v.x, p.y+v.y
		if !field[ny][nx] || arrived[ny][nx] {
			continue
		}
		arrived[ny][nx] = true
		nexts = append(nexts, &Pos{nx, ny, p.d + 1})
	}
	return 0
}
