package api


import (
    "chatt/internal/middlewares"
    "chatt/internal/models"
    "chatt/internal/storage"
    "net/http"
    "strings"
)


func Login(w http.ResponseWriter, message string) string {
	parts := strings.Split(message, " ")

    if len(parts) != 5 {
        return "Invalid command, use: login -u username -p password"
    }

    if parts[1] != "-u" || parts[3] != "-p" {
        return "Invalid command, use: login -u username -p password"
    }

    username := parts[2]
    password := parts[4]
    user := models.User{Username: username, Password: password}

    users, err := storage.LoadUsers()

    if err != nil {
        return "Error getting users"
    }

    for _, cur := range users.Users {
    	if user.Username == cur.Username && user.Password == cur.Password {
            jwtService := middlewares.NewJWTService()
            tokenString, err := jwtService.CreateToken(user.Username)

            if err != nil {
                return "Error generating token"
            }

            http.SetCookie(w, &http.Cookie{
                Name:  "jwtToken",
                Value: tokenString,
                Path:  "/",
            })

            return "<script>window.location.reload();</script>"
        } 
    }
    return "Invalid user or password"
}

func Logout(w http.ResponseWriter) string {
    http.SetCookie(w, &http.Cookie{
        Name:   "jwtToken",
        Value:  "",
        Path:   "/",
        MaxAge: -1, 
    })

    return "Successfully logged out"
}

