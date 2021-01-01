package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type number struct {
	begin int
	end   int
}

func compareNumbers(a []byte, left number, right number) int {
	if (left.end - left.begin) > (right.end - right.begin) {
		return 1
	} else if (left.end - left.begin) < (right.end - right.begin) {
		return -1
	}
	diff := right.begin - left.begin
	for i := left.begin; i < left.end; i += 1 {
		if a[i] > a[i+diff] {
			return 1
		} else if a[i] < a[i+diff] {
			return -1
		}
	}
	return 0
}

// https://codeforces.com/problemset/problem/580/A
// комментарии к решению
// т.к. числа достаточно большие и их много, то fmt.Scan не справляется,
// поэтому необходимо сравнивать их посимвольно, либо использовать strconv.atoi
// Сравнение строк-чисел достаточно простая операция, которой учат в первых главах K&R,
// поэтому с небольшой модификацией получаем compareNumbers(),
// которую используем для сравнения чисел
// Алгоритм простой:
// 0) постоянно проверяем на < len(a) чтобы не словить out of index
// 1) запоминаем позицию последнего считанного числа
// 2) находим следующее
// 3) сравниваем
// 4) go to 0
func taskSolution(a []byte) int {
	currLen, maxLen := 1, 1
	right := number{begin: 0, end: 0}
	// find first number
	right.end = right.begin
	for right.end < len(a) && a[right.end] >= '0' && a[right.end] <= '9' {
		right.end += 1
	}
	// iterate numbers
	for right.end < len(a) {
		left := right
		// find next number
		right.begin = left.end + 1
		right.end = left.end + 1
		for right.end < len(a) && a[right.end] >= '0' && a[right.end] <= '9' {
			right.end += 1
		}
		// compare
		if compareNumbers(a, left, right) <= 0 {
			currLen += 1
			if maxLen < currLen {
				maxLen = currLen
			}
		} else {
			currLen = 1
		}
	}
	return maxLen
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

	fmt.Println(taskSolution(bytes.TrimSpace(inRaw)))
}
