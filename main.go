package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"time"
)

// точка входа
func main() {
	botToken := "1374295300:AAEro6013Bvfu-LS5BoMvXZmA7qjzeU4cmE"
	botAPI := "https://api.telegram.org/bot"
	botURL := botAPI + botToken
	offset := 0
	for { // for ; ; { так записывается бесконечный цикл
		updates, err := getUpdates(botURL, offset)
		if err != nil {
			log.Println("Smth went wrong: ", err.Error())
		}

		for _, update := range updates {
			err = respond(botURL, update)
			offset = update.UpdateID + 1
		}
		fmt.Println(updates)

	}

}

// запрос обновлений
func getUpdates(botURL string, offset int) ([]Update, error) {
	resp, err := http.Get(botURL + "/getUpdates" + "?offset=" + strconv.Itoa(offset))
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		return nil, err
	}

	var restResponse RestResponse
	err = json.Unmarshal(body, &restResponse)
	if err != nil {
		return nil, err
	}
	return restResponse.Result, nil
}

// ответ на обновления
func respond(botURL string, update Update) error {
	var botMessage BotMessage
	T := time.Now()
	Tm := time.Unix(int64(update.Message.Date), 0)
	botMessage.ChatID = update.Message.Chat.ChatID
	botMessage.Text = T.Format("02 - 01 - 2006") + "\n" + update.Message.Text + "\n" + T.Format("15:04:05") + "\n" + Tm.Format("02 - 01 - 2006  15:04:05")
	buf, err := json.Marshal(botMessage)
	if err != nil {
		return err
	}

	_, err = http.Post(botURL+"/sendMessage", "application/json", bytes.NewBuffer(buf)) // application/json указатель на тип данных, используемых в запросе
	if err != nil {
		return err
	}
	return nil
}
