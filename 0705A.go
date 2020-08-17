package main

import "fmt"

// https://codeforces.com/problemset/problem/705/A
// комментарий к решению
// Задача на конкатенацию строк, причём решение в лоб через string
// ест 10Мб, поэтому приходится манипулировать []byte
func taskSolution(n int) string {
	ans := []byte("I hate")
	a := [][]byte{
		[]byte(" that I love"),
		[]byte(" that I hate"),
	}
	for i := 2; i <= n; i += 1 {
		ans = append(ans, a[i%2]...)
	}
	ans = append(ans, []byte(" it")...)
	return string(ans)
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
