package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/208/A
// комментарии к решению
// Для восстановления строки в изначальное состояние необходимо:
// удалить все "WUB", оставив вместо подряд идущих "WUB" между словами максимум один пробел
// Условие задачи гарантирует, что
// "Входные данные состоят из единственной непустой строки, состоящей только из заглавных латинских букв"
// Поэтому можно считать []byte и пройтись по нему без приведения к string/[]rune
func taskSolution(line []byte) string {
	res := []byte{}
	i := 0
	for i < len(line) && !(line[i] >= 'A' && line[i] <= 'Z') {
		i += 1
	}
	for i < len(line)-2 {
		if line[i] == 'W' && line[i+1] == 'U' && line[i+2] == 'B' {
			if len(res) > 0 && res[len(res)-1] != ' ' {
				res = append(res, ' ')
			}
			i += 3
		} else {
			res = append(res, line[i])
			i += 1
		}
	}
	for i < len(line) && line[i] >= 'A' && line[i] <= 'Z' {
		res = append(res, line[i])
		i += 1
	}
	return string(res)
}

func main() {
	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(inRaw))
}
