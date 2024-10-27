package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Minute}
	res, err := c.Get("https://google.com")

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	body, err := io.ReadAll(res.Body)
	if err != nil {
		panic(err)
	}

	println(string(body))

}

func Post() {
	c := http.Client{Timeout: time.Minute}
	jsonVar := bytes.NewBuffer([]byte(`{"name": "Alex"}`))
	res, err := c.Post("https://google.com", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}
	defer res.Body.Close()

	io.CopyBuffer(os.Stdout, res.Body, nil)
}
