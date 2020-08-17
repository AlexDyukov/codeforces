package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

type line struct {
	start int
	end   int
}

// https://codeforces.com/problemset/problem/4/C
// комментарии к решению
// При попытке считывания string через fmt.Scan мы падаем по времени, поэтому
// парсим входные данные самостоятельно, приводя строки к string только в крайнем случае
func taskSolution(in []byte) string {
	names := map[string]int{}
	ans := []byte{}
	next := line{start: 0, end: 0}
	// iterate lines
	for next.end < len(in) {
		next.start = next.end
		// skip spaces
		for next.start < len(in) && !(in[next.start] >= 'a' && in[next.start] <= 'z') {
			next.start += 1
		}
		// find end of word
		next.end = next.start
		for next.end < len(in) && in[next.end] >= 'a' && in[next.end] <= 'z' {
			next.end += 1
		}
		// break at EOF
		if next.end == len(in) {
			break
		}
		// check logins
		loginStr := string(in[next.start:next.end])
		if _, ok := names[loginStr]; ok {
			ans = append(ans, []byte(fmt.Sprintf("%s%d", loginStr, names[loginStr]))...)
			names[loginStr] += 1
		} else {
			ans = append(ans, []byte("OK")...)
			names[loginStr] = 1
		}
		ans = append(ans, byte('\n'))
	}
	return string(ans)
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
