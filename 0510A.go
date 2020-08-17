package main

import "fmt"

// https://codeforces.com/problemset/problem/510/A
// комментарии к решению
// Все используемые в выводе символы входят в ASCII таблицу
// и занимают 1 байт в кодировке UTF-8, поэтому можно выделить
// []byte необходимого размера и заполнить его, согласно
// предопределенной маске в задаче
func taskSolution(n int, m int) string {
	//if n < 3 || m < 3 {
	//	panic("invalid parameters")
	//}
	ans := make([]byte, 0, n*(m+1)+1)

	// init first line to prevent trimming \n at the end
	for j := 0; j < m; j += 1 {
		ans = append(ans, byte('#'))
	}

	for i := 1; i < n; i += 1 {
		ans = append(ans, byte('\n'))
		switch i % 4 {
		case 0, 2:
			for j := 0; j < m; j += 1 {
				ans = append(ans, byte('#'))
			}
		case 1:
			for j := 1; j < m; j += 1 {
				ans = append(ans, byte('.'))
			}
			ans = append(ans, byte('#'))
		case 3:
			ans = append(ans, byte('#'))
			for j := 1; j < m; j += 1 {
				ans = append(ans, byte('.'))
			}
		}
	}
	return string(ans)
}

func main() {
	var n, m int
	if _, err := fmt.Scan(&n, &m); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(n, m))
}
