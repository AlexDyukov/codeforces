package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/281/A
// комментарий к решению
// стандартная библиотека strings содержит соответствующий метод капитализации слова
// Производительность string меньше []byte, но для текущей задачи достаточно
func taskSolution(line string) string {
	return strings.Title(line)
}

func main() {
	var line string
	if _, err := fmt.Scan(&line); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(line))
}
