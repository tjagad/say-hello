package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	port := 9080

	http.HandleFunc("/helloworld", helloWorldHandler)
	http.HandleFunc("/byeworld", byeWorldHandler)

	log.Printf("Server starting on port %v\n", port)
	log.Fatal(http.ListenAndServe(fmt.Sprintf(":%v", port), nil))
}
