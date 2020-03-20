package main

import (
    "fmt"
    "io/ioutil"
    "net/http"
)

func sendMessageHttpHandler(w http.ResponseWriter, req *http.Request) {
    chanelid := req.URL.Query().Get("chanelid")
    sign := req.URL.Query().Get("sign")
    if len(chanelid) == 0 || len(sign) == 0 {
        http.Error(w, "missing chanelid or sign", 400)
        return
    }
    if !signedStringCheck(chanelid, sign, Config.BotToken) {
        http.Error(w, "sign mismatch", 401)
        return
    }
    b, err := ioutil.ReadAll(req.Body)
    defer req.Body.Close()
    if err != nil {
        http.Error(w, err.Error(), 500)
        return
    }
    if len(b) == 0 {
        http.Error(w, "content is empty", 500)
        return
    }
    m, err := sendMessageToDiscord(chanelid, string(b))
    if err != nil {
        http.Error(w, "fail", 500)
        return
    }
    addMessageReaction(m, "✔️")//❌
    addMessageReaction(m, "🇭🇰")
    fmt.Fprintf(w, "ok")
}

func defaultHttpHandler(w http.ResponseWriter, req *http.Request) {
    fmt.Fprintf(w, "hello world!\n")
}