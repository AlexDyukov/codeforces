package main

import "fmt"

// https://codeforces.com/problemset/problem/579/A
// комментарии к решению
// Т.к. бактерии удваиваются ежедневно, то нам достаточно
// контролировать только нечетные изменения. На примере:
// 5 {нужно получить}
// -> 4{было}+1{положили}
// -> 2{было}*2{1ночь}+1положили
// -> 1{было}*2{1ночь}*2{2ночь}+1положили
// -> (0{старт}+1{положили})*2{1ночь}*2{2ночь}+1положили
func taskSolution(n int) (ans int) {
	for n > 0 {
		if n%2 == 1 {
			ans += 1
		}
		n /= 2
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n))
}
