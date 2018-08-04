package main

import (
	"net/http"
	"net/url"
	"fmt"
	"bytes"
)

// Sends request to Telegram API
func sendRequest(method string, parameters map[string]string) (*http.Response, error) {
	var urlBuffer bytes.Buffer

	urlString, err := registry.config.Get("telegram.url")

	if err != nil {
		return nil, err
	}

	token, err := registry.env.get("TELEGRAM_BOT_TOKEN")

	if err != nil {
		return nil, err
	}

	urlBuffer.WriteString(urlString)
	urlBuffer.WriteString("bot")
	urlBuffer.WriteString(token)
	urlBuffer.WriteString("/")
	urlBuffer.WriteString(method)

	var queryBuffer bytes.Buffer

	parametersCount := len(parameters)

	i := 0
	for key := range parameters {
		value := parameters[key]
		value = url.QueryEscape(value)

		queryParameter := key + "=" + value

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

// Sends message by chat by bot
func sendMessage(message string) (bool, error){
	chatId, err := registry.env.get("TELEGRAM_CHAT_ID")

	if err != nil {
		return false, err
	}

	parseMode, err := registry.config.Get("telegram.send_message.parse_mode")

	if err != nil {
		return false, err
	}

	parameters := map[string]string{}
	parameters["chat_id"] = chatId
	parameters["text"] = message
	parameters["parse_mode"] = parseMode

	fmt.Println(parameters)

	response, err := sendRequest("sendMessage", parameters)

	fmt.Println(response)

	return true, nil
}
