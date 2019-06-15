package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	sc := bufio.NewScanner(os.Stdin)
	sc.Split(bufio.ScanWords)
	sc.Scan()
	n, _ := strconv.Atoi(sc.Text())

	monsters := make([]int, n)
	for i := 0; i < n; i++ {
		sc.Scan()
		a, _ := strconv.Atoi(sc.Text())
		monsters[i] = a
	}

	a := monsters[0]
	for _, b := range monsters[1:] {
		a = gcd(a, b)
	}

	fmt.Println(a)
}

func gcd(x, y int) int {
	for y != 0 {
		z := x % y
		x, y = y, z
	}
	return x
}
