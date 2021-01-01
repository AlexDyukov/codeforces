package main

import "fmt"

// https://codeforces.com/problemset/problem/25/A
// комментарии к решению
// В виду низких требований к входным данным == кол-во чисел <= 100
// задача решается в лоб считыванием int через fmt.Scan и простым алгоритмом:
// если четность одного из первых трех чисел не совпадает, то возвращаем его номер (индекс + 1)
// иначе ищем ближайший номер, где четность не совпадает
func taskSolution(a []int) int {
	first, second, third := a[0]%2, a[1]%2, a[2]%2
	if !(first == second && second == third) {
		switch {
		case first == second:
			return 3
		case first == third:
			return 2
		default:
			return 1
		}
	}
	for i := 3; i < len(a); i += 1 {
		if a[i]%2 != first {
			return i + 1
		}
	}
	return 0
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
