package main

import (
	"bytes"
	"strings"
	"testing"
)

func TestDo(t *testing.T) {
	testCases := []struct {
		in   string
		want string
	}{
		{in: "3 4", want: "Even"},
		{in: "1 21", want: "Odd"},
	}
	for _, tC := range testCases {
		t.Run(tC.in, func(t *testing.T) {
			r := strings.NewReader(tC.in)
			var w bytes.Buffer
			do(r, &w)

			if w.String() != tC.want {
				t.Errorf("want: %s got: %s", tC.want, w.String())
			}
		})
	}
}
