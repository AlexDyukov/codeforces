package main

import "fmt"

// https://codeforces.com/problemset/problem/677/A
// комментарии к решению
// Минимальная ширина дороги будет как минимум не меньше кол-ва друзей
// и за каждого друга, выше забора она должна быть увеличена на 1,
// т.к. высоким друзьям придётся нагинаться
func taskSolution(a []int, h int) int {
	ans := len(a)
	for _, ai := range a {
		if ai > h {
			ans += 1
		}
	}
	return ans
}

func main() {
	var n, h int
	if _, err := fmt.Scan(&n, &h); err != nil {
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

	fmt.Println(taskSolution(a, h))
}
