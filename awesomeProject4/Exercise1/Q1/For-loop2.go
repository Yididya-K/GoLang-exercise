package main

import "fmt"

func main() {
	i:=1
	start:
		if(i<=10){
			fmt.Print(i ," ")
			i++
			goto start
		}
		println()
}
