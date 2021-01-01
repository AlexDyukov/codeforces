package main

import "fmt"

func min(a, b int) int {
	if a > b {
		return b
	}
	return a
}

// https://codeforces.com/problemset/problem/466/A
// комментарии к решению
// Считаем выгоду от приобритения абонемента в долгосрочной перспективе
// и в краткосрочной перспективе на остаток дней
func taskSolution(n, m, a, b int) int {
	subscriptionLong := n / m * min(a*m, b)
	subscriptionShort := min((n%m)*a, b)
	return subscriptionLong + subscriptionShort
}

func main() {
	var n, m, a, b int
	if _, err := fmt.Scan(&n, &m, &a, &b); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, m, a, b))
}
