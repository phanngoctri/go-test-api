package main

import (
	"apis/blog_api"
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	//fmt.Println("Hello world!")
	//handlers := http.NewServeMux()
	router := mux.NewRouter()

	router.HandleFunc("/api/blog/findall", blog_api.FindAll()).Methods("GET")

	err := http.ListenAndServe(":5000", router)
	if err != nil {
		fmt.Println(err)
	}
}
