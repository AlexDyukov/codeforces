package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/1352/A
// комментарии к решению
// Т.к. требуется оперировать цифрами, а не числами, то задействуем посимвольную обработку
// входного потока. Т.к. числа из входного потока <= 10000, то и их количество в выходной строке
// будет не больше 4, за что отвечает 1 символ, поэтому нам не требуется предварительно считать
// на сколько чисел нам необходимо разбить текущее, а можно просто инкрементировать символ '0'
// на позиции кол-ва чисел. Для примера по итерациям разбив 5009 будет выглядеть так:
// 0: 0\n
// 1: 1\n5000
// 2: 2\n5000 9
func taskSolution(a [][]byte) string {
	ans := []byte{'0', '\n'}
	num := a[0]
	nCountPos := 0
	for j := 0; j < len(num); j += 1 {
		if num[j] > '0' && num[j] <= '9' {
			ans[nCountPos] = ans[nCountPos] + byte(1)
			ans = append(ans, num[j])
			for k := j + 1; k < len(num) && num[k] >= '0' && num[k] <= '9'; k += 1 {
				ans = append(ans, byte('0'))
			}
			ans = append(ans, byte(' '))
		}
	}
	for i := 1; i < len(a); i += 1 {
		ans[len(ans)-1] = '\n'
		num = a[i]
		nCountPos = len(ans)
		ans = append(ans, []byte("0\n")...)
		for j := 0; j < len(num); j += 1 {
			if num[j] > '0' && num[j] <= '9' {
				ans[nCountPos] = ans[nCountPos] + byte(1)
				ans = append(ans, num[j])
				for k := j + 1; k < len(num) && num[k] >= '0' && num[k] <= '9'; k += 1 {
					ans = append(ans, byte('0'))
				}
				ans = append(ans, byte(' '))
			}
		}
	}
	return string(ans[:len(ans)-1])
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	inRawArr := bytes.Split(bytes.TrimSpace(inRaw), []byte("\n"))

	fmt.Println(taskSolution(inRawArr))
}
