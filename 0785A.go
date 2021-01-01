package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/785/A
// комментарии к решению
// Можно заметить, что все многогранники имеют различную заглавную букву,
// поэтому сравнивать построчно избыточно. Кол-во граней для каждой фигуры
// задано в условии, поэтому буферизовываем весь входной поток и проверяем посимвольно
func taskSolution(a []byte) (ans int) {
	for _, char := range a {
		switch char {
		case 'T': // Tetrahedron
			ans += 4
		case 'C': // Cube
			ans += 6
		case 'O': // Octahedron
			ans += 8
		case 'D': // Dodecahedron
			ans += 12
		case 'I': // Icosahedron
			ans += 20
		}
	}
	return ans
}

func main() {
	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(inRaw))
}
