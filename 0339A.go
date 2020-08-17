package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/339/A
// комментарии к решению
// Т.к. цифры заранее известны, то самым простым решением будет
// 1) запомнить кол-во повторений нужного числа
// 2) выписать все числа вставив между ними знак '+'
func taskSolution(line []byte) string {
	count := []int{0, 0, 0}
	for _, b := range line {
		if b >= '1' && b <= '3' {
			count[b-'1'] += 1
		}
	}
	ans := []byte{}
	for i := 0; i < len(count); i += 1 {
		for j := 0; j < count[i]; j += 1 {
			ans = append(ans, byte(i)+'1', '+')
		}
	}

	return string(ans[:len(ans)-1])
}

func main() {
	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(taskSolution(inRaw))
}
