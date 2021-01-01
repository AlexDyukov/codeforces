package main

import "fmt"

// https://codeforces.com/problemset/problem/136/A
// комментарии к решению
// Малое ограничение на кол-во входных данных n (1 ≤ n ≤ 100)
// позволяет манипулировать числами напрямую.
// А для поддержания концепции "бизнес логики" в taskSolution
// используем fmt.Sprint и подрезаем []
func taskSolution(a []int) string {
	res := make([]int, len(a))
	for i, ai := range a {
		res[ai-1] = i + 1
	}
	output := fmt.Sprint(res)
	// trim []
	return output[1 : len(output)-1]
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
