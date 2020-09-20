package main

import (
	"fmt"
	"strings"
)

type gifts struct {
	n int
	a []int
	b []int
}

func max(a, b int) int {
	if a > b {
		return a
	}
	return b
}

// https://codeforces.com/problemset/problem/1399/B
// комментарий к решению
// Т.к. требуется арифметика с числами, то используем fmt.Scan для считывания массива.
// Для вывода воспользуемся fmt.Sprint переведя массив ответов в строку.
// Т.к. fmt.Sprint выводит данные через пробел, то задействуем strings.Replace
// для вывода ответов по одному на строку.
// Компилятор видит, что заменяется 1 однобайтовый символ unicode на 1 однобайтовый символ unicode,
// плюс мы указываем кол-во замен, поэтому будет сгенерирован оптимальный по памяти код,
// как если бы мы производили свою арифметику с []byte
func taskSolution(in []gifts) string {
	ans := make([]int64, len(in))
	for i := 0; i < len(in); i += 1 {
		n, a, b := in[i].n, in[i].a, in[i].b
		min_a, min_b := a[0], b[0]
		for j := 1; j < n; j += 1 {
			if a[j] < min_a {
				min_a = a[j]
			}
			if b[j] < min_b {
				min_b = b[j]
			}
		}
		for j := 0; j < n; j += 1 {
			ans[i] += int64(max(a[j]-min_a, b[j]-min_b))
		}
	}
	res := fmt.Sprint(ans)
	return strings.Replace(res[1:len(res)-1], " ", "\n", len(ans)-1)
}

func main() {
	var t int
	if _, err := fmt.Scan(&t); err != nil {
		panic(err)
	}

	in := make([]gifts, 0, t)
	var n, elem int
	for i := 0; i < t; i += 1 {
		if _, err := fmt.Scan(&n); err != nil {
			panic(err)
		}
		a := make([]int, 0, n)
		for j := 0; j < n; j += 1 {
			if _, err := fmt.Scan(&elem); err != nil {
				panic(err)
			}
			a = append(a, elem)
		}
		b := make([]int, 0, n)
		for j := 0; j < n; j += 1 {
			if _, err := fmt.Scan(&elem); err != nil {
				panic(err)
			}
			b = append(b, elem)
		}
		in = append(in, gifts{n: n, a: a, b: b})
	}

	fmt.Println(taskSolution(in))
}
