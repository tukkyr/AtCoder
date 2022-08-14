package main

import (
	"fmt"
	"io"
	"os"
)

func do(r io.Reader, w io.Writer) {
	var a, b int
	fmt.Fscanf(r, "%d %d", &a, &b)

	if a*b%2 == 0 {
		fmt.Fprint(w, "Even")
	} else {
		fmt.Fprintf(w, "Odd")
	}
}

func main() {
	do(os.Stdin, os.Stdout)
}
