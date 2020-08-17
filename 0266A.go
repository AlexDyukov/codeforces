package main

import "fmt"

// https://codeforces.com/problemset/problem/266/A
// комментарии к решению
// в задаче требуется подсчитать кол-во элементов,
// похожих на элемент на предшествующей позиции
// Т.к. требования к входным данным не высокие,
// то можно воспользоваться string вместо []byte
func taskSolution(line string) (ans int) {
	for i := 1; i < len(line); i += 1 {
		if line[i] == line[i-1] {
			ans += 1
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	var line string
	if _, err := fmt.Scan(&line); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(line))
}
