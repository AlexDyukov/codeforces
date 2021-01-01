package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/337/A
// комментарии к решению
// кол-во учеников n <= кол-ва пазлов m, поэтому можно пройтись по массиву окном n
// и найти минимальную разницу между самым сложным и самым простым пазлом в этом окне
// предварительно отсортировать массив, конечно
func taskSolution(a []int, n int) int {
	sort.Ints(a)
	ans := a[n-1] - a[0]
	for i := n; i < len(a); i += 1 {
		diff := a[i] - a[i-n+1]
		if diff < ans {
			ans = diff
		}
	}
	return ans
}

func main() {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic(err)
	}

	a := []int{}
	for i := 0; i < m; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a, n))
}
