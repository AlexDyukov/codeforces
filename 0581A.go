package main

import "fmt"

func abs(a int) int {
	if a < 0 {
		return -a
	}
	return a
}

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

// https://codeforces.com/problemset/problem/581/A
// комментарии к решению
// Для того, чтобы определить кол-во дней, во время которых Вася
// будет в разноцветных носках необходимо определить наименьшее
// из двух представленных чисел.
// Для того, чтобы определить кол-во дней, во время которых Вася
// будет в одноцветных (оставшихся) носках необходимо оставшиеся
// разбить на пары
func taskSolution(n int, m int) string {
	return fmt.Sprint(min(n, m), abs(n-m)/2)
}

func main() {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, m))
}
