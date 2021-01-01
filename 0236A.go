package main

import "fmt"

// https://codeforces.com/problemset/problem/236/A
// комментарии к решению
// Делаем подсчёт уникальных символов через set (map[T]struct{})
// и выводим ответ в зависимовсти от четности размера уникального множества
func taskSolution(line string) string {
	possibleAnswers := []string{"CHAT WITH HER!", "IGNORE HIM!"}
	type s struct{}
	exists := s{}
	runes := map[rune]s{}

	for _, c := range line {
		runes[c] = exists
	}
	return possibleAnswers[len(runes)%2]
}

func main() {
	var line string
	if _, err := fmt.Scan(&line); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(line))
}
