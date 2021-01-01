package main

import "fmt"

// https://codeforces.com/problemset/problem/977/A
// комментарии к решению
// Числа во входных параметрах не особо большие, поэтому
// можно в лоб реализовать Танино вычитание, что и
// сделано ниже, за тем лишь исключением, что добавлена
// оптимизация для вычитания последнего разряда
func taskSolution(n int64, k int64) int64 {
	for k > 0 {
		remainder := n % 10
		switch {
		case remainder == 0:
			k -= 1
			n /= 10
		case remainder < k:
			k -= remainder
			n -= remainder
		default:
			return n - k
		}
	}
	return n
}

func main() {
	var n, k int64
	if _, err := fmt.Scan(&n, &k); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, k))
}
