package bot

import (
	"net/http"
	"net/url"
	"bytes"
)

type Telegram struct {
	Client    client
	Address   string
	Token     string
	ChatId    string
	ParseMode string
}

// Initializes Telegram structure
func Init() (*Telegram, error) {
	/**
	urlString, err := reg.config.Get("telegram.url")

	if err != nil {
		return nil, err
	}

	token, err := reg.env.get("TELEGRAM_BOT_TOKEN")

	chatId, err := reg.env.get("TELEGRAM_CHAT_ID")

	if err != nil {
		return nil, err
	}

	parseMode, err := reg.config.Get("telegram.send_message.parse_mode")

	if err != nil {
		return nil, err
	}

	if err != nil {
		return nil, err
	}
	*/

	urlString := "123"
	token := "token"
	chatId := "chat"
	parseMode := "parseMode"

	telegram := &Telegram{
		&Client{},
		urlString,
		token,
		chatId,
		parseMode,
	}

	return telegram, nil
}

// Creates url to send request
func (t *Telegram) createUrl(method string, parameters map[string]string) string {
	var urlBuffer bytes.Buffer

	urlBuffer.WriteString(t.Address)
	urlBuffer.WriteString("bot")
	urlBuffer.WriteString(t.Token)
	urlBuffer.WriteString("/")
	urlBuffer.WriteString(method)

	i := 0
	for key, value := range parameters {
		if i == 0 {
			urlBuffer.WriteString("?")
		} else {
			urlBuffer.WriteString("&")
		}

		value = url.QueryEscape(value)

		urlBuffer.WriteString(key + "=" + value)

		i++
	}

	return urlBuffer.String()
}

// Sends request to Telegram API
func (t *Telegram) sendRequest(method string, parameters map[string]string) (*http.Response, error) {
	url := t.createUrl(method, parameters)
	res, err := t.Client.SendRequest("POST", url, nil)

	if err != nil {
		panic(err)
	}

	return res, nil
}

// Sends message by chat by bot
func (t *Telegram) SendMessage(message string) (bool, error) {

	parameters := map[string]string{}
	parameters["chat_id"] = t.ChatId
	parameters["text"] = message
	parameters["parse_mode"] = t.ParseMode

	_, err := t.sendRequest("sendMessage", parameters)

	if err != nil {
		panic(err)
	}

	return true, nil
}
