package main

import "fmt"

// https://codeforces.com/problemset/problem/500/A
// комментарии к решению
// Т.к. в задаче не предусмотрено передвижение кроме как через порталы,
// то нам достаточно проитерироваться по массиву и по прыжкам попытаться
// добраться до нужного нам элемента
func taskSolution(a []int, n int) string {
	i := 1
	for i <= len(a) && i < n {
		i += a[i-1]
	}
	if i == n {
		return "YES"
	}
	return "NO"
}

func main() {
	var n, t int
	if _, err := fmt.Scan(&n, &t); err != nil {
		panic(err)
	}

	a := []int{}
	for i := 0; i < n-1; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a, t))
}
