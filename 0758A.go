package main

import "fmt"

// https://codeforces.com/problemset/problem/758/A
// комментарии к решению
// Сначала находим максимум, а после суммируем разницу
// каждого элемента массива и максимального
func taskSolution(a []int) (ans int) {
	max := 0
	for _, ai := range a {
		if max < ai {
			max = ai
		}
	}
	for _, ai := range a {
		ans += max - ai
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []int{}
	for i := 0; i < n; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a))
}
