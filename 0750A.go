package main

import "fmt"

const timeLimit = 240
const k = 5

// https://codeforces.com/problemset/problem/750/A
// комментарии к решению
// Т.к. числа маленькие, то проще проитерироваться, чем искать
// ответ по формуле суммы арифметической последовательности
func taskSolution(n int, m int) int {
	leftTime := timeLimit - m
	ans := 0
	for ans < n && leftTime >= (ans+1)*k {
		ans += 1
		leftTime -= ans * k
	}
	return ans
}

func main() {
	var n, k int
	if _, err := fmt.Scan(&n, &k); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, k))
}
