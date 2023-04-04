package main

import (
	"fmt"
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", TimeoutFunc)
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		log.Fatalf("Error server %s\n", err.Error())
	}
}
func TimeoutFunc(w http.ResponseWriter, r *http.Request) {
	// sleep for 10s
	time.Sleep(time.Second * 10)
	fmt.Fprintf(w, "Hello world!")
}
