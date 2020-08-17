package main

import "fmt"

// https://codeforces.com/problemset/problem/479/A
// комментарий к решению
// кол-во чисел заранее известно = 3
// поэтому мы можем просто посчитать все возможные произведения
// и выбрать среди них максимальное, которое и будет ответом
func taskSolution(a, b, c int) (max int) {
	results := []int{
		a + b + c,
		a*b + c,
		a + b*c,
		a * b * c,
		a * (b + c),
		(a + b) * c,
	}
	for _, item := range results {
		if item > max {
			max = item
		}
	}
	return max
}

func main() {
	var a, b, c int
	if _, err := fmt.Scan(&a, &b, &c); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(a, b, c))
}
