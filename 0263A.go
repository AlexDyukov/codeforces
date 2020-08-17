package main

import "fmt"

const dimension = 5
const mean = 2

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

// https://codeforces.com/problemset/problem/263/A
// комментарии к решению
// В задаче требуется посчитать расстояние от "1" до середины матрицы,
// что несложно сделать принимая во внимание известный размер матрицы
// и индексацию элементов с 0 - середина будет в [2,2]:
// 0 1 2 3 4
// 1
// 2   *
// 3
// 4
func taskSolution(a [][]int) int {
	for i, line := range a {
		for j, elem := range line {
			if elem == 1 {
				return abs(mean-i) + abs(mean-j)
			}
		}
	}
	panic("unexpected behavior")
}

func main() {
	a := [][]int{}
	for i := 0; i < dimension; i += 1 {
		line := []int{}
		for j := 0; j < dimension; j += 1 {
			var n int
			if _, err := fmt.Scan(&n); err != nil {
				panic(err)
			}
			line = append(line, n)
		}
		a = append(a, line)
	}

	fmt.Println(taskSolution(a))
}
