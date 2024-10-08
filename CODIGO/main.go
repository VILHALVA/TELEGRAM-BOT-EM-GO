package main

import (
    "log"
    "os"
    "time"

    "github.com/go-telegram-bot-api/telegram-bot-api/v5"
    "github.com/joho/godotenv"
)

func main() {
    err := godotenv.Load()
    if err != nil {
        log.Fatalf("Erro ao carregar arquivo .env: %v", err)
    }

    token := os.Getenv("TOKEN")
    if token == "" {
        log.Fatal("Token do Telegram não encontrado. Verifique o arquivo .env.")
    }

    bot, err := tgbotapi.NewBotAPI(token)
    if err != nil {
        log.Fatalf("Erro ao criar bot: %v", err)
    }

    log.Printf("Autenticado como %s", bot.Self.UserName)

    updateConfig := tgbotapi.NewUpdate(0)
    updateConfig.Timeout = 60
    updates := bot.GetUpdatesChan(updateConfig)

    for update := range updates {
        if update.Message == nil { 
            continue
        }

        if update.Message.IsCommand() {
            switch update.Message.Command() {
            case "start":
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "😃Bem-vindo ao bot! 😁Envie /help para ver os comandos!")
                if _, err := bot.Send(msg); err != nil {
                    log.Printf("Erro ao enviar mensagem: %v", err)
                }
            case "help":
                helpMsg := "Comandos disponíveis:\n" +
                    "/start - Iniciar o bot\n" +
                    "/help - Mostrar esta mensagem de ajuda\n" +
                    "/about - Informações sobre o bot\n" +
                    "/echo <mensagem> - Repetir a mensagem que você enviou\n" +
                    "/date - Mostrar a data e hora atuais"
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, helpMsg)
                if _, err := bot.Send(msg); err != nil {
                    log.Printf("Erro ao enviar mensagem: %v", err)
                }
            case "about":
                aboutMsg := "Este bot foi criado para demonstrar a API do Telegram com Go. " +
                    "Você pode usar comandos como /help, /echo e /date."
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, aboutMsg)
                if _, err := bot.Send(msg); err != nil {
                    log.Printf("Erro ao enviar mensagem: %v", err)
                }
            case "echo":
                if len(update.Message.CommandArguments()) > 0 {
                    echoMsg := update.Message.CommandArguments()
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, echoMsg)
                    if _, err := bot.Send(msg); err != nil {
                        log.Printf("Erro ao enviar mensagem: %v", err)
                    }
                } else {
                    msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Por favor, forneça uma mensagem para repetir.")
                    if _, err := bot.Send(msg); err != nil {
                        log.Printf("Erro ao enviar mensagem: %v", err)
                    }
                }
            case "date":
                currentTime := time.Now().Format("02/01/2006 15:04:05")
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Data e hora atuais: "+currentTime)
                if _, err := bot.Send(msg); err != nil {
                    log.Printf("Erro ao enviar mensagem: %v", err)
                }
            default:
                msg := tgbotapi.NewMessage(update.Message.Chat.ID, "Comando não reconhecido. Use /help para ver os comandos disponíveis.")
                if _, err := bot.Send(msg); err != nil {
                    log.Printf("Erro ao enviar mensagem: %v", err)
                }
            }
        }
    }
}
