package Q6

import "fmt"

func order(n1 ,n2 int)(int,string,int){
	if(n1>n2){
		return n2," , ",n1
	}else{
		return n1," , ",n2
	}
}
func main() {

	fmt.Println(order(4,7))
}


