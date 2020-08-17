package main

import "fmt"

// https://codeforces.com/problemset/problem/427/A
// комментарии к решению
// Для решения нам необходимо на каждом событии знать кол-во
// свободных полицейских, а далее мы просто итерируемся по массиву,
// идентифицируя каждое событие как один из 3-х возможных вариантов
func taskSolution(a []int) (ans int) {
	for i, free_officers := 0, 0; i < len(a); i += 1 {
		if a[i] > 0 {
			free_officers += a[i]
		} else if free_officers > 0 {
			free_officers -= 1
		} else {
			ans += 1
		}
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
