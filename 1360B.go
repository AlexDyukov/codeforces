package main

import (
	"fmt"
	"sort"
	"strings"
)

type test struct {
	a []int
}

const TEST_MAX_SIZE = 50

// https://codeforces.com/problemset/problem/1360/B
// комментарий к решению
// Т.к. размер конечных команд не определен, то можно разделять спортсменов
// хоть как 1 vs n-1 , поэтому наша задача сводится к разделению на
// самых слабых vs самых сильных так, чтобы разница между max(слабых) и min(сильных)
// была минимальной. Т.е. сортируем всех спротсменов по силе и ищем минимальную
// разницу между любыми соседними - это и будет ответом
func taskSolution(in []test) string {
	ans := make([]int, len(in))
	for i := 0; i < len(in); i += 1 {
		a := in[i].a
		sort.Ints(a)
		ans[i] = a[1] - a[0]
		for j := 2; j < len(a); j += 1 {
			if a[j]-a[j-1] < ans[i] {
				ans[i] = a[j] - a[j-1]
			}
		}
	}
	res := fmt.Sprint(ans)
	return strings.Replace(res[1:len(res)-1], " ", "\n", len(ans)-1)
}

func main() {
	var t int
	if _, err := fmt.Scan(&t); err != nil {
		panic(err)
	}

	in := make([]test, 0, t)
	for i := 0; i < t; i += 1 {
		var n, ai int
		a := make([]int, 0, TEST_MAX_SIZE)
		if _, err := fmt.Scan(&n); err != nil {
			panic(err)
		}
		for j := 0; j < n; j += 1 {
			if _, err := fmt.Scan(&ai); err != nil {
				panic(err)
			}
			a = append(a, ai)
		}
		in = append(in, test{a})
	}

	fmt.Println(taskSolution(in))
}
