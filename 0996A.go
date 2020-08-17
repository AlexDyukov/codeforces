package main

import "fmt"

// https://codeforces.com/problemset/problem/996/A
// комментарии к решению
// Т.к. номиналы купюр известны заранее и они полностью перекрывают друг друга
// т.е. любая старшая купюра легко может быть представлена определенным кол-вом
// более младшей купюры, то минимальное кол-во купюр для выдачи n денег
// легко поддаётся подсчету по алгоритму ниже
func taskSolution(n int) (ans int) {
	nominals := []int{100, 20, 10, 5, 1}
	for i := 0; i < len(nominals) && n > 0; i += 1 {
		ans, n = ans+n/nominals[i], n%nominals[i]
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
