package main

import (
	"fmt"
	"unicode"
)

func isCapslocked(a []rune) bool {
	if len(a) > 1 {
		return unicode.IsUpper(a[1])
	}
	return true
}

func toLowerTitled(a []rune) []rune {
	if a[0] == unicode.ToUpper(a[0]) {
		a[0] = unicode.ToLower(a[0])
	} else {
		a[0] = unicode.ToUpper(a[0])
	}
	for i := 1; i < len(a); i += 1 {
		a[i] = unicode.ToLower(a[i])
	}
	return a
}

// https://codeforces.com/problemset/problem/131/A
// комментарии к решению
// Входные параметры в виде небольшой строки позволяют задействовать пакет unicode.
// Пакет unicode работает с rune, а string это массив rune, поэтому
// уместно считывать сразу string.
// Входные параметры не содержат слов с лесенкой, поэтому
// если любой не первый символ IsUpper, то у нас зажат CapsLock и надо исправлять.
// При исправлении меняем регистр первого символа на противоположный
// и приводим к нижнему регистру все остальные символы
func taskSolution(line string) string {
	a := []rune(line)
	if isCapslocked(a) {
		return string(toLowerTitled(a))
	}
	return line
}

func main() {
	var line string
	if _, err := fmt.Scan(&line); err != nil {
		panic(err)
	}

	fmt.Println(taskSolution(line))
}
