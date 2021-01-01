package main

import (
	"fmt"
	"io/ioutil"
	"os"
	"strings"
)

// parse positive int from byte slice
func parseInt(a []byte, start int) (ans int) {
	for i := start; i < len(a) && a[i] >= '0' && a[i] <= '9'; i += 1 {
		// skip overflow check
		ans = 10*ans + int(a[i]-'0')
	}
	return ans
}

func findX(mersennes []int, n int) int {
	for i := 0; i < len(mersennes); i += 1 {
		if n%mersennes[i] == 0 {
			return n / mersennes[i]
		}
	}
	panic("unexpected behaviour")
}

// https://codeforces.com/problemset/problem/1343/A
// комментарии к решению
// Можно заметить, что уравнение приводиться к:
// x + 2 * x + 4 * x +..+ 2^(k−1) * x = n =>
// (2^k - 1) * x = n
// где коэффициент перед х не что иное как числа Мерсенна, которых до 10^9 не много.
// Выписываем их и итерируемся по ним для каждой строки, пока не найдём целый х:
// х = n / (2^k - 1)
func taskSolution(a []byte, t int) string {
	// Mersenne numbers
	mersennes := []int{
		3, 7, 15, 31, 63, 127, 255, 511, 1023, 2047, 4095, 8191,
		16383, 32767, 65535, 131071, 262143, 524287, 1048575,
		2097151, 4194303, 8388607, 16777215, 33554431, 67108863,
		134217727, 268435455, 536870911,
	}
	intStart := 0
	ans := []int{}
	for i := 0; i < t; i += 1 {
		// find next
		// skip len check, because t value control number count
		for !(a[intStart] >= '0' && a[intStart] <= '9') {
			intStart += 1
		}
		cur := parseInt(a, intStart)
		ans = append(ans, findX(mersennes, cur))
		// skip parsed
		for a[intStart] >= '0' && a[intStart] <= '9' {
			intStart += 1
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
