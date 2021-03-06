package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// parse integer from byte slice
func parseInt(a []byte, pos int) (ans, newpos int) {
	for pos < len(a) && a[pos] >= '0' && a[pos] <= '9' {
		pos, ans = pos+1, 10*ans+int(a[pos]-'0')
	}
	return ans, pos + 1
}

// https://codeforces.com/problemset/problem/1374/A
// комментарии к решению
// fmt.Scan для нас удобнее, но на большом входном потоке он падает по TLE,
// поэтому читаем []byte и парсим самостоятельно.
// Для буферизованного вывода используем fmt.Sprint , с заменой пробелов
// на \n. Кол-во замен и их размер (1 байт == ' ' на 1 байт == '\n') известно,
// поэтому будет сгенерирован оптимальный код без большого оверхеда
func taskSolution(in []byte, t int) string {
	ans := make([]int, t)
	var x, y, n int
	for i, pos := 0, 0; i < t; i += 1 {
		for in[pos] < '0' || in[pos] > '9' {
			pos += 1
		}
		x, pos = parseInt(in, pos)
		y, pos = parseInt(in, pos)
		n, pos = parseInt(in, pos)
		ans[i] = n - n%x + y
		if ans[i] > n {
			ans[i] -= x
		}
	}
	res := fmt.Sprint(ans)
	return strings.Replace(res[1:len(res)-1], " ", "\n", len(ans)-1)
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
