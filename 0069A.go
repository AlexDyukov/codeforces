package main

import "fmt"

type point struct {
	x int
	y int
	z int
}

// https://codeforces.com/problemset/problem/69/A
// комментарии к решению
// В виду низких требований к входным данным == кол-во чисел <= 100
// задача решается в лоб считыванием int через fmt.Scan и простым алгоритмом:
// Считываем все точки -> суммируем соответствующие координаты -> проверяем на нуль.
// Можно реализовать подсчёт координат как работу с массивом,
// но со структурой выглядит понятнее
func taskSolution(a []point) string {
	res := point{}
	for _, p := range a {
		res.x += p.x
		res.y += p.y
		res.z += p.z
	}

	if res.x == 0 && res.y == 0 && res.z == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []point{}
	for i := 0; i < n; i += 1 {
		var x, y, z int
		if _, err := fmt.Scan(&x, &y, &z); err != nil {
			panic(err)
		}
		a = append(a, point{x, y, z})
	}

	fmt.Println(taskSolution(a))
}
