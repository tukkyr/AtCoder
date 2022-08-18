package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func input(r io.Reader) (ns []int) {
	var n int
	fmt.Fscan(r, &n)
	ns = make([]int, n)
	for i := range ns {
		fmt.Fscan(r, &ns[i])
	}
	sort.Sort(sort.Reverse(sort.IntSlice(ns)))
	return ns
}

func output(r io.Reader) (w int) {
	fmt.Fscan(r, &w)
	return
}

func do(r io.Reader) int {
	ns := input(r)
	var a, b int
	for i, n := range ns {
		if i&1 == 0 {
			a += n
		} else {
			b += n
		}
	}
	return a - b
}

func main() {
	fmt.Println(do(os.Stdin))
}
