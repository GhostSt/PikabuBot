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

type ParameterBag struct {
	key   string
	value string
}

// Initializes Telegram structure
func Init(registry *Registry) (*Telegram, error) {
	urlString, err := registry.config.Get("telegram.url")

	if err != nil {
		return nil, err
	}

	token, err := registry.env.get("TELEGRAM_BOT_TOKEN")

	if err != nil {
		return nil, err
	}

	chatId, err := registry.env.get("TELEGRAM_CHAT_ID")

	if err != nil {
		return nil, err
	}

	parseMode, err := registry.config.Get("telegram.send_message.parse_mode")

	if err != nil {
		return nil, err
	}

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
func (t *Telegram) createUrl(method string, parameters []*ParameterBag) string {
	var urlBuffer bytes.Buffer

	urlBuffer.WriteString(t.Address)
	urlBuffer.WriteString("bot")
	urlBuffer.WriteString(t.Token)
	urlBuffer.WriteString("/")
	urlBuffer.WriteString(method)

	i := 0
	for _, parameter := range parameters {

		if i == 0 {
			urlBuffer.WriteString("?")
		} else {
			urlBuffer.WriteString("&")
		}

		value := url.QueryEscape(parameter.value)

		urlBuffer.WriteString(parameter.key + "=" + value)

		i++
	}

	return urlBuffer.String()
}

// Sends request to Telegram API
func (t *Telegram) sendRequest(method string, parameters []*ParameterBag) (*http.Response, error) {
	url := t.createUrl(method, parameters)
	res, err := t.Client.SendRequest("POST", url, nil)

	if err != nil {
		panic(err)
	}

	return res, nil
}

// Sends message by chat by bot
func (t *Telegram) SendMessage(message string) (bool, error) {

	chat := &ParameterBag{"chat_id", t.ChatId}
	text := &ParameterBag{"text", message}
	parseMode := &ParameterBag{"parse_mode", t.ParseMode}

	parameters := [...]*ParameterBag{chat, text, parseMode}

	_, err := t.sendRequest("sendMessage", parameters[:])

	if err != nil {
		panic(err)
	}

	return true, nil
}
