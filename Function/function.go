package main

import (
	"fmt"
	"net/http"
)

func contentType(url string) (string, error) {
	resp, err := http.Get(url)

	if err != nil {
		return "", err
	}
	defer resp.Body.Close()

	return resp.Header.Get("Content-Type"), nil
}

func main() {
	resp, err := contentType("https://golang.org")

	if err != nil {
		fmt.Printf("Error: %s\n", err)
	} else {
		fmt.Println(resp)
	}
}
