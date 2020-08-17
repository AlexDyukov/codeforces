package main

import "fmt"

// https://codeforces.com/problemset/problem/791/A
// комментарии к решению
// В задаче требуется найти такое минимальное n, что
// a*(3**n) < b*(2**n)
// при таких небольших a и b как в задаче проще перебрать,
// чем пытаться выводить формулу с логарифмами
func taskSolution(a, b int) (ans int) {
	for a <= b {
		a *= 3
		b *= 2
		ans += 1
	}
	return ans
}

func main() {
	var a, b int
	if _, err := fmt.Scan(&a, &b); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(a, b))
}
