package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestMain(t *testing.T) {
	input := `1
2 3
test`
	r := strings.NewReader(input)
	var b bytes.Buffer
	do(r, &b)
	var ans = "6 test"

	if b.String() != ans {
		t.Errorf("want %s got %s", ans, b.String())
	}
}
