package main

import "fmt"

// https://codeforces.com/problemset/problem/50/A
// комментарии к решению
// Задача на алгоритм по замощению прямоугольника фиксированной фигурой
func taskSolution(m, n int) int {
	return m // n / 2
}

func main() {
	var m, n int
	if _, err := fmt.Scan(&m, &n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(m, n))
}
