package main

import (
	"fmt"
	"net"
	"sync"
)

func main() {
	var wg sync.WaitGroup
	for i := 0; i < 1024; i++ {
		wg.Add(1) // Each time you create a go routine
		go func(j int) {
			defer wg.Done() // Every time you do some owrk you call this to -1
			address := fmt.Sprintf("scanme.nmap.org:%d", j)
			conn, err := net.Dial("tcp", address)
			if err != nil {
				fmt.Printf("%d port is closed", j)
				return
			}
			conn.Close()
			fmt.Printf("%d port is open", j)
		}(i)
	}
	wg.Wait()
}
