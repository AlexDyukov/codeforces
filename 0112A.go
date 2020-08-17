package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/112/A
// комментарии к решению
// В виду низких требований к входным данным == длина строк <= 100
// задача решается в лоб:
// считываем две строки -> приводим к одному регистру -> сравниваем строки
// При больших строках производительность сильно не просядет:
// вариант с проходом по []byte экономнее по памяти, но
// занимает в 4 раза больше кода и отличается лишь на пару memcpy
// поэтому пока нас не жмут по памяти данное решение в приоритете
func taskSolution(a, b string) (ans int) {
	a, b = strings.ToLower(a), strings.ToLower(b)
	return strings.Compare(a, b)
}

func main() {
	var first, second string
	if _, err := fmt.Scan(&first, &second); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(first, second))
}
