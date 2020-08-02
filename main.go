package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strconv"
	"strings"
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

			if strings.HasPrefix(update.Message.Text, "%") {

				if strings.Contains((strings.ToLower(update.Message.Text)), "выход") || strings.Contains((strings.ToLower(update.Message.Text)), "exit") {
					hours, minutes := getCheck(update.Message.Date)
					update.Message.Text = "Время учтено\nПереработка за сегодня: " + strconv.Itoa(hours) + " час, " + strconv.Itoa(minutes) + " минут"
					fmt.Print("Переработка за сегодня составляет: ", hours, ":", minutes, "\n")
				} else {
					update.Message.Text = update.Message.Text + "\nВне учёта"
				}
			}

			err = respond(botURL, update)
			offset = update.UpdateID + 1
		}
		fmt.Println(updates)
		time.Sleep(1 * time.Second)

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
	var (
		botMessage BotMessage
	)
	T := time.Unix(int64(update.Message.Date), 0)
	T = T.Add((time.Hour * 3)) // Если бот запущен на Heroku - добавляет три часа к локальному времени
	botMessage.ChatID = update.Message.Chat.ChatID
	botMessage.Text = T.Format("02 - 01 - 2006") + "\n" + update.Message.Text + "\n" + T.Format("15:04:05")
	botMessage.Button.Keyboard[0][0].Text = "%Выход"
	botMessage.Button.Keyboard[0][1].Text = "%Тест"
	botMessage.Button.SizeKeyboard = true // этот параметр изменяет размер кнопки в боте под текст
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

// учет времени переработки
func getCheck(overTime int) (int, int) {
	var tHourRegime, tMinuteRegime int
	t := time.Unix(int64(overTime), 0)
	t = t.Add((time.Hour * 3)) // Если бот запущен на Heroku - добавляет три часа к локальному времени
	fmt.Println(t)

	switch t.Format("Mon") {
	case "Sun": // это для проверки при разработке
		tHourRegime = 7
		tMinuteRegime = 5
	case "Sat":
		tHourRegime = 7
		tMinuteRegime = 5
	case "Fri":
		tHourRegime = 14
		tMinuteRegime = 50
	default:
		tHourRegime = 15
		tMinuteRegime = 50
	}

	tHourExit, err := strconv.Atoi(t.Format("15"))
	if err != nil {
		fmt.Println(err)
	}
	tMinuteExit, err := strconv.Atoi(t.Format("04"))
	if err != nil {
		fmt.Println(err)
	}

	if tMinuteExit < tMinuteRegime && tHourExit <= tHourRegime {
		return 0, 0
	}
	if tMinuteExit < tMinuteRegime && tHourExit > tHourRegime {
		tHourExit = tHourExit - 1
		tMinuteExit = tMinuteExit + 60
	}

	tOverHours := tHourExit - tHourRegime
	tOverMinutes := tMinuteExit - tMinuteRegime

	return tOverHours, tOverMinutes
}
