package main

import (
    "flag"
    "github.com/jinzhu/configor"
    "github.com/bwmarrin/discordgo"
    "fmt"
    "os"
)

var (
    ConfigPATH string
    DiscordSession *discordgo.Session
    Closed chan struct{}
)

// Variables used for command line parameters
var Config struct {
    BotToken string `default:"895444309:AADdfntqx8sOXV4qxusM34qVGDDgH1Ls6-C"`
    BotAPIUrl string `default:"http://123.95.96.103"`
    BotServerPort string `default:":80"`
}

func init() {
    flag.StringVar(&ConfigPATH, "c", "config.yml", "config file path")
    flag.Parse()
    if !IsFile(ConfigPATH) || !Exists(ConfigPATH) {
        fmt.Printf("Failed to find configuration %s\n", ConfigPATH)
        os.Exit(3)
        return
    }
    configor.Load(&Config, ConfigPATH)
}

func main() {
    Closed = make(chan struct{})
    go startDiscordBot()// start a thread
    go startHttpServer()// start another thread
    // to protected main thread
    for range Closed {
        close(Closed)
    }
}