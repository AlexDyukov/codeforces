package main

import "fmt"

// https://codeforces.com/problemset/problem/489/C
// комментарии к решению
// Максимальное число всегда будет выглядеть как:
// 99..99Х00..00 , где Х - цифра
// а минимальное число будет выглядеть как:
// 100...0Х99..99 или Х99..99 , где Х - цифра
func taskSolution(m int, s int) string {
	if (s == 0 && m != 1) || (s > 9*m) {
		return "-1 -1"
	}

	ans := make([]byte, 2*m+1)
	last := 0
	for i := 0; i < m; i += 1 {
		switch {
		case s == 0:
			ans[m+i+1] = '0'
			ans[m-i-1] = '0'
		case s > 9:
			ans[m+i+1] = '9'
			ans[m-i-1] = '9'
			s -= 9
		default:
			last = m - i - 1
			ans[m+i+1] = byte(s) + '0'
			ans[m-i-1] = byte(s) + '0'
			s = 0
		}
	}
	if last > 0 {
		ans[0] = '1'
		ans[last] -= byte(1)
	}
	// put space between two numbers
	ans[m] = ' '
	return string(ans)
}

func main() {
	var m, s int
	if _, err := fmt.Scan(&m, &s); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(m, s))
}
