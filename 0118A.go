package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
	"os"
)

// https://codeforces.com/problemset/problem/118/A
// комментарии к решению
// В виду низких требований к входным данным == n(1≤n≤100)
// задачу можно решить через конкатенацию строк,
// но в go строки неизменяемы, поэтому это дорогое удовольствие даже для n == 100
// Т.к. нам необходимо учитывать только согласные, то лучшим выбором будет
// slice и set из гласных, т.к. их меньше, чем согласных.
// Т.е. копируем всё, что не гласные, вставляя перед ними точку
func taskSolution(in []byte) []byte {
	type s struct{}
	exists := s{}
	vowels := map[byte]s{
		'a': exists,
		'o': exists,
		'y': exists,
		'e': exists,
		'u': exists,
		'i': exists,
	}

	lowercased := bytes.ToLower(bytes.TrimSpace(in))
	ans := []byte{}

	for i := 0; i < len(lowercased); i += 1 {
		r := lowercased[i]
		if _, ok := vowels[r]; !ok {
			ans = append(ans, '.', r)
		}
	}
	return ans
}

func main() {
	inRaw, err := ioutil.ReadAll(os.Stdin)
	if err != nil {
		panic(err)
	}

	fmt.Println(string(taskSolution(inRaw)))
}
