package main

import (
	"fmt"
	"io"
	"os"
	"sort"
)

func nCount(n int, b int) (c int) {
	for b%n == 0 {
		b /= n
		c++
	}
	return c
}

func count(b int) (s int) {
	return nCount(2, b)
}

func min(ns ...int) int {
	sort.Sort(sort.IntSlice(ns))
	return ns[0]
}

func max(ns ...int) int {
	sort.Sort(sort.IntSlice(ns))
	n := len(ns)
	return ns[n-1]
}

func input(r io.Reader) (ns []int) {
	var n int
	fmt.Fscan(r, &n)
	ns = make([]int, n)
	for i := range ns {
		fmt.Fscan(r, &ns[i])
	}
	return ns
}

func output(r io.Reader) (w int) {
	fmt.Fscan(r, &w)
	return
}

func do(r io.Reader) int {
	ns := input(r)
	var cs []int
	for _, n := range ns {
		c := count(n)
		if c == 0 {
			return 0
		}
		cs = append(cs, count(n))
	}
	return min(cs...)
}

func main() {
	fmt.Println(do(os.Stdin))
}
