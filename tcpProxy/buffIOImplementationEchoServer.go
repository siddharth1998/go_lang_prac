package main

import(
	"log"
	"net"
	"bufio"
)

func Echo(conn net.Conn){
	// defer conn.Close()
	reader := bufio.NewReader(conn)// Take a Reader as parameter
	writer :=bufio.NewWriter(conn)
	for {
		s, err := reader.ReadString('\n')
		if err!=nil{
			log.Fatalln("Some error while reading")
		}
	
		
		
		if _,err :=writer.WriteString(s);err!=nil{
			log.Fatalln("Some error while writing")
		}
		writer.Flush()
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
		go Echo(conn)

		// Send to connection handler :: Compute
		

	}
}