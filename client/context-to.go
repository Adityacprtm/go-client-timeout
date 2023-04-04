package main

import (
	"context"
	"fmt"
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	url := "http://localhost:8080/"

	// initial the context
	ctx, cncl := context.WithTimeout(context.Background(), time.Second*5)
	defer cncl()

	// initial the request with time out
	req, err := http.NewRequestWithContext(ctx, http.MethodGet, url, nil)

	if err != nil {
		fmt.Println(err)
		return
	}

	// do the request
	res, err := http.DefaultClient.Do(req.WithContext(ctx))

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
