package main

import (
	"fmt"
	"strings"
)

type numpair struct {
	a int
	b int
}

// https://codeforces.com/problemset/problem/1360/A
// комментарий к решению
// Для получения минимальной площади квадрата необходимо
// параллельно класть прямоугольники вдоль большей стороны
func taskSolution(in []numpair) string {
	ans := make([]int, len(in))
	for i := 0; i < len(in); i += 1 {
		a, b := in[i].a, in[i].b
		if a > b {
			a, b = b, a
		}
		if 2*a > b {
			ans[i] = 4 * a * a
		} else {
			ans[i] = b * b
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

	in := make([]numpair, 0, t)
	for i := 0; i < t; i += 1 {
		var a, b int
		if _, err := fmt.Scan(&a, &b); err != nil {
			panic(err)
		}
		in = append(in, numpair{a: a, b: b})
	}

	fmt.Println(taskSolution(in))
}
