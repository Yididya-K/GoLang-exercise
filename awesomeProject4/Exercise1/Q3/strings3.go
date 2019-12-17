package main

import (
	"fmt"
	"strings"
)

func main() {
	A:="asSASA ddd dsjkdsjs dk"

	new := strings.Replace(A, "ASA", "abc", 1)
	fmt.Println(new)
}
