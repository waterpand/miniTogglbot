package main

// Update : получает структуру от бота
type Update struct {
	UpdateID int     `json:"update_id"` // данная система тегов показывает какое поле структуры соответствует какому полю в json-файле
	Message  Message `json:"message"`
}

// Message : раскрытие структуры Message
type Message struct {
	Chat Chat   `json:"chat"`
	Text string `json:"text"`
}

// Chat : дальнейшее раскрытие структуры
type Chat struct {
	ChatID int `json:"id"`
}

// RestResponse : структура для получения ответа от бота
type RestResponse struct {
	Result []Update `json:"result"`
}

// BotMessage : посылаем сообщение ботом
type BotMessage struct {
	ChatID int    `json:"chat_id"`
	Text   string `json:"text"`
}
