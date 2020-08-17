package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/133/A
// комментарии к решению
// Задача требует проверить, выведет ли что-то поданная на вход программа.
// Т.о. мы должно просто удостовериться, что во входящем потоке символов
// есть один из трёх символов-кодов, которые хоть что-то выводят: 'H'/'Q'/'9'
// На вход подаются строго символы ASCII из диапазона [33;126], поэтому
// без негативных последствий можно считать []byte и пройтись по нему
func taskSolution(line []byte) string {
	for _, r := range line {
		switch r {
		case 'H', 'Q', '9':
			return "YES"
		}
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
