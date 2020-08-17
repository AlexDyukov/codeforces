package main

import "fmt"

// https://codeforces.com/problemset/problem/61/A
// комментарии к решению
// реализация XOR на строках, состоящих из '0' и '1'
// Можно ускорить решение переписав его на []byte,
// но в виду низких требований к входным данным - не целесообразно
func taskSolution(w1, w2 string) string {
	// по условию длина совпадает
	//if len(w1) != len(w2) {
	//    panic("diff len")
	//}
	res := []rune{}
	for i := 0; i < len(w1); i += 1 {
		if w1[i] == w2[i] {
			res = append(res, '0')
		} else {
			res = append(res, '1')
		}
	}

	return string(res)
}

func main() {
	var word1, word2 string
	if _, err := fmt.Scan(&word1, &word2); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(word1, word2))
}
