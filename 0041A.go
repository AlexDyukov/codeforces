package main

import "fmt"

// https://codeforces.com/problemset/problem/41/A
// комментарии к решению
// задача аналогична классической реализации strcmp за исключением,
// что по второму слову мы итерируемся с конца
// Лимиты на входные данные низкие, поэтому переменные string и fmt.Scan достаточно
func taskSolution(w1, w2 string) string {
	if len(w1) != len(w2) {
		return "NO"
	}

	length := len(w1)

	for i := 0; i < length; i += 1 {
		if w1[i] != w2[length-i-1] {
			return "NO"
		}
	}

	return "YES"
}

func main() {
	var word1, word2 string
	if _, err := fmt.Scan(&word1, &word2); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(word1, word2))
}
