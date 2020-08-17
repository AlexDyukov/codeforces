package main

import (
	"fmt"
	"strings"
)

func ipow2(exp int) int {
	res := 1
	for exp > 0 {
		res, exp = res*2, exp-1
	}
	return res
}

// https://codeforces.com/problemset/problem/1348/A
// комментарии к решению
// n=2: 2 4                    => (4)         - (2)            => 2
// n=4: 2 4 8 16               => (16+2)      - (4+8)          => 6
// n=6: 2 4 8 16 32 64         => (64+2+4)    - (32+16+8)      => 14
// n=8: 2 4 8 16 32 64 128 256 => (256+2+4+8) - (128+64+32+16) => 30
// => 2^(n/2+1)-2
func taskSolution(a []int) string {
	ans := make([]int, len(a))
	for i := 0; i < len(a); i += 1 {
		ans[i] = ipow2(a[i]/2+1) - 2
	}
	res := fmt.Sprint(ans)
	return strings.Replace(res[1:len(res)-1], " ", "\n", len(ans)-1)
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
