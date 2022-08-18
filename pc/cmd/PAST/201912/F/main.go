package main

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"reflect"
	"sort"
	"strings"
	"unicode"
)

type question struct {
	s string
}

type answer struct {
	s string
}

func (a answer) String() string {
	return fmt.Sprintf("%s", a.s)
}

func (a *answer) eq(o *answer) bool {
	return reflect.DeepEqual(a, o)
}

func input(r io.Reader) (q question) {
	fmt.Fscan(r, &q.s)
	return
}

func output(r io.Reader) (a answer) {
	fmt.Fscan(r, &a.s)
	return
}

func do(in question) (as []answer) {
	var words []string
	for i := 0; i < len(in.s); {
		j := i + 1
		for {
			if j < len(in.s) && unicode.IsLower(rune(in.s[j])) {
				j++
				continue
			}
			break
		}
		word := in.s[i : j+1]
		words = append(words, word)
		i = j + 1
	}
	sort.Slice(words, func(i, j int) bool { return strings.ToLower(words[i]) < strings.ToLower(words[j]) })
	as = append(as, answer{s: strings.Join(words, "")})
	return as
}

func main() {
	r := bufio.NewReader(os.Stdin)
	as := do(input(r))
	fmt.Printf("%v\n", as[0])
}
