package models

type User struct {
    Username string `json:"username"`
    Password string `json:"password"`
}

type UserStorage struct {
    Users []User `json:"users"`
}