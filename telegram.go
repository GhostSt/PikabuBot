package main

import (
	"net/http"
	"fmt"
	"bytes"
)

func sendRequest(method string, parameters map[string]string) (*http.Response, error) {
	var urlBuffer bytes.Buffer

	url, err := reg.config.Get("telegram.url")

	if err != nil {
		return nil, err
	}

	token, err := reg.config.Get("telegram.token")

	if err != nil {
		return nil, err
	}

	urlBuffer.WriteString(url)
	urlBuffer.WriteString("bot")
	urlBuffer.WriteString(token)
	urlBuffer.WriteString("/")
	urlBuffer.WriteString(method)

	var queryBuffer bytes.Buffer

	parametersCount := len(parameters)

	i := 0
	for key := range parameters {
		queryParameter := key + "=" + parameters[key]

		queryBuffer.WriteString(queryParameter)

		if i != parametersCount - 1 {
			queryBuffer.WriteString("&")
		}

		i++
	}

	urlBuffer.WriteString("?")
	urlBuffer.Write(queryBuffer.Bytes())

	fmt.Println(urlBuffer.String())

	requestUrl := urlBuffer.String()

	req, _ := http.NewRequest("POST", requestUrl, nil)

	client := &http.Client{}

	res, err := client.Do(req)

	if err != nil {
		panic(err)
	}

	return res, nil
}

func sendMessage(message string) (bool, error){
	chatId, err := reg.config.Get("telegram.chat_id")

	if err != nil {
		return false, err
	}

	parseMode, err := reg.config.Get("telegram.send_message.parse_mode")

	if err != nil {
		return false, err
	}

	parameters := map[string]string{}
	parameters["chat_id"] = "-" + chatId
	parameters["text"] = message
	parameters["parse_mode"] = parseMode

	fmt.Println(parameters)

	response, err := sendRequest("sendMessage", parameters)

	fmt.Println(response)

	return true, nil
}
