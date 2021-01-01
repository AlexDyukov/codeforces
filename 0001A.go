package main

import "fmt"

// https://codeforces.com/problemset/problem/1/A
// комментарии к решению
// Задача имеет простой алгоритм:
// 1) Найти ближайшие (сверху) кратные А стороны квадрата
// 2) Замостить квадратами со стороной А, учитывая возможное переполнение
func taskSolution(n, m, a int64) int64 {
	n = topDividend(n, a) / a
	m = topDividend(m, a) / a
	return n // m
}

func topDividend(a, b int64) int64 {
	if a%b == 0 {
		return a
	}
	return a - a%b + b
}

func main() {
	var n, m, a int64
	if _, err := fmt.Scan(&n, &m, &a); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, m, a))
}
