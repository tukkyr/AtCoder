package main

import (
	"fmt"
	"io"
	"os"
)

func do(r io.Reader, w io.Writer) {
	var a, b, c int
	var s string

	fmt.Fscanf(r, "%d\n", &a)
	fmt.Fscanf(r, "%d %d\n", &b, &c)
	fmt.Fscanf(r, "%s", &s)

	fmt.Fprintf(w, "%d %s", a+b+c, s)
}

func main() {
	do(os.Stdin, os.Stdout)
}
