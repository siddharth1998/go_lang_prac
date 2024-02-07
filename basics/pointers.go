package main 

import "fmt"

func main(){

	var count="charlizard"
	address :=&count
	fmt.Println(*address,address)
	*address = "10000"
	fmt.Println(*address,address)
	
}