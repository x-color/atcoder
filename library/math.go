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

func max(x, y int) int {
	if x > y {
		return x
	}
	return y
}

func min(x, y int) int {
	if x < y {
		return x
	}
	return y
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}
