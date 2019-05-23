package main

import "fmt"

var N = 1000000007

func main() {
	var n, m int
	fmt.Scanf("%d %d", &n, &m)
	a := factorization(m)
	r := 1
	for _, v := range a {
		r = r * comb(v+n-1, v) % N
	}
	fmt.Println(r)
}

func factorization(n int) map[int]int {
	a := make(map[int]int)
	for i := 2; i*i <= n; {
		if n%i == 0 {
			a[i]++
			n /= i
		} else {
			i++
		}
	}
	if n > 1 {
		a[n]++
	}
	return a
}

func pow(x, n int) int {
	r := 1
	for ; n > 0; n /= 2 {
		if n%2 == 1 {
			r = r * x % N
		}
		x = x * x % N
	}
	return r
}

func comb(n, m int) int {
	return factorial(n, n-m+1) * pow(factorial(m, 2), N-2) % N
}

func factorial(n, m int) int {
	r := 1
	for i := m; i <= n; i++ {
		r = r * i % N
	}
	return r
}
