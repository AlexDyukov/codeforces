package main

import "fmt"

// https://codeforces.com/problemset/problem/486/A
// комментарии к решению
// в виду большого n посчитать функцию итеративно очень затратно,
// поэтому выписав значения для первых 10 (к примеру) n,
// определим закономерности и реализуем их в коде
func taskSolution(n int64) int64 {
	if n%2 == 0 {
		return n / 2
	}
	return -(n + 1) / 2
}

func main() {
	var n int64
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
