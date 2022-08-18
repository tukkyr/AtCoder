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

func want(t *testing.T, i int, w interface{}, g interface{}) {
	t.Helper()
	t.Errorf("case:%d, want: %v, got: %v", i, w, g)
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

func td(i int, do func(in question) []answer, isLog bool) (d []answer, a answer, du time.Duration, _ error) {
	pwd, _ := os.Getwd()
	qr, err := os.Open(filepath.Join(pwd, "cases", fmt.Sprintf("%d.in", i)))
	if err != nil {
		return d, a, du, isFinish(err)
	}
	defer qr.Close()
	ar, err := os.Open(filepath.Join(pwd, "cases", fmt.Sprintf("%d.out", i)))
	if err != nil {
		return d, a, du, isFinish(err)
	}
	defer ar.Close()
	in := input(qr)
	s := time.Now()
	d = do(in)
	a = output(ar)
	du = time.Since(s)
	if isLog {
		fmt.Printf("case:%d %dms\n", i, du.Milliseconds())
	}
	return d, a, du, nil
}

var TS = 2 * time.Second

func TestSample(t *testing.T) {
	for i := 1; i < 10; i++ {
		gs, w, du, err := td(i, do, true)
		if err == done {
			break
		}
		if err != nil {
			t.Fatal(err)
		}
		var ok bool
		for _, g := range gs {
			if g.eq(&w) {
				ok = true
				break
			}
		}
		if !ok {
			want(t, i, w, gs)
		}
		if du > TS {
			t.Errorf("case:%d timeout: %dms", i, (du - TS).Milliseconds())
		}
	}
}

func BenchmarkDo(b *testing.B) {
	for i := 0; i < b.N; i++ {
		for n := 0; n < 10; n++ {
			td(n, do, false)
		}
	}
}
