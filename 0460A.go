package main

import "fmt"

// https://codeforces.com/problemset/problem/460/A
// комментарии к решению
// Маленькие числа во входном потоке и теги перебора позволяют пройтись
// в цикле по дням. Можно также заметить, что кол-во носков в каждый
// i/m и i/m+1 день неизменяется для i кратных m, поэтому
// итерируемся, пока счетчик кол-ва носков на утро не обнулится
func taskSolution(n int, m int) (ans int) {
	for n > 0 {
		ans += 1
		if ans%m != 0 {
			n -= 1
		}
	}
	return ans
}

func main() {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, m))
}
