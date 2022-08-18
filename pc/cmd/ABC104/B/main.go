package main

import (
	"fmt"
	"io"
	"os"
	"strings"
	"unicode"
)

func input(r io.Reader) (s string) {
	fmt.Fscan(r, &s)
	return
}

func isA(s string) bool {
	return string(s[0]) == "A"
}

func oneC(s string) bool {
	return strings.Count(s[2:len(s)-1], "C") == 1
}

func countU(s string) (c int) {
	for i := range s {
		if unicode.IsUpper(rune(s[i])) {
			c++
		}
	}
	return
}

func isTwice(s string) bool {
	return countU(s) == 2
}

func do(r io.Reader) string {
	s := input(r)
	if isA(s) && oneC(s) && isTwice(s) {
		return "AC"
	} else {
		return "WA"
	}
}

func main() {
	fmt.Println(do(os.Stdin))
}
