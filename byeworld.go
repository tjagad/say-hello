package main

import (
	"fmt"
	"net/http"
)

func byeWorldHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Bye World\n")
}
