package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// parse uint from byte slice
func parseUInt(a []byte, start int) (ans uint64) {
	for i := start; i < len(a) && a[i] >= '0' && a[i] <= '9'; i += 1 {
		// skip overflow check
		ans = 10*ans + uint64(a[i]-'0')
	}
	return ans
}

// https://codeforces.com/problemset/problem/339/B
// комментарии к решению
// Т.к. движение одностороннее, то каждый раз, когда нам нужен дом с меньшим ID,
// приходится делать неполный круг. В остальных случаях достаточно посчитать
// сколько времени у нас займёт путь до следующего дома
// При использовании fmt.Scan со считыванием int мы не укладываемся по времени:
// $ shuf -i 1-100000 -n 100000 | tr '\n' ' ' > input
// $ sed  -i '1i 100000 100000' input
// $ cat input | time -p go run 0339B.go
// 4999974741
// real 4.54
// user 1.07
// sys 3.90
// Поэтому задействуем ioutil.ReadAll() и самописную функцию парсинга []byte в uint,
// которые получаются быстрее в 4-5 раз:
// real 0.90
// user 0.51
// sys 0.81
func taskSolution(a []byte, n uint64) uint64 {
	intStart := 0
	cur := parseUInt(a, intStart)
	// no overflow because input numbers start from 1
	ans := cur - 1
	for intStart < len(a) {
		prev := cur
		// find next
		for intStart < len(a) && a[intStart] >= '0' && a[intStart] <= '9' {
			intStart += 1
		}
		for intStart < len(a) && !(a[intStart] >= '0' && a[intStart] <= '9') {
			intStart += 1
		}
		if intStart >= len(a) {
			break
		}
		// found
		cur = parseUInt(a, intStart)
		if prev > cur {
			ans += cur - prev + n
		} else {
			ans += cur - prev
		}
	}
	return ans
}

func main() {
	var n, m uint64
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic(err)
	}

	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(bytes.TrimSpace(inRaw), n))
}
