package main

import (
	"fmt"
	"io"
	"net/http"
)

func main() {

	// fazer chamada http
	request, err := http.Get("https://ifconfig.me")
	if err != nil {
		panic(err)
	}

	response, err := io.ReadAll(request.Body)

	if err != nil {
		panic(err)
	}

	fmt.Println(string(response))

	request.Body.Close()
}
