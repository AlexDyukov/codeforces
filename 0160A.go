package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/160/A
// комментарии к решению
// В виду низких требований к входным данным == длина последовательности <= 100
// задача решается в лоб:
// 1) сортируем slice по убыванию
// 2) ищем позицию k такую, что sum(from 0 to k) > sum(from k to end)
func taskSolution(a []int) (ans int) {
	sort.Sort(sort.Reverse(sort.IntSlice(a)))
	sum := 0
	for _, e := range a {
		sum += e
	}
	for value := 0; ans < len(a) && (value <= sum-value); ans += 1 {
		value += a[ans]
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	in := []int{}
	for i := 0; i < n; i += 1 {
		var a int
		if _, err := fmt.Scan(&a); err != nil {
			panic(err)
		}
		in = append(in, a)
	}

	fmt.Println(taskSolution(in))
}
