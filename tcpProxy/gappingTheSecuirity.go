package main
// Remote SHELL

import (
	"fmt"
	"io"
	"log"
	"net"
	"os"
	"os/exec"
)


type FooReader struct{}

type FooWriter struct{}

func (fooWriter * FooWriter) Write(b []byte)(int,error){
	fmt.Print("out > ")
	return os.Stderr.Write(b)
}



func execution(conn net.Conn){
	

	commandExecution:=exec.Command("/bin/bash","-i")
	readp,writep:=io.Pipe()
	commandExecution.Stdin=conn
	commandExecution.Stdout=writep
	go io.Copy(conn,readp)
	commandExecution.Run()
	// if err:=commandExecution.Run();err!=nil{
	// 	fmt.Println(err)
	// 	var val [] byte
	// 	val=[]byte("Closed")
	// 	conn.Write(val)
	// }
	fmt.Println("Connection closed for "+conn.RemoteAddr().String())

	defer conn.Close()
	
}

func main(){
	fmt.Println("Starting the server")

	listner,err:=net.Listen("tcp","localhost:5000")
	if err!=nil{
		fmt.Println("Server not able to start error")
		log.Fatalln(err)
	}
	

	for{
		conn,err:=listner.Accept()
		if err!=nil{
			fmt.Println("Server not able to accept new connection")
			log.Fatalln(err)		
		}
		// send conn
		fmt.Println("New Connection From "+conn.RemoteAddr().String())
		go execution(conn)
	}
}