package models


type Message struct {
    Name string `json:"name"`
    Msg  string `json:"msg"`
}

type ChatData struct {
    Vip []Message `json:"vip"`
    Pub []Message `json:"pub"`
}

