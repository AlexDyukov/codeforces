package main

import (
	"fmt"
	"sort"
)

const alen = 4

// https://codeforces.com/problemset/problem/1154/A
// комментарии к решению
// Пусть s1 <= s2 <= s3 <= s4 и
// a + b = s1
// b + c = s2
// a + c = s3
// a + b + c = s4
// Тогда
// a = s4 - s2
// b = s4 - s3
// c = s4 - s1
func taskSolution(in []int) string {
	sort.Ints(in)
	s1, s2, s3, s4 := in[0], in[1], in[2], in[3]
	a, b, c := s4-s2, s4-s3, s4-s1
	return fmt.Sprint(a, b, c)
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
