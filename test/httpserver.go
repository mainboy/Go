package main

import (
	"fmt"
	"log"
	"net/http"
)

type MyMux struct {
}

func (m *MyMux) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path == "/hello" {
		sayhelloName(w, r)
		return
	}
	http.NotFound(w, r)
	return
}

func sayhelloName(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello everybody!") //这个写入到w的是输出到客户端的
}

func main() {
	mux := &MyMux{}
	err := http.ListenAndServe(":9090", mux) //设置监听的端口
	if err != nil {
		log.Fatal("ListenAndServe: ", err)
	}
}
