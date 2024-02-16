package main

import(
	"fmt"
	"net"
	"sort"
)


func worker(ports,result chan int){
	// distrubted worker implementation 
	// when got error send the values as 0 in the pipe or channel 
	// 
	for p:= range ports{
		// fmt.Printf("%d scanning the following port\n",p)
		address :=fmt.Sprintf("chargearth.com:%d",p);
		conn, err := net.Dial("tcp",address);
		if err != nil{
			fmt.Printf("%d is closed \n",p)
			result <- 0
			continue
		}
		conn.Close()// Close the connection before going ahead 
		fmt.Printf("%d Following port is open\n",p)

		result <- p
	}
}
func main(){
	ports := make(chan int,100)
	results := make(chan int)
	var openports [] int
	for i:=0; i < cap(ports); i++{

		// pool the workers 
		go worker(ports,results)
	}
	go func(){
		for i:=1;i <1024; i++{
			// assing the value of the port to the channel
			ports <-i
		}
	}()
	for i :=0;i < 1024;i ++{
		port :=<- results

		if port!=0{
			openports =append(openports,port)
		}
	}
	close(ports)
	close(results)
	sort.Ints(openports);
	for _, port:=range openports{

		fmt.Println("%d open\n",port)
	}
}