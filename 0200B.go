package main

import "fmt"

// https://codeforces.com/problemset/problem/200/B
// комментарии к решению
// В задаче требуется вывести среднее среди предоставленного массива с допустимой погрешностью
// По идее необходимо также посчитать сколько знаков после запятой следует выводить
// для удовлетворения условия по абсолютной и относительной погрешности, но для такого
// уровня сложности задачи должно быть лень => смотрим кол-во знаков после запятой в примерах
func taskSolution(a []int) string {
	sum := 0
	for _, ai := range a {
		sum += ai
	}
	avg := float64(sum) / float64(len(a))
	return fmt.Sprintf("%.12f", avg)
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []int{}
	for i := 0; i < n; i += 1 {
		var ai int
		if _, err := fmt.Scan(&ai); err != nil {
			panic(err)
		}
		a = append(a, ai)
	}

	fmt.Println(taskSolution(a))
}
