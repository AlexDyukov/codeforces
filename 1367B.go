package main

import (
	"fmt"
	"strings"
)

type test struct {
	a []int
}

const TEST_MAX_SIZE = 40

func calc_wrongs(a []int) (wrong_even, wrong_odd int) {
	// even == четные
	// odd  == нечетные
	for i := 0; i < len(a)/2; i += 1 {
		if a[2*i]%2 != 0 {
			wrong_even += 1
		}
		if a[2*i+1]%2 == 0 {
			wrong_odd += 1
		}
	}

	if len(a)%2 != 0 && a[len(a)-1]%2 != 0 {
		wrong_even += 1
	}
	return wrong_even, wrong_odd
}

// https://codeforces.com/problemset/problem/1367/B
// комментарий к решению
// Т.к. от нас не требуется выводить результирующий хороший массив,
// то достаточно считать кол-во элементов требующих перестановки
// на четных и нечетных местах
func taskSolution(in []test) string {
	ans := make([]int, len(in))
	for i, t := range in {
		wrong_even, wrong_odd := calc_wrongs(t.a)
		if wrong_even != wrong_odd {
			ans[i] = -1
		} else {
			ans[i] = wrong_even
		}
	}
	res := fmt.Sprint(ans)
	return strings.Replace(res[1:len(res)-1], " ", "\n", len(ans)-1)
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
