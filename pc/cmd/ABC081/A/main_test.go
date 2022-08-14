package main

import (
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	r := strings.NewReader("101")
	n := do(r)
	if n != 2 {
		t.Fatal(n)
	}
}
