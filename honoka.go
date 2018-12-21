package main

import (
  "log"
  "github.com/go-telegram-bot-api/telegram-bot-api"
)

func piket_review() string {
  return "tes review"
}

func piket_makan() string {
  return "tes makan"
}

func makan_makan() string {
  return "makan makan"
}

func main() {
  bot, err := tgbotapi.NewBotAPI("713405222:AAFVw4cobDesBWaMfobprSZN1kermofs7gI")
  if err != nil {
    log.Panic(err)
  }

  bot.Debug = true

  log.Printf("Authorized on account %s", bot.Self.UserName)

  u := tgbotapi.NewUpdate(0)
  u.Timeout = 60

  updates, err := bot.GetUpdatesChan(u)

  for update := range updates {
    if update.Message == nil || !update.Message.IsCommand() {
      continue
    }

    log.Printf("[%s] %s", update.Message.From.UserName, update.Message.Text)

    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "")

    switch update.Message.Command() {
    case "jajan":
      msg.Text = piket_makan()
    case "review":
      msg.Text = piket_review()
    case "makan":
      msg.Text = makan_makan()
    case "about":
      msg.Text = bot.Self
    case "start":
      msg.Text =
    default:
      msg.Text = "Honoka wakaranai desu~"
    }

    msg.ReplyToMessageID = update.Message.MessageID

    bot.Send(msg)
  }
}
