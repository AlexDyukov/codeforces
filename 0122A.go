package main

import "fmt"

// https://codeforces.com/problemset/problem/122/A
// комментарии к решению
// Задача требует проверить, делится ли число на любое счастливое число без остатка
// Входные параметры в виде n (1 ≤ n ≤ 1000) позволяют перечислить все возможные
// взаимнопростые делители и проверить для них остаток от деления
func taskSolution(n int) string {
	type s struct{}
	exists := s{}
	lucky := map[int]s{
		4:   exists,
		7:   exists,
		47:  exists,
		74:  exists,
		447: exists,
		474: exists,
		477: exists,
		747: exists,
		774: exists,
	}
	for l := range lucky {
		if n%l == 0 {
			return "YES"
		}
	}
	return "NO"
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
