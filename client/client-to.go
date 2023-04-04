package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {

	url := "http://localhost:8080/"
	method := "GET"

	client := &http.Client{
		// set client time out
		Timeout: 5 * time.Second,
	}
	req, err := http.NewRequest(method, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}
	res, err := client.Do(req)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer res.Body.Close()

	body, err := ioutil.ReadAll(res.Body)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(string(body))
}
