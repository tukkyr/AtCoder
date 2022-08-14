package main

import (
	"fmt"
	"io"
	"os"
	"strings"
)

func do(r io.Reader) (a int) {
	var s string
	fmt.Fscan(r, &s)
	return strings.Count(s, "1")
}

func main() {
	a := do(os.Stdin)
	fmt.Print(a)
}
