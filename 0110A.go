package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/110/A
// комментарии к решению
// Задача больше не про числа, а больше про символы
// "Почти счастливое число - это где кол-во счастливых символов = счастливое число"
// Т.к. n (1 ≤ n ≤ 10**18), то входной поток ограничен 18 символами
// т.е. почти счастливое число - число, где кол-во счастливых чисел равно 4 или 7
func taskSolution(in []byte) string {
	counter := 0
	for _, char := range in {
		if char == '4' || char == '7' {
			counter += 1
		}
	}

	if counter == 4 || counter == 7 {
		return "YES"
	}
	return "NO"
}

func main() {
	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(inRaw))
}
