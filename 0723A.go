package main

import "fmt"

func sorted(a, b, c int) (int, int, int) {
	if a > b {
		a, b = b, a
	}
	if b > c {
		b, c = c, b
	}
	if a > b {
		a, b = b, a
	}
	return a, b, c
}

// https://codeforces.com/problemset/problem/723/A
// комментарии к решению
// Оптимальная точка встречи - это средний дом.
func taskSolution(a, b, c int) int {
	a, b, c = sorted(a, b, c)
	return c - a // (b-a) + (c-b)
}

func main() {
	var x1, x2, x3 int
	if _, err := fmt.Scan(&x1, &x2, &x3); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(x1, x2, x3))
}
