package main

import (
	"fmt"
	"net"
)

func main() {

	fmt.Printf("Starting the connection")
	_,err :=net.Dial("tcp","scanme.nmap.org:80")
	if (err==nil){
		fmt.Print("\nConnection Successfull")
	}else { 
		fmt.Print(err);
	}
}
