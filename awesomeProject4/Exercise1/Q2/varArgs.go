package main

import "fmt"

func main() {

	for i:=1;i<=100;i++{
		if(i%5==0 && i%3==0){
			fmt.Print("fizzbuzz")
		}else if  (i%3==0){
			fmt.Print("fizz")
		}else if(i%5==0){
			fmt.Print("buzz")
		}else{
			fmt.Print(i)
		}
		fmt.Print(" , ")
	}
}
