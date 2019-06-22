package main

import (
	"fmt"
)

func main() {
	var a, b, c, d int
	fmt.Scan(&a, &b, &c, &d)

	a--
	n := lcm(c, d)

	ac := a / c
	ad := a / d
	acd := a / n

	bc := b / c
	bd := b / d
	bcd := b / n

	fmt.Println((b - a) - (bc + bd - bcd) + (ac + ad - acd))
}

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
}
