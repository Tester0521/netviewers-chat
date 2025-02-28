package api


import (
    "net/http"
    "chatt/internal/middlewares"
)


func SetupRoutes() {
    jwt := middlewares.NewJWTService()

    http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("web/static"))))
    http.HandleFunc("/", HandleHomePage)
    http.HandleFunc("/nochat", jwt.AuthMiddleware(HandleChatPage))
    http.HandleFunc("/cmd", HandleCmd)
    http.HandleFunc("/get-username", HandleUsername)
    http.HandleFunc("/ws", HandleConnection)
}