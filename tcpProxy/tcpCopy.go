package main

import(
	"fmt"
	"log"
	"io"
	"os"
)

type FooReader struct{}

type FooWriter struct{}

func (fooWriter * FooWriter) Write(b []byte)(int,error){
	fmt.Print("out > ")
	return os.Stderr.Write(b)
}

func (fooReader * FooReader) Read(b []byte)(int,error){
	fmt.Print("in >")
	return os.Stderr.Read(b)
}

func main(){
	var ( 
		reader FooReader
		writer FooWriter
	)

	if _,err:=io.Copy(&writer,&reader); err!=nil{
		log.Fatalln("Can not copy from reader to writer")
	}
}