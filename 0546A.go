package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://codeforces.com/problemset/problem/546/A
// комментарии к решению
// Конечная сумма для покупки бананов = сумма геометрической прогрессии S
// сумма которой не хватает бойцу = max(S-n, 0), т.к. в минус мы уйти не можем
func taskSolution(k, n, w int) int {
	total := k*(w+1)*w/2 - n
	return max(total, 0)
}

func main() {
	var k, n, w int
	if _, err := fmt.Scan(&k, &n, &w); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(k, n, w))
}
