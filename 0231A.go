package main

import (
	"fmt"
)

// https://codeforces.com/problemset/problem/231/A
// комментарии к решению
// Т.к. кол-во друзей известно заранее,
// то и критерий условия можно посчитать заранее
// поэтому всё сводиться к подсчету кол-во элементов
// входного потока из заданного множества
func taskSolution(a []string) (ans int) {
	type s struct{}
	exists := s{}
	winnerLines := map[string]s{"1 1 0": exists, "1 0 1": exists, "0 1 1": exists, "1 1 1": exists}

	for _, line := range a {
		if _, ok := winnerLines[line]; ok {
			ans += 1
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	var in []string
	for i := 0; i < n; i += 1 {
		var line string
		if _, err := fmt.Scan(&line); err != nil {
			panic(err)
		}
		in = append(in, line)
	}

	fmt.Println(taskSolution(in))
}
