package main

import (
	"fmt"
	"unicode"
)

// https://codeforces.com/problemset/problem/520/A
// комментарии к решению
// Т.к. строка не большая, то можно использовать fmt.Scan и считать string
// Для приведения букв к одному регистру используем unicode.ToLower
// Дл подсчета кол-ва уникальных букв используем map[rune]struct{}
func taskSolution(line string) string {
	type s struct{}
	exists := s{}
	uniq_symbols := map[rune]s{}
	for _, char := range line {
		uniq_symbols[unicode.ToLower(char)] = exists
	}
	if len(uniq_symbols) == int('z')-int('a')+1 {
		return "YES"
	}
	return "NO"
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
