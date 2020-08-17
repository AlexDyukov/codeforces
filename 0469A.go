package main

import "fmt"

// https://codeforces.com/problemset/problem/469/A
// комментарии к решению
// В задаче достаточно большой объем кода для чтения данных:
// n
// m a[1] a[2] ... a[m]
// k b[1] b[2] ... b[k]
// Для решения необходимо проверить, эквивалентно ли
// объединение множеств A и B и множество последовательности 1 2 .. n
// для чего удобнее использовать set (map[T]struct{} в go)
func taskSolution(n int, a []int, b []int) string {
	type s struct{}
	exists := s{}
	levels := map[int]s{}
	for _, ai := range a {
		levels[ai] = exists
	}
	for _, bi := range b {
		levels[bi] = exists
	}
	if len(levels) == n {
		return "I become the guy."
	}
	return "Oh, my keyboard!"
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []int{}
	var alen int
	if _, err := fmt.Scan(&alen); err != nil {
		panic(err)
	}
	for i := 0; i < alen; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	b := []int{}
	var blen int
	if _, err := fmt.Scan(&blen); err != nil {
		panic(err)
	}
	for i := 0; i < blen; i += 1 {
		var bi int
		if _, err := fmt.Scan(&bi); err != nil {
			panic(err)
		}
		b = append(b, bi)
	}

	fmt.Println(taskSolution(n, a, b))
}
