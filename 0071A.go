package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const maxLen = 10

// https://codeforces.com/problemset/problem/71/A
// комментарии к решению
// В виду низких требований к входным данным == длина строк <= 100
// задача решается в лоб через строки:
// Читаем строку -> если длина больше 10, то
// в ответ кидаем
// 1) первый символ
// 2) кол-во символом между первой и последней буквой
// 3) последний символ
// 4) \n
// Не забываем обрезать последний лишний перевод строки
func taskSolution(a [][]byte) (ans string) {
	for i := 0; i < len(a); i += 1 {
		line := bytes.TrimSpace(a[i])
		if len(line) > maxLen {
			ans = fmt.Sprintf("%s%c%d%c\n", ans, line[0], len(line)-2, line[len(line)-1])
		} else {
			ans = fmt.Sprintf("%s%s\n", ans, line)
		}
	}
	return ans[:len(ans)-1]
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
