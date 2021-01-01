package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

const maxLen = 6

// https://codeforces.com/problemset/problem/96/A
// комментарии к решению
// В виду низких требований к входным данным == длина строки <= 100
// и она не изменяется, то достаточно использовать string и fmt.Scan
// Задача просто требует проверить наличие подпоследовательности
// больше определенной длины
func taskSolution(line []byte) string {
	last := byte(' ')
	currLen := 0
	for i := 0; i < len(line); i += 1 {
		if currLen > maxLen {
			return "YES"
		} else if line[i] == last {
			currLen += 1
		} else {
			currLen = 1
			last = line[i]
		}
	}
	if currLen > maxLen {
		return "YES"
	}
	return "NO"
}

func main() {
	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}
	fmt.Println(taskSolution(bytes.TrimSpace(inRaw)))
}
