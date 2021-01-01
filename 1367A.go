package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/1367/A
// комментарии к решению
// Как видно из примера и алгоритма из строки:
// bccddaaf
// получается:
// bcdaf
// т.е. все серединные (не первый и не последний) парные символы
// нужно превратить в одиночные. В примере выше это достигается
// путём выбора символов с позиций 1 - 3 - 5 - 7 - 8
func taskSolution(a []byte, t int) string {
	ans := []byte{}
	for i, pos := 0, 0; i < t; i += 1 {
		// find next
		// skip len check, because t value control number count
		for !(a[pos] >= 'a' && a[pos] <= 'z') {
			pos += 1
		}
		// skip len check, because t value control number count
		for a[pos] >= 'a' && a[pos] <= 'z' {
			ans = append(ans, a[pos])
			pos += 2
		}
		ans = append(ans, a[pos-1], '\n')
	}
	// truncate last new-line
	return string(ans[:len(ans)-1])
}

func main() {
	var t int
	if _, err := fmt.Scan(&t); err != nil {
		panic(err)
	}

	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(inRaw, t))
}
