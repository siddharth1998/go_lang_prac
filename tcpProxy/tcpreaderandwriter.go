package main

import(
	"fmt"
	"log"
	"os"
)

// struct to read from io.Reader from Stdin
type FooReader struct{}

// struct to read from io.Writer to write to stdout
type FooWriter struct{}

func(fooReader * FooReader) Read(b []byte)(int,error){
	fmt.Print("in > ")
	return os.Stdin.Read(b)
}

func (fooWriter * FooWriter ) Write(b [] byte)(int,error){
	fmt.Print("out >")
	return os.Stdout.Write(b)
}

func main(){
	var(
		reader FooReader
		writer FooWriter
	)

	// A buffer channel to take inputs bytes because from input we will get the bytes
	input := make([] byte,4096)
	s,err := reader.Read(input)
	if err!=nil{
		log.Fatalln("Unable to read data")
	}
	fmt.Printf("Read %d bytes fropm stdin \n ",s)

	// Use writer to write output
	s, err= writer.Write(input)
	if err!= nil{
		log.Fatalln("Unable to write to output")
	} 
	fmt.Printf("Wrote %d bytes to stdout \n",s)
}