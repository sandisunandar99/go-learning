package main

import (
	"fmt"
	"net/url"
)

func main() {

	var urlString = "https://jsonplaceholder.typicode.com/todos/2?userid=1"
	var u, e = url.Parse(urlString)

	if e != nil {
		panic(e)
	}

	fmt.Printf("url : %s \n", urlString)
	fmt.Printf("scheme : %s \n", u.Scheme)
	fmt.Printf("host : %s \n", u.Host)
	fmt.Printf("path : %s \n", u.Path)

	var userid = u.Query().Get("userid")
	fmt.Printf("userid : %s \n", userid)

}
