package main

import (
	"io"
	"log"
	"net"
)

func echo(conn net.Conn) {
	// Connection handler
	defer conn.Close()

	// Buffer to handle the read operation
	b := make([]byte, 1024)
	for {
		size, err := conn.Read(b[0:])
		if err == io.EOF {
			log.Println("The client got disconnected")
			break
		}
		if err != nil {
			log.Println("Unexpected Error")
			break
		}
		log.Printf("Bytes %d : %s \n", size, string(b))

		log.Println("Writing your data ")
		if _, err := conn.Write(b[0:size]); err != nil {
			log.Println("Not able to write")
		}
	}
}

func main() {

	// Port Binding
	listner, err := net.Listen("tcp", ":2080")
	if err != nil {
		log.Fatalln("Unable to bind the port")
	}
	log.Println("Listening on port 2080")

	for {
		conn, err := listner.Accept()

		log.Println("Recived Connection")
		if err != nil {
			log.Fatalln("Not able to accept the connection")

		}

		// Send to connection handler :: Compute
		go echo(conn)

	}
}
