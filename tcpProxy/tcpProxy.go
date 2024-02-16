package main

import (
	"fmt"
	"io"
	"log"
	"net"
)

func handle(src net.Conn) {

	dst, err := net.Dial("tcp", "localhost:10000")

	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Problem in connecting with dst server")

	}
	
	defer dst.Close()

	go func() {

		if _, err := io.Copy(src, dst); err != nil {
			log.Fatalln(err)
	fmt.Printf("b")

			log.Fatalln("Source not able to send data")
		}
	}()


	if _, err := io.Copy(dst, src); err != nil {
		log.Fatalln(err)
		fmt.Printf("c")
		log.Fatalln("dst not able to send the data")
	}
}

func main() {
	fmt.Println("Starting the TCP server")
	// accetp the connection
	// send the data from the source to dst
	listner, err := net.Listen("tcp", ":80")
	if err != nil {
		log.Fatalln(err)
		log.Fatalln("Problem starting the Proxy server (listner module)")

	}
	for {
		conn, err := listner.Accept()
		if err != nil {
			log.Fatalln("Problem Accepting Connections")
		}else{
			fmt.Printf("New client connected")
		}
		// TODO send conn
		go handle(conn)
	}
}
