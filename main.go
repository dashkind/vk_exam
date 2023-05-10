package main

import (
	"fmt"
	"log"
	"os"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api"
)

func main() {
	bot, err := tgbotapi.NewBotAPI(os.Getenv("TELEGRAM_TOKEN"))
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = false

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 120

	updates, err := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.CallbackQuery == nil { // ignore non-CallbackQuery updates
			if update.Message != nil {
				chatId := update.Message.Chat.ID
				if update.Message.IsCommand() {
					if update.Message.Command() == "start" {
						msg := tgbotapi.NewMessage(chatId, fmt.Sprint("Привет ", update.Message.From.UserName, "!"))
						bot.Send(msg)
						msg = tgbotapi.NewMessage(chatId, fmt.Sprint("Это главное меню"))
						msg.ReplyMarkup = tgbotapi.NewInlineKeyboardMarkup(
							tgbotapi.NewInlineKeyboardRow(
								tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "button1_callback_data"),
								tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "button2_callback_data"),
							),
							tgbotapi.NewInlineKeyboardRow(
								tgbotapi.NewInlineKeyboardButtonData("Кнопка 3", "button3_callback_data"),
								tgbotapi.NewInlineKeyboardButtonData("Кнопка 4", "button4_callback_data"),
							),
						)
						bot.Send(msg)
					}
				}
			}
		}
		if update.CallbackQuery != nil {
			chatId := update.CallbackQuery.Message.Chat.ID
			messageId := update.CallbackQuery.Message.MessageID
			switch update.CallbackQuery.Data {
			case "button1_callback_data":
				newKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Сделай 1", "button11_callback_data"),
						tgbotapi.NewInlineKeyboardButtonData("Назад", "button_back_callback_data"),
					),
				)
				editMessage := tgbotapi.NewEditMessageText(chatId, messageId, "Вы нажали кнопку 1")
				editMessage.ReplyMarkup = &newKeyboard
				bot.Send(editMessage)
			case "button2_callback_data":
				newKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Сделай 2", "button21_callback_data"),
						tgbotapi.NewInlineKeyboardButtonData("Назад", "button_back_callback_data"),
					),
				)
				editMessage := tgbotapi.NewEditMessageText(chatId, messageId, "Вы нажали кнопку 2")
				editMessage.ReplyMarkup = &newKeyboard
				bot.Send(editMessage)
			case "button3_callback_data":
				newKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Сделай 3", "button31_callback_data"),
						tgbotapi.NewInlineKeyboardButtonData("Назад", "button_back_callback_data"),
					),
				)
				editMessage := tgbotapi.NewEditMessageText(chatId, messageId, "Вы нажали кнопку 3")
				editMessage.ReplyMarkup = &newKeyboard
				bot.Send(editMessage)
			case "button4_callback_data":
				newKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Сделай 4", "button41_callback_data"),
						tgbotapi.NewInlineKeyboardButtonData("Назад", "button_back_callback_data"),
					),
				)
				editMessage := tgbotapi.NewEditMessageText(chatId, messageId, "Вы нажали кнопку 4")
				editMessage.ReplyMarkup = &newKeyboard
				bot.Send(editMessage)
			case "button_back_callback_data":
				newKeyboard := tgbotapi.NewInlineKeyboardMarkup(
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Кнопка 1", "button1_callback_data"),
						tgbotapi.NewInlineKeyboardButtonData("Кнопка 2", "button2_callback_data"),
					),
					tgbotapi.NewInlineKeyboardRow(
						tgbotapi.NewInlineKeyboardButtonData("Кнопка 3", "button3_callback_data"),
						tgbotapi.NewInlineKeyboardButtonData("Кнопка 4", "button4_callback_data"),
					),
				)
				editMessage := tgbotapi.NewEditMessageText(chatId, messageId, "Это главное меню")
				editMessage.ReplyMarkup = &newKeyboard
				bot.Send(editMessage)
			}
		}
	}
}
