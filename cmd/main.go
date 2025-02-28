



package main

import (
    "fmt"
    "net/http"
    "chatt/internal/api"
)

func main() {
    api.SetupRoutes()

    
    fmt.Println("Сервер запущен на http://localhost:8080")

    if err := http.ListenAndServe(":8080", nil); err != nil {
        fmt.Println("Ошибка при запуске сервера:", err)
    } 
}








// package main

// import (
//     "encoding/json"
//     "fmt"
//     "net/http"
//     "sync"
//     "github.com/gorilla/websocket"
// )

// var upgrader = websocket.Upgrader{
//     CheckOrigin: func(r *http.Request) bool {
//         return true
//     },
// }

// type Client struct {
//     Conn *websocket.Conn 
//     Name string
// }

// var clients = make(map[*websocket.Conn]*Client)
// var mu sync.Mutex

// func loginHandler(w http.ResponseWriter, req *http.Request) {
//     var user struct {
//         Name string `json:"name"`
//     }

//     if err := json.NewDecoder(req.Body).Decode(&user); err != nil {
//         http.Error(w, "Invalid request", http.StatusBadRequest)
//         return
//     }

//     http.SetCookie(w, &http.Cookie{
//         Name:  "name",
//         Value: user.Name,
//         Path:  "/",
//     })

//     w.WriteHeader(http.StatusOK)
// }

// func handleConnection(w http.ResponseWriter, req *http.Request) {
//     conn, err := upgrader.Upgrade(w, req, nil)

//     if err != nil {
//         fmt.Println("Error while upgrading connection:", err)
//         return
//     }

//     defer conn.Close()

//     cookie, err := req.Cookie("name")
//     if err != nil {
//         fmt.Println("Error get cookie: ", err)
//         return
//     }

//     client := &Client{ Conn: conn, Name: cookie.Value }

//     mu.Lock()
//     clients[conn] = client
//     mu.Unlock()

//     for {
//         var msg []byte
//         _, msg, err := conn.ReadMessage();
        
//         if err != nil {
//             fmt.Println("Error while reading message:", err)
//             mu.Lock()
//             delete(clients, conn)
//             mu.Unlock()
//             break
//         }
//         broadcastMessage(client.Name, string(msg))
//     }
// }

// func broadcastMessage(name, msg string) {
//     mu.Lock()
//     defer mu.Unlock()
//     for client := range clients {
//         err := client.WriteMessage(websocket.TextMessage, []byte(fmt.Sprintf("%s: %s", name, msg)))
//         if err != nil {
//             fmt.Println("Error while sending message:", err)
//             client.Close()
//             delete(clients, client)
//         }
//     }
// }

// func main() {
    //     http.HandleFunc("/login", loginHandler)
    // http.HandleFunc("/", func(w http.ResponseWriter, req *http.Request) { http.ServeFile(w, req, "web/index.html") })
    // http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))

    // http.HandleFunc("/ws", handleConnection)
// }