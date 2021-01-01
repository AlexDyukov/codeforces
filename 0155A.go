package main

import "fmt"

// https://codeforces.com/problemset/problem/155/A
// комментарии к решению
// Т.к. чисел не много и ограничение сверху целых 2 сек,
// то используем fmt.Scan для парсинга входного потока
func taskSolution(a []int) (ans int) {
	min := a[0]
	max := a[0]
	for i := 1; i < len(a); i += 1 {
		if min > a[i] {
			min, ans = a[i], ans+1
		} else if max < a[i] {
			max, ans = a[i], ans+1
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []int{}
	for i := 0; i < n; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a))
}
