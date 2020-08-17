package main

import "fmt"

// https://codeforces.com/problemset/problem/1030/A
// комментарии к решению
// В виду низких требований к входным данным == n(1≤n≤100)
// задача решается в лоб:
// считываем числа и проверяем входной массив на не нулевое значение
// При измении условия в сторону max(n) > 1000,
// будет выгоднее перейти к []byte и сравнению посимвольно с '1'
// т.к. семейство fmt.Scan функций сильно проседают по производительности
func taskSolution(a []int) string {
	for _, ai := range a {
		if ai > 0 {
			return "HARD"
		}
	}
	return "EASY"
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
