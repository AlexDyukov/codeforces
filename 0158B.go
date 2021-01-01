package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

func min(a, b int) int {
	if a < b {
		return a
	}
	return b
}

func extractGroups(line []byte) []int {
	ans := []int{0, 0, 0, 0}
	for _, e := range line {
		switch e {
		case '1':
			ans[0] += 1
		case '2':
			ans[1] += 1
		case '3':
			ans[2] += 1
		case '4':
			ans[3] += 1
		}
	}
	return ans
}

// https://codeforces.com/problemset/problem/158/B
// комментарии к решению
// Т.к. (1 ≤ si ≤ 4) жестко задан одним символом, то профитнее обработать входной поток как []byte
// Т.к. размер групп не произвольное число,
// то можно описать все возможные комбинации без перебора:
// 1) Группы из 4-х человек полностью влезают в такси
// 2) Группы из 3-х человек могут добрать к себе одиночек или поехать втроём
// 3) Группы из 2-х человек могут скооперироваться и при нечетном их кол-ве добрать одиночек
// 4) Оставшиеся одиночки просто заказывают минимальное необходимое кол-во такси
func taskSolution(in []byte) (ans int) {
	groups := []int{0, 0, 0, 0}
	groups = extractGroups(in)

	ans += groups[4-1]

	ans += groups[3-1]
	groups[1-1] -= min(groups[3-1], groups[1-1])

	quotient := groups[2-1] / 2
	remainder := groups[2-1] % 2
	ans += quotient
	ans += min(remainder, 1)
	groups[1-1] -= min(remainder*2, groups[1-1])

	ans += groups[1-1] / 4
	ans += min(groups[1-1]%4, 1)

	return ans
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
