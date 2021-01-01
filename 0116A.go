package main

import "fmt"

type stopDiff struct {
	pout int
	pin  int
}

// https://codeforces.com/problemset/problem/116/A
// комментарии к решению
// В виду низких требований к входным данным == n(1≤n≤1000)
// и наличия строго правила "сначала выходят, потом заходят"
// задача решается в лоб:
// После каждой остановки смотрим заполненность трамвая
// и считаем его необходимый размер согласно требованиям задачи
func taskSolution(in []stopDiff) (ans int) {
	size := 0
	for i := 0; i < len(in); i += 1 {
		size += in[i].pin - in[i].pout
		if ans < size {
			ans = size
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	in := []stopDiff{}
	for i := 0; i < n; i += 1 {
		var pout, pin int
		if _, err := fmt.Scan(&pout, &pin); err != nil {
			panic(err)
		}
		in = append(in, stopDiff{pout, pin})
	}

	fmt.Println(taskSolution(in))
}
