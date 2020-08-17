package main

import (
	"fmt"
	"strings"
)

// https://codeforces.com/problemset/problem/1335/A
// комментарии к решению
// Т.к. требуется арифметика с числами, то используем fmt.Scan для считывания массива.
// Для вывода воспользуемся fmt.Sprint переведя массив ответов в строку.
// Т.к. fmt.Sprint выводит данные через пробел, то задействуем strings.Replace
// для вывода ответов по одному на строку.
// Компилятор видит, что заменяется 1 однобайтовый символ unicode на 1 однобайтовый символ unicode,
// плюс мы указываем кол-во замен, поэтому будет сгенерирован оптимальный по памяти код,
// как если бы мы производили свою арифметику с []byte
func taskSolution(a []int) string {
	ans := a[:]
	for i := 0; i < len(a); i += 1 {
		ans[i] = (a[i] - 1) / 2
	}
	res := fmt.Sprint(ans)
	return strings.Replace(res[1:len(res)-1], " ", "\n", len(ans)-1)
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []int{}
	for i := 0; i < n; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a))
}
