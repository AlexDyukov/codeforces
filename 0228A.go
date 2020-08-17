package main

import "fmt"

const alen = 4

// https://codeforces.com/problemset/problem/228/A
// комментарии к решению
// Требуется подсчитать кол-во неуникальных элементов,
// поэтому уместно воспользоваться множествами set (map[T]struct{} в go)
func taskSolution(a []int) int {
	type s struct{}
	exists := s{}
	numbers := map[int]s{}
	for _, ai := range a {
		numbers[ai] = exists
	}
	return len(a) - len(numbers)
}

func main() {
	a := []int{}
	for i := 0; i < alen; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a))
}
