package main

import "fmt"

const friends = 2

type form struct {
	home int
	out  int
}

// https://codeforces.com/problemset/problem/268/A
// комментарий к решению
// решаем перебором проходя по каждой паре с проверкой,
// что домашняя форма итерируемого в данный момент совпадает
// с формой на выезде для любой другой команды
func taskSolution(a []form) (ans int) {
	for _, h := range a {
		for _, o := range a {
			if h != o && h.home == o.out {
				ans += 1
			}
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []form{}
	for i := 0; i < n; i += 1 {
		var h, o int
		if _, err := fmt.Scan(&h, &o); err != nil {
			panic(err)
		}
		a = append(a, form{home: h, out: o})
	}

	fmt.Println(taskSolution(a))
}
