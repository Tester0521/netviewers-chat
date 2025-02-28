package api

import (
    "fmt"
    "net/http"
    "strings"
    "chatt/internal/middlewares"
)

func HandleHomePage(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path != "/" {
        http.NotFound(w, req)
        return
    }

    http.ServeFile(w, req, "web/index.html")
}

func HandleChatPage(w http.ResponseWriter, req *http.Request) {
    if req.URL.Path != "/nochat" {
        http.NotFound(w, req)
        return
    }
    http.ServeFile(w, req, "web/views/chat.html")
}

func HandleUsername(w http.ResponseWriter, req *http.Request) {
    jwt := middlewares.NewJWTService()

    username, _ := jwt.GetUsernameFromReq(req)

    fmt.Fprintf(w, `%s $`, username)
}

func HandleCmd(w http.ResponseWriter, req *http.Request) {
    
    message := req.FormValue("message")
    jwt := middlewares.NewJWTService()

    username, _ := jwt.GetUsernameFromReq(req)

    w.Header().Set("Content-Type", "text/html")

    if strings.HasPrefix(message, "login") {
        if jwt.CheckToken(w, req) {
            fmt.Fprintf(w, `<span>%s $ %s</span><br><span>Already logged in</span><br>`, username, message)
            return
        } 
        loginRes := Login(w, message)

        if strings.HasPrefix(loginRes, "Success:") {
            fmt.Fprintf(w, `<script>window.location.reload()</script>`)
        } else {
            fmt.Fprintf(w, `<span>%s $ %s</span><br><span>%s</span><br>`, username, message, loginRes)
        }
    } else if message == "logout" {
        if !jwt.CheckToken(w, req) {
            fmt.Fprintf(w, `<span>%s $ %s</span><br><span>U r not logged in</span><br>`, username, message)
            return
        }

        fmt.Fprintf(w, `%s <script>window.location.reload()</script>`, Logout(w))
    } else if strings.HasPrefix(message, "echo") {
        if len(strings.TrimSpace(message)) < 5 {
            fmt.Fprintf(w, `<span>%s $ %s</span><br>`, username, message)
            return
        }

        fmt.Fprintf(w, `<span>%s $ %s</span><br><span>%s</span><br>`, username, message, strings.TrimSpace(message[5:]))
    } else if message == "nochat" {
        if !jwt.CheckToken(w, req) {
            fmt.Fprintf(w, `<span>%s $ %s</span><br><span>U r not authorized</span><br>`, username, message)
            return
        }

        fmt.Fprintf(w, `<script>window.location.href = "/nochat"</script>`)
    } else if message == "clear" {
        fmt.Fprintf(w, `<script>document.getElementById("cout").innerHTML = ""</script>`)
    } else if message == "" {
        fmt.Fprintf(w, `<span>%s $ %s</span><br>`, username, message)
    } else {
        fmt.Fprintf(w, `<span>%s $ %s</span><br><span>bash: command not found: %s</span><br>`, username, message, message)
    }
}