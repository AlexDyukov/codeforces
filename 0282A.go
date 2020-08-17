package main

import "fmt"

// https://codeforces.com/problemset/problem/282/A
// комментарии к решению
// центральный символ каждой строки однозначно идентифицирует операцию
// поэтому для инкремента/декремента ориентируемся на него
func taskSolution(a []string) (ans int) {
	for _, str := range a {
		// check central symbol of string
		if str[1] == '+' {
			ans += 1
		} else {
			ans -= 1
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []string{}
	for i := 0; i < n; i += 1 {
		var ai string
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a))
}
