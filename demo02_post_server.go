package main

import (
	"fmt"
	"net/http"
)

func main() {
   http.HandleFunc("/login",nil)

   http.ListenAndServe("127.0.0.1:9999",nil)

}
func loginFunc(writer http.ResponseWriter, request *http.Request) {
	name:=request.FormValue("username")
	pwd:=request.FormValue("password")
	fmt.Println("用户名：",name)
	fmt.Println("密码：",pwd)

	writer.Write([]byte("欢迎你："+name))
}
