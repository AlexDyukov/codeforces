package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/443/A
// комментарии к решению
// Требуется подсчитать кол-во неуникальных латинских букв,,
// поэтому уместно воспользоваться множествами set (map[T]struct{} в go)
func taskSolution(in []byte) int {
	type s struct{}
	exists := s{}
	chars := map[byte]s{}

	for _, c := range in {
		if c >= 'a' && c <= 'z' {
			chars[c] = exists
		}
	}

	return len(chars)
}

func main() {
	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(inRaw))
}
