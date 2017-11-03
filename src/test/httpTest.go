package main

import (
	"fmt"
	"math/rand"
	"net/http"
	"log"
)

func main() {
	/* for 循环 */
	for a := 0; a < 10; a++ {
		var s int = rand.Intn(100)%3
		if s == 0 {
			fmt.Println("------> 小强")
		}else if s== 1 {
			fmt.Println("------> 牛肉")
		}else if s == 2{
			fmt.Println("------> 吃")
		}

	}

	fmt.Println("我启动监听了 http 8080 端口：http://localhost:8080")
	//step1. 建立 URL:Handler映射表
	mux := http.NewServeMux()
	/*mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		var  _url string = r.URL.RequestURI()
		if strings.Contains(_url,"test") {
			fmt.Fprintln(w, "Hello, world ---------for test")
		}else if strings.Contains(_url,"f") {
			fmt.Fprintln(w, "Hello, world ---------for fffffffffffffffffff")
		}else if strings.Contains(_url,"z"){
			fmt.Fprintln(w, "Hello, world ---------for zzzzzzzzzzzzzzzzzzzz")
		}else{
		fmt.Fprintln(w, "Hello, world ---------吃牛肉")
		}
	})*/
	mux.HandleFunc("/users", viewUsers)
	mux.HandleFunc("/companies", viewCompanies)

	//step2. 创建并运行HTTP server
	server := http.Server{Addr: ":8080", Handler: mux}
	log.Fatal(server.ListenAndServe())
}

func viewUsers(w http.ResponseWriter, r *http.Request) {

}

func viewCompanies(w http.ResponseWriter, r *http.Request) {

}
