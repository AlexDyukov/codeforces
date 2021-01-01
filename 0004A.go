package main

import "fmt"

// https://codeforces.com/problemset/problem/4/A
// комментарии к решению
// любое четное число больше 3 можно поделить на 2 четные части
func taskSolution(n int) string {
	if n > 3 && n%2 == 0 {
		return "YES"
	}
	return "NO"
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
