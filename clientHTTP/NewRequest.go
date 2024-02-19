package main

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"strings"
)


func main(){
	fmt.Println("Sending a Request")
	req, err := http.NewRequest("DELETE","http://www.google.com/robots.txt",nil)
	var client http.Client

	resp,err := client.Do(req)

	if err!=nil{
		fmt.Printf("Error")
	}

	resp.Body.Close()

	form:=url.Values{}
	form.Add("foo","bar")
	var client1 http.Client

	request, err:= http.NewRequest("PUT","https://www.google.com/robots.txt",strings.NewReader(form.Encode()))
	if err!=nil{
		fmt.Printf("Error sending the data")
	}
	response,err:=client1.Do(request)
	body,err:=io.ReadAll(response.Body)
	fmt.Println(string(body))
}