package main

import "fmt"

// https://codeforces.com/problemset/problem/144/A
// комментарии к решению
// В виду низких требований к входным данным == кол-во чисел <= 100
// задача решается в лоб считыванием int через fmt.Scan и простым алгоритмом:
// Находим где стоит самый высокий ближе к началу
// Находим где стоит самый низкий ближе к концу
// суммируем с проверкой, что если нам понадобиться менять местами
// самого высокого и низкого, то выиграем на этом секунду
func taskSolution(a []int) int {
	maxi, mini := 0, len(a)-1
	for i, ai := range a {
		if ai > a[maxi] {
			maxi = i
		}
		if ai <= a[mini] {
			mini = i
		}
	}
	if maxi < mini {
		return maxi + len(a) - mini - 1
	}
	return maxi + len(a) - mini - 2
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
