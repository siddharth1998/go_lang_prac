package main

import (
	"fmt"
	"net/http"
	"net/url"
	"strings"
)

func main() {
	r1, err := http.Get("http://www.google.com/robots.txt")
	defer r1.Body.Close()

	r2, err := http.Head("http://www.google.com/robots.txt")
	r2.Body.Close()
	defer r2.Body.Close()

	form := url.Values{}
	form.Add("foo", "bar")
	r3, err := http.Post("https://www.google.com/robots.txt", "application/x-www-urlencoded", strings.NewReader(form.Encode()))
	r3.Body.Close()
	defer r3.Body.Close()

	form1 := url.Values{}
	form1.Add("foo", "bar")
	r4, err := http.PostForm("https://www.google.com/robots.txt", form1)
	if err!=nil{
		fmt.Println(err)
	}
	defer r4.Body.Close()



	if err != nil {
		fmt.Println(err)
	}
}