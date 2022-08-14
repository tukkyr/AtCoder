package main

import (
	"errors"
	"fmt"
	"io/fs"
	"os"
	"path/filepath"
	"testing"
	"time"
)

func want(t *testing.T, w, g interface{}) {
	t.Helper()
	t.Errorf("want: %v, got: %v", w, g)
}

func TestCount(t *testing.T) {
	testCases := []struct {
		b int
		c int
	}{
		{b: 8, c: 3},
		{b: 7, c: 0},
		{b: 6, c: 1},
	}
	for _, tC := range testCases {
		t.Run(fmt.Sprint(tC.b), func(t *testing.T) {
			a := count(tC.b)
			if a != tC.c {
				want(t, tC.c, a)
			}
		})
	}
}

func TestMinMax(t *testing.T) {
	testCases := []struct {
		desc string
		ns   []int
		s    int
	}{
		{desc: "max", ns: []int{3, 4, 1, 6, 7}, s: 7},
		{desc: "min", ns: []int{3, 4, 1, 6, 7}, s: 1},
	}
	for _, tC := range testCases {
		t.Run(tC.desc, func(t *testing.T) {
			var a int
			switch tC.desc {
			case "max":
				a = max(tC.ns...)
			case "min":
				a = min(tC.ns...)
			}
			if a != tC.s {
				want(t, tC.s, a)
			}
		})
	}
}

var done = errors.New("done")

func isFinish(err error) error {
	var pe *fs.PathError
	if errors.As(err, &pe) {
		// log.Printf("err:%v op:%s path:%s", pe.Err, pe.Op, pe.Path)
		return done
	} else {
		return err
	}
}

func td(i int) (int, int, error) {
	pwd, _ := os.Getwd()
	q, err := os.Open(filepath.Join(pwd, "cases", fmt.Sprintf("%d.in", i)))
	if err != nil {
		return 0, 0, isFinish(err)
	}
	defer q.Close()
	a, err := os.Open(filepath.Join(pwd, "cases", fmt.Sprintf("%d.out", i)))
	if err != nil {
		return 0, 0, isFinish(err)
	}
	defer a.Close()
	w := output(a)
	s := time.Now()
	g := do(q)
	fmt.Printf("case:%d %dμs\n", i, time.Since(s).Microseconds())
	return g, w, nil
}

func TestSample(t *testing.T) {
	for i := 1; i < 10; i++ {
		g, w, err := td(i)
		if err == done {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		if g != w {
			want(t, w, g)
		}
	}
}
