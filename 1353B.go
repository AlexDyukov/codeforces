package main

import (
	"fmt"
	"sort"
	"strings"
)

type test struct {
	n int
	k int
	a []int
	b []int
}

// https://codeforces.com/problemset/problem/1353/B
// комментарий к решению
// Небольшой и разнообразные входной поток требует
// задействовать fmt.Scan , т.к. с парсингом ioutil.ReadAll
// сложность и количество кода увеличивается
func taskSolution(in []test) string {
	ans := make([]int, len(in))
	for i, ei := range in {
		sort.Ints(ei.a)
		sort.Sort(sort.Reverse(sort.IntSlice(ei.b)))
		for j := 0; j < ei.k && ei.a[j] < ei.b[j]; j += 1 {
			ei.a[j] = ei.b[j]
		}
		for _, aj := range ei.a {
			ans[i] += aj
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
		var n, k, ai, bi int
		var a, b []int
		if _, err := fmt.Scan(&n, &k); err != nil {
			panic(err)
		}
		for j := 0; j < n; j += 1 {
			if _, err := fmt.Scan(&ai); err != nil {
				panic(err)
			}
			a = append(a, ai)
		}
		for j := 0; j < n; j += 1 {
			if _, err := fmt.Scan(&bi); err != nil {
				panic(err)
			}
			b = append(b, bi)
		}
		in = append(in, test{n: n, k: k, a: a, b: b})
	}

	fmt.Println(taskSolution(in))
}
