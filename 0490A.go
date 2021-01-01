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

// https://codeforces.com/problemset/problem/490/A
// комментарии к решению
// Входной поток состоит только из цифр 1, 2 и 3, поэтому можно
// считать его как []byte и разобрать посимвольно.
// Т.к. для ответа требуется указывать порядковые номера, то
// распределяем учеников по группам с соответствующими способностями
func taskSolution(in []byte, n int) (ans string) {
	type s []int
	children := map[byte]s{
		'1': s{},
		'2': s{},
		'3': s{},
	}
	// no overflow потому что мы знаем объем входных данных
	for i, pos := 1, 0; i <= n; i, pos = i+1, pos+1 {
		// find next
		for in[pos] < '1' || in[pos] > '3' {
			pos += 1
		}
		children[in[pos]] = append(children[in[pos]], i)
	}
	// calculate
	count := min(len(children['1']), len(children['2']))
	count = min(count, len(children['3']))
	ans = fmt.Sprint(count)
	for i := 0; i < count; i += 1 {
		ans = fmt.Sprintf("%s\n%d %d %d", ans, children['1'][i], children['2'][i], children['3'][i])
	}
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

	fmt.Println(taskSolution(inRaw, n))
}
