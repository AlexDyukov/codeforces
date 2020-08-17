package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://codeforces.com/problemset/problem/617/A
// комментарии к решению
// Т.к. слоник может ходить на 1,2,3,4 или 5 позиций,
// то наикратчайший путь будет, когда слоник прошагает большую часть пути
// шагами по 5 позиций, а последний рывок сделает за один шаг требуемой длины
func taskSolution(n int) int {
	return n/5 + min(n%5, 1)
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
