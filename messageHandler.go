package main

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
)

// This function will be called (due to AddHandler above) every time a new
// message is created on any channel that the autenticated bot has access to.
func messageCreate(s *discordgo.Session, m *discordgo.MessageCreate) {
    // Ignore all messages created by the bot itself
    // This isn't required in this specific example but it's a good practice.
    if m.Author.ID == s.State.User.ID {
        return
    }
    // If the message is "ping" reply with "Pong!"
    if m.Content == "ping" {
        s.ChannelMessageSend(m.ChannelID, "Pong!")
    }

    // If the message is "pong" reply with "Ping!"
    if m.Content == "pong" {
        s.ChannelMessageSend(m.ChannelID, "Ping!")
    }

    // If the message is "getmyid" reply with "Ping!"
    if m.Content == "getchannelid" {
        s.ChannelMessageSend(m.ChannelID, "ChannelID:" + m.ChannelID)
    }

    if m.Content == "getpushurl" {
        s.ChannelMessageSend(m.ChannelID, sendMessageURLGen(m.ChannelID))
    }
}

func sendMessageURLGen(chanelid string) string {
    sign := stringSign(chanelid, Config.BotToken)
    return fmt.Sprintf("%s/send?chanelid=%s&sign=%s", Config.BotAPIUrl, chanelid, sign)
}

func sendMessageToDiscord(chanelid string, content string) (*discordgo.Message, error) {
    m,err := DiscordSession.ChannelMessageSend(chanelid, content)
    if err != nil {
        fmt.Println("send message was wrong!")
        return nil, err
    }
    return m, nil
}

func addMessageReaction(m *discordgo.Message, emojiID string) error {
    err := DiscordSession.MessageReactionAdd(m.ChannelID, m.ID, emojiID)
    return err
}