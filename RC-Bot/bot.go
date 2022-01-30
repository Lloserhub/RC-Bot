package main

import (
	"log"
	"strings"

	tgbotapi "github.com/go-telegram-bot-api/telegram-bot-api/v5"
)

func main() {
	bot, err := tgbotapi.NewBotAPI("TOKEN HERE")
	if err != nil {
		log.Panic(err)
	}

	bot.Debug = true

	log.Printf("Authorized on account %s", bot.Self.UserName)

	u := tgbotapi.NewUpdate(0)
	u.Timeout = 60

	updates := bot.GetUpdatesChan(u)

	for update := range updates {
		if update.Message != nil {

			text := strings.ToLower(update.Message.Text)

			switch text {

			case "/start":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Привет, "+update.SentFrom().FirstName+
					". Ты написал нашему боту, который может помочь тебе справиться с буллингом"+
					".\nВведи команду /help чтобы увидеть возможные команды.")

				bot.Send(msg)
				continue

			case "/help":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "\n\n*Введи эту команду, чтобы получить доступ к номерам экстренных служб: /emergency"+
					"\n\n*Введи эту команду, чтобы получить доступ к информации и ресурсам, связанными с буллингом: /bullying"+
					"\n\n*Введи эту команду, чтобы получить доступ к информации и ресурсам, связанными с кибер буллингом: /cyberbullying"+
					"\n\n*Введи эту команду, чтобы связаться с одним из наших волонтёров: /volunteer")

				bot.Send(msg)
				continue

			case "/emergency":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"\n\n* Полиция: 102"+
						"\n* Национальная детская горячая линия: 0800500225 или 116111"+
						"\n* Офис бесплатной правовой помощи: 0800213103"+
						"\n* Аварийная газовая служба: 104")

				bot.Send(msg)
				continue

			case "/bullying":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Булинг - это намеренное повторяющееся поведение, которое"+
					" проявляется в виде: травли, бойкота, насмешек, распространения слухов, высмеивания, порчи имущества, "+
					"физического насилия, запугивания, вымогания денег и т.д.\n\n"+
					"Кто может стать жертвой буллинга? Люди с определенными особенностями, такими как:"+
					"\n*Внешний вид\n*Поведенческие особенности\n*Расовая принадлежность\n*Религия\n*Инвалидность или определенные болезни\n*Сексуальная ориентация\n\n"+
					"Но жертвой может стать кто-угодно.")

				bot.Send(msg)
				continue

			case "/cyberbullying":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Кибербуллинг - это неприятные сообщения, комментарии, ссылки на твой профиль в унижающих постах."+
						" Помни, что всё, что ты делаешь в интернете, остается в нем навсегда.")

				bot.Send(msg)
				continue

			case "/volunteer":
				msg := tgbotapi.NewMessage(update.Message.Chat.ID,
					"Помни, что это не твоя вина. Никто не заслуживает унижения и ничто не оправдывает буллинг. "+
						"Пострадавшие от насилия никогда не бывают виноватыми в его совершении. Они не могут"+
						" его спровоцировать, вызвать или обусловить. Виноват и должен отвечать всегда только насильник.")

				bot.Send(msg)
				continue
			}
		}
	}
}
