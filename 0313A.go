package main

import "fmt"

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://codeforces.com/problemset/problem/4/A
// комментарии к решению
// n без последней цифры == n%10
// n без предпоследней цифры == (n/100)*10 + n%10
func taskSolution(n int) int {
	ans := max(n, n/10)
	n = (n/100)*10 + n%10
	return max(ans, n)
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
