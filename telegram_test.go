package bot_test

import (
	"testing"
	"PikabuBot"
	"github.com/golang/mock/gomock"
	"PikabuBot/mocks"
	"fmt"
	"net/http"
	"net/url"
)

func TestTelegram_SendMessage(t *testing.T) {
	mockCtrl := gomock.NewController(t)
	defer mockCtrl.Finish()

	address := "http://localhost/"
	token := "some_token"
	chatId := "some_chat_id"
	parseMode := "some mode"
	method := "sendMessage"
	message := "some message"

	mockHttpClient := mock_bot.NewMockclient(mockCtrl)

	expectedUrl := fmt.Sprintf("%s%s%s/%s", address, "bot", token, method)
	expectedUrl += "?chat_id=" + chatId
	expectedUrl += "&text=" + url.QueryEscape(message)
	expectedUrl += "&parse_mode=" + url.QueryEscape(parseMode)

	mockHttpClient.EXPECT().SendRequest("POST", expectedUrl, nil).Return(&http.Response{}, nil).Times(1)

	telegram := &bot.Telegram{
		Client:    mockHttpClient,
		Address:   address,
		Token:     token,
		ChatId:    chatId,
		ParseMode: parseMode,
	}

	telegram.SendMessage("some message")
}
