package main

import (
	"fmt"
	"sort"
)

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://codeforces.com/problemset/problem/492/B
// комментарии к решению
// Для решения нам необходимо найти максимальное расстояние между соседними фонарями,
// с учетом границ улицы, и поделить его на 2
// Т.к. фонари идут не по порядку, то сортируем входные данные
func taskSolution(a []int, l int) string {
	sort.Sort(sort.IntSlice(a))
	maxDiff := max(a[0], l-a[len(a)-1]) * 2
	for i := 1; i < len(a); i += 1 {
		maxDiff = max(maxDiff, a[i]-a[i-1])
	}
	ans := fmt.Sprintf("%.10f", float64(maxDiff)/float64(2))
	return ans
}

func main() {
	var n, l int
	if _, err := fmt.Scan(&n, &l); err != nil {
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

	fmt.Println(taskSolution(a, l))
}
