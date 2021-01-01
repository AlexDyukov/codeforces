package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/344/A
// комментарии к решению
// Островки считаются как непохожесть последнему проверенному магниту
// т.е. 101010: 10 == 10 == 10 => 1 остров
// 10011001: 10 != 01 != 10 != 01 => 4 острова
// для исключение повторного просчёта надо сдвигаться вправо на 2 символа
// вместо одного после проверки каждого магнита
func taskSolution(a []byte) (res int) {
	last := ""
	for i := 1; i < len(a); i += 1 {
		switch {
		case a[i-1] == '0' && a[i] == '1':
			if last != "01" {
				res += 1
				last = "01"
			}
			i += 1
		case a[i-1] == '1' && a[i] == '0':
			if last != "10" {
				res += 1
				last = "10"
			}
			i += 1
		}
	}
	return res
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

	fmt.Println(taskSolution(inRaw))
}
