package main

import(
	"fmt"
	"net"
)

func main(){
	for i := 80;i<1024;i++{
		address := fmt.Sprintf("chargearth.com:%d",i)
		fmt.Printf(" The number :: %d\n",i)
		conn,err := net.Dial("tcp",address)
		if(err!=nil){
			// It's been closed
			fmt.Print("Inside this ")
			fmt.Print("%s",err)
			continue
		}else{
			fmt.Printf("\n The %d Port is OPEN !!!!!",i)
			conn.Close()
		}

	
	}
}