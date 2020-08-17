package main

import (
	"fmt"
	"strings"
)

type pair struct {
	n int
	m int
}

// https://codeforces.com/problemset/problem/1353/A
// комментарий к решению
// Проход через fmt.Scan сильно медленнее решения на ioutil.ReadAll
// однако в лимиты укладываемся (0.5сек при лимитах в 1сек)
func taskSolution(in []pair) string {
	ans := make([]int, len(in))
	for i := 0; i < len(in); i += 1 {
		n, m := in[i].n, in[i].m
		switch n {
		case 1:
			ans[i] = 0
		case 2:
			ans[i] = m
		default:
			ans[i] = 2 * m
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

	in := make([]pair, 0, t)
	for i := 0; i < t; i += 1 {
		var n, m int
		if _, err := fmt.Scan(&n, &m); err != nil {
			panic(err)
		}
		in = append(in, pair{n: n, m: m})
	}

	fmt.Println(taskSolution(in))
}
