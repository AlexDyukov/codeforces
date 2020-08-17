package main

import "fmt"

type roll struct {
	m int
	c int
}

// https://codeforces.com/problemset/problem/703/A
// комментарии к решению
// Малый объем входных данных (n <= 100) позволяет задействовать fmt.Scan .
func taskSolution(a []roll) string {
	res := 0
	for _, r := range a {
		if r.m > r.c {
			res += 1
		} else if r.m < r.c {
			res -= 1
		}
	}
	if res > 0 {
		return "Mishka"
	} else if res < 0 {
		return "Chris"
	}
	return "Friendship is magic!^^"
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []roll{}
	for i := 0; i < n; i += 1 {
		var m, c int
		if _, err := fmt.Scan(&m, &c); err != nil {
			panic(err)
		}
		a = append(a, roll{m, c})
	}

	fmt.Println(taskSolution(a))
}
