package api

import (
    "fmt"
    "net/http"
    "sync"
    "github.com/gorilla/websocket"
    "chatt/internal/models"
    "chatt/internal/middlewares"
)

var upgrader = websocket.Upgrader{
    CheckOrigin: func(req *http.Request) bool {
        return true
    },
}
var mu sync.Mutex

func HandleConnection(w http.ResponseWriter, req *http.Request) {
    conn, err := upgrader.Upgrade(w, req, nil)

    if err != nil {
        fmt.Println("Error while upgrading connection:", err)
        return
    }
    defer conn.Close()

    jwt := middlewares.NewJWTService()

    username, _ := jwt.GetUsernameFromReq(req)


    mu.Lock()
    models.Clients[conn] = &models.Client{Conn: conn, Name: username}
    mu.Unlock()

    for {
        var msg []byte
        _, msg, err := conn.ReadMessage()
        if err != nil {
            fmt.Println("Error while reading message:", err)
            mu.Lock()
            delete(models.Clients, conn)
            mu.Unlock()
            break
        }
        jwt := middlewares.NewJWTService()
        if !jwt.CheckToken(w, req) {
            fmt.Println("Unauthorized")
            return
        }

        BroadcastMessage(msg, username)
    }
}

func BroadcastMessage(msg []byte, username string) {
            mu.Lock()
    defer   mu.Unlock()

    for conn := range models.Clients {
        err := conn.WriteMessage(websocket.TextMessage, []byte (fmt.Sprintf(`%s: %s`, username, msg)))
        if err != nil {
            fmt.Println("Error while sending message:", err)
            conn.Close()
            delete(models.Clients, conn)
        }
    }
}
