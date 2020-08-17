package main

import "fmt"

// https://codeforces.com/problemset/problem/158/A
// комментарии к решению
// Входной массив отсортирован и его длина <= 100,
// поэтому решаем задачу в лоб:
// Если k-тый участник набрал 0 баллов, то ищем ближайшего не нулевого слева
// Если k-тый участник набрал более 0 балов, то ищем крайнего правого с таким же кол-вом баллов
func taskSolution(n int, k int, a []int) int {
	if a[k-1] > 0 {
		for i := k; i < len(a); i += 1 {
			if a[i] != a[k-1] {
				return i
			}
		}
		return len(a)
	}
	for i := k; i > 0; i -= 1 {
		if a[i-1] != 0 {
			return i
		}
	}
	return 0
}

func main() {
	var n, k int
	if _, err := fmt.Scan(&n, &k); err != nil {
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

	fmt.Println(taskSolution(n, k, a))
}
