package main

import (
	"bytes"
	"fmt"
	"net/http"
	"net/url"
)

func main() {
  urlStr:="http://127.0.0.1:9999/login"

  datas:=url.Values{
  	"username":{"刘义"},
  	"password":{"1111111"},
  }
  requestBody:=bytes.NewBufferString(datas.Encode())
  response,err:=http.Post(urlStr,"application/x-www-form-urlencoded",requestBody)
	if err != nil {
		fmt.Println(err)
	}
	defer response.Body.Close()
}
