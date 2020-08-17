package main

import "fmt"

func nonUniq(n int) bool {
	a1, a2, a3, a4 := n/1000, (n/100)%10, (n/10)%10, n%10
	return a1 == a2 || a1 == a3 || a1 == a4 || a2 == a3 || a2 == a4 || a3 == a4
}

// https://codeforces.com/problemset/problem/271/A
// комментарии к решению
// Т.к. требуется число, цифры которого уникальны, то используем int
// цифру на i месте для четырехзначного числа можно получить формулой (n/10**(5-i)%10
func taskSolution(n int) int {
	n += 1
	for nonUniq(n) {
		n += 1
	}
	return n
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
