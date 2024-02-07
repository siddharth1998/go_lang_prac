// This program shows how worker pool works

package main

import (
	"fmt"
	"sync"
)

func worker(ports chan int, wg *sync.WaitGroup){
	// Print the port for now 
	for p:= range ports{
		fmt.Println(p)
		wg.Done()
	}
}

func main(){
	fmt.Print("/n Starting the program /n")
	var wg sync.WaitGroup
	ports :=make(chan int, 100)// creating a pipe for sending data to from worker
	

	// Create a pool of hundred workers
	for i:=0; i<cap(ports);i++{
		go worker(ports,&wg)
	}
	for i:=0; i<1024;i++{
		wg.Add(1)
		ports<-i
	}
	wg.Wait()
	close(ports)
}