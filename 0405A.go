package main

import (
	"fmt"
	"sort"
)

// https://codeforces.com/problemset/problem/405/A
// комментарии к решению
// если рассматривать поле не как совокупность кубиков, поставленных друг на друга
// а как список вертикальных столбцов, то легко понять, что с нас просто требуют
// отсортировать входной массив в порядке возрастания
// Ограничения на входные данные достаточно низкие, поэтому можно использовать fmt.Scan
func taskSolution(a []int) string {
	sort.Ints(a)
	res := fmt.Sprint(a)
	return res[1 : len(res)-1]
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
