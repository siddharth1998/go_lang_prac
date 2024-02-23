package main

import (
	"fmt"
	"os"

	"shodan/myPkg"
)

func main(){
	api_key:= os.Getenv("SHODAN_API_KEY")
	client:=myPkg.New(api_key)
	api_info,err:=client.APIInfo()
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Printf("%s \n",api_info)
	fmt.Printf(
		"Query Credits: %d\nScan Credits: %d\n\n",
		api_info.QueryCredits,
		api_info.ScanCredits)
	
	hostSearch,err :=client.HostSearch(os.Args[1])
	if err!=nil{
		fmt.Println("there is some problem in host search ")
		fmt.Println(err)
	}
	for _,host:= range hostSearch.Matches{
		fmt.Printf("\n %18s:: %8d \n",host.IPString,host.Port)
	}
}