package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

type s struct{}

var exists s = s{}

const max = uint64(1000000000000)

func isPrime(primes []uint64, num uint64) bool {
	for i := 0; i < len(primes) && primes[i]*primes[i] <= num; i += 1 {
		if num%primes[i] == 0 {
			return false
		}
	}
	return true
}

func generatePrimes(m uint64) []uint64 {
	ans := []uint64{2}
	for i := uint64(3); i*i <= m; i += uint64(2) {
		if isPrime(ans, i) {
			ans = append(ans, i)
		}
	}
	return ans
}

func generateTPrimes(m uint64) map[uint64]s {
	ans := map[uint64]s{}
	primes := generatePrimes(m)
	for _, p := range primes {
		ans[p*p] = exists
	}
	return ans
}

// parse uint64 from byte slice
func parseUInt64(a []byte, start int) (ans uint64) {
	for i := start; i < len(a) && a[i] >= '0' && a[i] <= '9'; i += 1 {
		// skip overflow check
		ans = 10*ans + uint64(a[i]-'0')
	}
	return ans
}

// https://codeforces.com/problemset/problem/230/B
// комментарии к решению
// Задача сводиться к проверке является ли число квадратом простого числа или нет
// Сначала необходимо сгенерировать все простые числа до нужного порога,
// потом сгенерировать из них Т-простые числа,
// после чего можно итерироваться по входному потоку.
// Для хранения Т-простых хорошо подходит set (тип map[int]struct{} в golang)
// Плюсом реализуем свою parseInt, т.к. fmt.Scan не справляется с такими входными данными
func taskSolution(a []byte, n int) string {
	// 4 == "YES\n"
	ans := make([]byte, 0, 4*n)
	tPrimes := generateTPrimes(max)

	intStart := 0
	// check first and iterate from second to prevent trimming end space
	cur := parseUInt64(a, intStart)
	if _, ok := tPrimes[cur]; ok {
		ans = append(ans, []byte("YES")...)
	} else {
		ans = append(ans, []byte("NO")...)
	}
	for intStart < len(a) {
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
		cur = parseUInt64(a, intStart)
		if _, ok := tPrimes[cur]; ok {
			ans = append(ans, []byte("\nYES")...)
		} else {
			ans = append(ans, []byte("\nNO")...)
		}
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

	fmt.Println(taskSolution(bytes.TrimSpace(inRaw), n))
}
