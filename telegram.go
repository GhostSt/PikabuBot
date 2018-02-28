package main

import (
	"net/http"
)

func sendRequest(url string, data []int) *http.Response {
	req, _ := http.NewRequest("POST", url, nil)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return res
}
