package main


import "fmt"

func main(){
	fmt.Println("this is slices and map")
	var s=make([]string,0)
	s=append(s,"hi there")
	s=append(s,"pikachu")
	fmt.Println(s)
	// above is the slice
	var maper=make(map[string]string)
	maper["pika"]="chu"
	fmt.Println(maper)
}
