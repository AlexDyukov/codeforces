package main

import (
	"fmt"
	"strings"
	"unicode"
)

func countLowerUpper(s string) (lower, upper int) {
	for _, c := range s {
		if unicode.IsUpper(c) {
			upper += 1
		} else {
			lower += 1
		}
	}
	return lower, upper
}

// https://codeforces.com/problemset/problem/59/A
// комментарии к решению
// Требуется сменить регистр регисторового меньшинства букв на противоположный
// Наиболее производительным вариантом будет итерация по []byte,
// но в виду низких требований к входным данным реализуем через strings,
// т.к. решение короче
func taskSolution(line string) string {
	if lower, upper := countLowerUpper(line); lower < upper {
		return strings.ToUpper(line)
	}
	return strings.ToLower(line)
}

func main() {
	var line string
	if _, err := fmt.Scan(&line); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(line))
}
