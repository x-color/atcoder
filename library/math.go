package main

func lcm(x, y int) int {
	return x * y / gcd(x, y)
}

func gcd(x, y int) int {
	for x%y != 0 {
		x, y = y, x%y
	}
	return y
}

func max(l ...int) int {
	r := l[0]
	for _, n := range l {
		if n > r {
			r = n
		}
	}
	return r
}

func min(l ...int) int {
	r := l[0]
	for _, n := range l {
		if n < r {
			r = n
		}
	}
	return r
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
