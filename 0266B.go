package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/266/B
// комментарии к решению
// в виду низких требований входных данных
// решаем задачу в лоб
// при увеличении t,n на несколько порядков имеет смысл итерироваться
// с первого B, а не с 1, как указано в решении и до последней G,
// т.к. B будут постепенно смещаться в конец, а G в начало
func taskSolution(in []byte, t int) string {
	boy := byte('B')
	for ; t > 0; t -= 1 {
		for i := 1; i < len(in); i += 1 {
			if in[i-1] == boy && in[i] != boy {
				in[i-1], in[i] = in[i], in[i-1]
				i += 1
			}
		}
	}
	return string(in)
}

func main() {
	var n, t int
	if _, err := fmt.Scan(&n, &t); err != nil {
		panic(err)
	}

	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(bytes.TrimSpace(inRaw), t))
}
