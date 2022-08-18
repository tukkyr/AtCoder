package main

import (
	"fmt"
	"io"
	"os"
	"reflect"
)

type question struct {
	n int
	d []int
}

type answer struct {
	x int
}

func defaultVal() (a answer) {
	a.x = 0
	return
}

func (a answer) String() string {
	return fmt.Sprintf("%d", a.x)
}

func (a *answer) eq(o *answer) bool {
	return reflect.DeepEqual(a, o)
}

func input(r io.Reader) (q question) {
	fmt.Fscan(r, &q.n)
	q.d = make([]int, q.n)
	for i := range q.d {
		fmt.Fscan(r, &q.d[i])
	}
	return
}

func output(r io.Reader) (a answer) {
	fmt.Fscan(r, &a.x)
	return
}

func do(in question) (as []answer) {
	/*
		var set = make(map[int]struct{})
		for _, d := range in.d {
			set[d] = struct{}{}
		}
		return []answer{{x: len(set)}}
	*/
	var exist = make([]int, 101)
	for _, d := range in.d {
		exist[d] = 1
	}
	var r int
	for _, n := range exist {
		r += n
	}
	return []answer{{x: r}}
}

func main() {
	as := do(input(os.Stdin))
	fmt.Printf("%v\n", as[0])
}
