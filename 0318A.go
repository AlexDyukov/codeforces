package main

import "fmt"

// https://codeforces.com/problemset/problem/318/A
// комментарии к решению
// входные параметры достаточно большие, поэтому
// работаем по формуле: определяем четность/нечетность
// будущего числа и определяем его порядок
// к примеру для "7 7":
// 1 3 5 7 2 4 6 == получившаяся последовательность
// 1 2 3 4 5 6 7 == позиции
//       ^
//    середина
func taskSolution(n int64, k int64) int64 {
	if k > n/2+n%2 {
		k -= n/2 + n%2
		return k * 2
	}
	return (k-1)*2 + 1
}

func main() {
	var n, k int64
	if _, err := fmt.Scan(&n, &k); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, k))
}
