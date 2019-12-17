package main

import "fmt"

func main() {
	var A[]float64
	A=append(A,2.34,6.74,4.56,5.43,23.98)
	var sum float64 =0
	for i:=0;i<len(A);i++{
		sum=sum+A[i]
	}
	fmt.Println(sum/float64(len(A)))
}
