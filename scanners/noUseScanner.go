package main

import(
	"fmt"
	"net"
)


func main(){
	fmt.Print("this is a no use print")
	
	for i :=0; i <1024 ; i++{
		go func(j int){
			address :=fmt.Sprintf("scanme.nmap.org:%d \n",j)
			conn,err :=net.Dial("tcp",address)
			if err!=nil{
				//port is closed or filtered
				return
			}
			conn.Close()
			fmt.Printf("open port %d\n",j)

		}(i)
		
	}
}