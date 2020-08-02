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
	Date int    `json:"date"`
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
	ChatID int                 `json:"chat_id"`
	Text   string              `json:"text"`
	Button ReplyKeyboardMarkup `json:"reply_markup"`
}

// ReplyKeyboardMarkup : структура для кнопки
type ReplyKeyboardMarkup struct {
	Keyboard     [1][2]KeyboardButton `json:"keyboard"`
	SizeKeyboard bool                 `json:"resize_keyboard"`
}

// KeyboardButton : текст кнопки
type KeyboardButton struct {
	Text string `json:"text"`
}

//BotButton : Попытка создать кнопку
type BotButton struct {
	Text         string `json:"text"`
	CallbackData string `json:"callback_data"`
}
