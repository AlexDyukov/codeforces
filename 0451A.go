package main

import "fmt"

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://codeforces.com/problemset/problem/451/A
// комментарии к решению
// Можно заметить, что независимо от выбранного пересечения,
// на поле после хода всегда остаётся сетка из (n-1) x (m-1) палок,
// поэтому ещё до начала игры можно определить кто проиграет.
// Т.к. играет всего двое и порядок ходов задан заранее, то
// определить програвшего можно по четности минимального из n или m
func taskSolution(n int, m int) string {
	if min(n, m)%2 == 0 {
		return "Malvika"
	}
	return "Akshat"
}

func main() {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, m))
}
