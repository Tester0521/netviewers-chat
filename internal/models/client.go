package models

import "github.com/gorilla/websocket"

type Client struct {
    Conn *websocket.Conn
    Name string
}

var Clients = make(map[*websocket.Conn]*Client)