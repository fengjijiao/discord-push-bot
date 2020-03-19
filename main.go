package main

import (
    "flag"
    "github.com/jinzhu/configor"
    "github.com/bwmarrin/discordgo"
    //"net/http"
)

var (
    ConfigPATH string
    DiscordSession *discordgo.Session
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
    configor.Load(&Config, ConfigPATH)
}

func main() {
    startDiscordBot()
    startHttpServer()
}