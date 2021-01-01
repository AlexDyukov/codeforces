package main

import (
	"fmt"
	"sort"
)

type test struct {
	a []int
}

const TEST_MAX_SIZE = 50

func is_editable(a []int) bool {
	sort.Ints(a)
	for i := 1; i < len(a); i += 1 {
		if a[i]-a[i-1] > 1 {
			return false
		}
	}
	return true
}

// https://codeforces.com/problemset/problem/1399/A
// комментарий к решению
// Т.к. от нас не требуется выводить результирующий элемент массива,
// то достаточно проверить, что из массива можно сделать неубывающий
// массив с разницей соседних любых двух элементов <= 1
func taskSolution(in []test) string {
	ans := make([]byte, 0, len(in)*4) //YES\n == 4 byte
	for i := 0; i < len(in); i += 1 {
		if is_editable(in[i].a) {
			ans = append(ans, []byte("YES\n")...)
		} else {
			ans = append(ans, []byte("NO\n")...)
		}
	}
	return string(ans[:len(ans)-1])
}

func main() {
	var t int
	if _, err := fmt.Scan(&t); err != nil {
		panic(err)
	}

	in := make([]test, 0, t)
	for i := 0; i < t; i += 1 {
		var n, ai int
		a := make([]int, 0, TEST_MAX_SIZE)
		if _, err := fmt.Scan(&n); err != nil {
			panic(err)
		}
		for j := 0; j < n; j += 1 {
			if _, err := fmt.Scan(&ai); err != nil {
				panic(err)
			}
			a = append(a, ai)
		}
		in = append(in, test{a})
	}

	fmt.Println(taskSolution(in))
}
