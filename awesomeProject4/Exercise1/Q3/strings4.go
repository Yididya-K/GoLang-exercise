package main

import "fmt"

func main() {
	var a []string
	A:="foobar"

	for i:=len(A)-1;i>=0;i--{
		a=append(a,string(A[i]))
	}
	fmt.Println(a)
}
