package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/432/A
// комментарии к решению
// Отбрасываем участников, которые не могут поучаствовать k раз.
// Из остальных формируем команды
func taskSolution(a []byte, k int) int {
	members := map[byte]int{0: 0, 1: 0, 2: 0, 3: 0, 4: 0, 5: 0}
	for _, c := range a {
		if c >= '0' && c <= '5' {
			members[c-'0'] += 1
		}
	}

	membersCount := 0
	for i := 0; i < len(members)-k; i += 1 {
		membersCount += members[byte(i)]
	}
	return membersCount / 3
}

func main() {
	var n, k int
	if _, err := fmt.Scan(&n, &k); err != nil {
		panic(err)
	}

	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(inRaw, k))
}
