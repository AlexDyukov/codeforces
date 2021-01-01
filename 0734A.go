package main

import (
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/734/A
// комментарии к решению
// В виду большого n - длины строки, то посимвольное считывание + проверка
// может не уложиться по времени и приходиться буферизовывать весь входной поток
// подсчёт партий выполняется счётчиками. Для улучшения читаемости их два, а не один
func taskSolution(a []byte) string {
	antonWins := 0
	danikWins := 0
	for _, b := range a {
		switch b {
		case 'A':
			antonWins += 1
		case 'D':
			danikWins += 1
		}
	}

	if antonWins > danikWins {
		return "Anton"
	} else if antonWins < danikWins {
		return "Danik"
	} else {
		return "Friendship"
	}
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

	fmt.Println(taskSolution(inRaw))
}
