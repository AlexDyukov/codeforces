package main

import "fmt"

// https://codeforces.com/problemset/problem/148/A
// комментарии к решению
// По идее можно сообразить формулу по сумме всех непересекающихся частных,
// но если для двух уникальных делителей она читаема:
// x/a + x/b - x/(a*b)
// то для трёх и более уже адекватнее использовать перебор
// при таких маленьких ограничениях на входные параметры
func taskSolution(k, l, m, n, d int) int {
	ans := 0
	for num := 1; num <= d; num += 1 {
		if num%k == 0 || num%l == 0 || num%m == 0 || num%n == 0 {
			ans += 1
		}
	}
	return ans
}

func main() {
	var k, l, m, n, d int
	if _, err := fmt.Scan(&k, &l, &m, &n, &d); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(k, l, m, n, d))
}
