package main

import "fmt"

const sub = "hello"

// https://codeforces.com/problemset/problem/58/A
// комментарии к решению
// Задача требует проверить соответствие входной строки на соответствие regex "*h*e*l*l*o*"
// Чуть модифицированный strcmp справляется с поставленной задачей за 1 проход
func taskSolution(line string) string {
	var i, j int
	for i, j = 0, 0; i < len(sub) && j < len(line); j += 1 {
		if sub[i] == line[j] {
			i += 1
		}
	}
	if i == len(sub) && j <= len(line) {
		return "YES"
	}
	return "NO"
}

func main() {
	var line string
	if _, err := fmt.Scan(&line); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(line))
}
