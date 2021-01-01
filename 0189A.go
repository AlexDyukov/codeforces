package main

import "fmt"

// https://codeforces.com/problemset/problem/189/A
// комментарии к решению
// Задача решается полным перебором. Для сокращения кол-ва операций
// итерируемся с максимального возможного вниз, т.к. в цикле for
// вычисления из init производяться один раз (сокращается кол-во операций)
func taskSolution(n, a, b, c int) int {
	max := 0
	for i := n / a; i >= 0; i -= 1 {
		for j := (n - i*a) / b; j >= 0; j -= 1 {
			left := n - i*a - j*b
			if left%c == 0 && max < i+j+left/c {
				max = i + j + left/c
			}
		}
	}
	return max
}

func main() {
	var n, a, b, c int
	if _, err := fmt.Scan(&n, &a, &b, &c); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, a, b, c))
}
