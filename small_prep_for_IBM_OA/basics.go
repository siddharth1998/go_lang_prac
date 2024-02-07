// Package name
package main


import "fmt"

func main(){
	fmt.Print("Hi there")
	// Implicit type casting
	var impVar= 10
	impVar2 := 2	

	// Explicit type casting 
	var expVar int 
	expVar = 10
	
	fmt.Print(impVar)
	fmt.Print(impVar2)
	fmt.Print(expVar)	
	

	fmt.Printf("\nThe values are the following  %d %d %d",impVar,impVar2,expVar)
	

	// Calling func Hello
	fmt.Print(Hello("Sid"))
}

func Hello (name string) string{
	return ("\nHello "+name)
}
