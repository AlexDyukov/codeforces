package main

import (
	"fmt"
	"sort"
)

type fight struct {
	x int
	y int
}

// https://codeforces.com/problemset/problem/230/A
// комментарии к решению
// Т.к. порядок сражений с драконами не задан заранее, то мы сортируем их от легкого к сложному,
// чтобы Кирито минимально рисковал своей жизнью в сражениях. Фактом победы считаем проитерированного дракона.
// Фактом прохождения подземелья - проитерирования всех драконов текущего уровня
func taskSolution(a []fight, s int) string {
	sort.SliceStable(a, func(i, j int) bool { return a[i].x < a[j].x })
	i := 0
	for i < len(a) && s > a[i].x {
		s += a[i].y
		i += 1
	}
	if i == len(a) {
		return "YES"
	}
	return "NO"
}

func main() {
	var s, n int
	if _, err := fmt.Scan(&s, &n); err != nil {
		panic(err)
	}

	a := []fight{}
	for i := 0; i < n; i += 1 {
		var x, y int
		if _, err := fmt.Scan(&x, &y); err != nil {
			panic(err)
		}
		a = append(a, fight{x, y})
	}

	fmt.Println(taskSolution(a, s))
}
