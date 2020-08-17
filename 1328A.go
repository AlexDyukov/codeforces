package main

import (
	"fmt"
	"strings"
)

type numpair struct {
	a int
	b int
}

// https://codeforces.com/problemset/problem/1328/A
// комментарий к решению
// Т.к. требуется арифметика с числами, то используем fmt.Scan для считывания массива.
// Для вывода воспользуемся fmt.Sprint переведя массив ответов в строку.
// Т.к. fmt.Sprint выводит данные через пробел, то задействуем strings.Replace
// для вывода ответов по одному на строку.
// Компилятор видит, что заменяется 1 однобайтовый символ unicode на 1 однобайтовый символ unicode,
// плюс мы указываем кол-во замен, поэтому будет сгенерирован оптимальный по памяти код,
// как если бы мы производили свою арифметику с []byte
func taskSolution(a []numpair) string {
	ans := []int{}
	for _, p := range a {
		ans = append(ans, (p.b-(p.a%p.b))%p.b)
	}
	res := fmt.Sprint(ans)
	return strings.Replace(res[1:len(res)-1], " ", "\n", len(ans)-1)
}

func main() {
	var t int
	if _, err := fmt.Scan(&t); err != nil {
		panic(err)
	}

	in := []numpair{}
	for i := 0; i < t; i += 1 {
		var a, b int
		if _, err := fmt.Scan(&a, &b); err != nil {
			panic(err)
		}
		in = append(in, numpair{a: a, b: b})
	}

	fmt.Println(taskSolution(in))
}
