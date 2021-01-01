package main

import (
	"fmt"
)

func getAlphas(line string) map[rune]int {
	ans := map[rune]int{}
	for _, r := range line {
		ans[r] += 1
	}
	return ans
}

// https://codeforces.com/problemset/problem/141/A
// комментарии к решению
// В виду низких требований к входным данным == длина строк <= 100
// их можно считать через fmt.Scan как string.
// Для подсчета задействованных букв удобно использовать map[rune]int.
// Если из букв кучи отнимать необходимое кол-во букв для каждого имени, то
// отрицательные значения покажут, что буквы не хватает,
// а положительное значение - что эти буквы лишние.
// Т.к. с нас просят проверить нет ли лишних или недостающих букв, то проверяем != 0
func taskSolution(guest, owner, left string) string {
	guestAlphaMap := getAlphas(guest)
	ownerAlphaMap := getAlphas(owner)
	leftAlphaMap := getAlphas(left)
	for i, c := range guestAlphaMap {
		leftAlphaMap[i] -= c
	}

	for i, c := range ownerAlphaMap {
		leftAlphaMap[i] -= c
	}

	for _, c := range leftAlphaMap {
		if c != 0 {
			return "NO"
		}
	}
	return "YES"
}

func main() {
	var guest, owner, alphasLeft string
	if _, err := fmt.Scan(&guest, &owner, &alphasLeft); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(guest, owner, alphasLeft))
}
