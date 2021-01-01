package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// parse positive int from byte slice
func parseInt(a []byte, start int) (ans int) {
	for i := start; i < len(a) && a[i] >= '0' && a[i] <= '9'; i += 1 {
		// skip overflow check
		ans = 10*ans + int(a[i]-'0')
	}
	return ans
}

// left evens:
// (2 + n    )*(n/2)/2
// right odds without last:
// (1 + n - 3)*(n/2-1)/2
// last odd = left_sum - right_sum = 3*n/2-1
func generateNumbers(n int) []byte {
	ans := []byte{}
	for i := 2; i <= n; i += 2 {
		ans = append(ans, []byte(fmt.Sprintf("%d ", i))...)
	}
	for i := 1; i <= n-3; i += 2 {
		ans = append(ans, []byte(fmt.Sprintf("%d ", i))...)
	}
	ans = append(ans, []byte(fmt.Sprintln(3*n/2-1))...)
	return ans
}

// https://codeforces.com/problemset/problem/1343/B
// комментарии к решению
// Во первых для решения подходят только n кратные 4, т.к. нужна четная сумма нечетных
// Во вторых в самом простом решении при анализе примеров прослеживаются арифметические
// подпоследовательности в левой/правой частях решения для каждого n
// В третьих в виду ограничений по времени в 1 секунду лучше использовать ioutil.ReadAll
// вместо fmt.Scan
func taskSolution(a []byte, t int) string {
	intStart := 0
	ans := []byte{}
	for i := 0; i < t; i += 1 {
		// find next
		// skip len check, because t value control number count
		for !(a[intStart] >= '0' && a[intStart] <= '9') {
			intStart += 1
		}
		cur := parseInt(a, intStart)
		if cur%4 != 0 {
			ans = append(ans, []byte("NO\n")...)
		} else {
			ans = append(ans, []byte("YES\n")...)
			ans = append(ans, generateNumbers(cur)...)
		}
		// skip parsed
		for a[intStart] >= '0' && a[intStart] <= '9' {
			intStart += 1
		}
	}
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
