package main

import "fmt"

const friends = 2

type room struct {
	size     int
	capacity int
}

// https://codeforces.com/problemset/problem/467/A
// комментарий к решению
// для удобочитаемости заводим структуру
// с полями текущего размера и вместимости
// а после в цикле считаем, перебирая комнаты
// на предмет вместимости нашей пары друзей
func taskSolution(a []room) (ans int) {
	for _, r := range a {
		if friends+r.size <= r.capacity {
			ans += 1
		}
	}
	return ans
}

func main() {
	var n int
	if _, err := fmt.Scan(&n); err != nil {
		panic(err)
	}

	a := []room{}
	for i := 0; i < n; i += 1 {
		var s, c int
		if _, err := fmt.Scan(&s, &c); err != nil {
			panic(err)
		}
		a = append(a, room{size: s, capacity: c})
	}

	fmt.Println(taskSolution(a))
}
