package main

import (
    "github.com/bwmarrin/discordgo"
    "fmt"
    //"os"
    //"os/signal"
    //"syscall"
    "net/http"
)

func startDiscordBot() {
    // Create a new Discord session using the provided bot token.
    dg, err := discordgo.New("Bot " + Config.BotToken)
    if err != nil {
        fmt.Println("error creating Discord session,", err)
        return
    }
    DiscordSession = dg
    // Register the messageCreate func as a callback for MessageCreate events.
    dg.AddHandler(messageCreate)

    // Open a websocket connection to Discord and begin listening.
    err = dg.Open()
    if err != nil {
        fmt.Println("error opening connection,", err)
        return
    }

    // Wait here until CTRL-C or other term signal is received.
    /*fmt.Println("Bot is now running.  Press CTRL-C to exit.")
    sc := make(chan os.Signal, 1)
    signal.Notify(sc, syscall.SIGINT, syscall.SIGTERM, os.Interrupt, os.Kill)
    <-sc

    // Cleanly close down the Discord session.
    dg.Close()*/
}

func startHttpServer() {
    http.HandleFunc("/send", sendMessageHttpHandler)
    http.HandleFunc("/", defaultHttpHandler)
    http.ListenAndServe(Config.BotServerPort, nil)
}