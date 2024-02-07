package main

import "fmt"

func main(){
	CalculateCost(100)	
}
	
 // CalculateCost works out the cost of producing the given number of cars.
 func CalculateCost(carsCount int) uint {

	fmt.Print((carsCount-carsCount%10)/10)
	fmt.Print("\n")        
data:= uint((carsCount-carsCount%10)/10*95000+carsCount*10000)
     fmt.Print(data)
    return data
 }

