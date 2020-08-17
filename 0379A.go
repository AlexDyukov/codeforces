package main

import "fmt"

// https://codeforces.com/problemset/problem/379/A
// комментарии к решению
// Итеративным подходом зажигаем все свежие свечи,
// потом делаем из необходимого кол-ва использованных свежие свечи
// и так до тех пор, пока не закончатся нормальные свечи и
// из использованных нельзя будет сделать ниодной нормальной
func taskSolution(a, b int) (ans int) {
	r, n := a, 0
	for r > 0 || n >= b {
		ans, r, n = ans+r, r/b+n/b, n%b+r%b
	}
	return ans
}

func main() {
	var a, b int
	if _, err := fmt.Scan(&a, &b); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(a, b))
}
