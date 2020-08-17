package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const maxLen = 100001

func max(a, b int64) int64 {
	if a > b {
		return a
	}
	return b
}

// parse positive int from byte slice
func parseInt(a []byte, start int) (ans int) {
	for i := start; i < len(a) && a[i] >= '0' && a[i] <= '9'; i += 1 {
		// skip overflow check
		ans = 10*ans + int(a[i]-'0')
	}
	return ans
}

// parse byte slice (string) into structurized array/map; faster then fmt.Scan
func parseInput(a []byte, n int) []int {
	intStart := 0
	ans := make([]int, 0, n)
	for intStart < len(a) {
		ans = append(ans, parseInt(a, intStart))
		// find next
		for intStart < len(a) && a[intStart] >= '0' && a[intStart] <= '9' {
			intStart += 1
		}
		for intStart < len(a) && !(a[intStart] >= '0' && a[intStart] <= '9') {
			intStart += 1
		}
	}
	return ans
}

func calcCounts(a []int) [maxLen]int {
	ans := [maxLen]int{}
	for _, v := range a {
		ans[v] += 1
	}
	return ans
}

// https://codeforces.com/problemset/problem/455/A
// комментарии к решению
// Порядок чисел в последовательности ничего не значит. Плюс лимиты на возможные значения не высокие,
// поэтому нам достаточно подсчитать кол-во одинаковых чисел и проитерироваться через весь промежуток
// возможных значений на предмет эффективности удаления i (итерируемого) из последовательности
func taskSolution(in []byte, n int) int64 {
	a := parseInput(in, n)
	counts := calcCounts(a)

	f2, f1 := int64(0), int64(0)
	for i := 0; i < maxLen; i += 1 {
		f2, f1 = f1, max(f1, f2+int64(i)*int64(counts[i]))
	}
	return f1
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

	fmt.Println(taskSolution(bytes.TrimSpace(inRaw), n))
}
