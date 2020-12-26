package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/login",loginFunc1)

	http.ListenAndServe("127.0.0.1:8899",nil)
}
func loginFunc1(writer http.ResponseWriter, request *http.Request)  {
	request.ParseForm()

	for k,v:=range request.Form{
	fmt.Println(k,v)
	}
	fmt.Println(writer,"欢迎你："+request.Form.Get("user"))
}