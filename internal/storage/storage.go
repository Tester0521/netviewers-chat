package storage

import (
    "encoding/json"
    "os"
    "chatt/internal/models"
)


func LoadUsers() (models.UserStorage, error) {
    data, err := os.ReadFile("internal/db/users.json")
    if err != nil {
        return models.UserStorage{}, err
    }

    var users models.UserStorage
    if err := json.Unmarshal(data, &users); err != nil {
        return models.UserStorage{}, err
    }
    return users, nil
}

// func AppendMessageToFile(message Message, category string) error {
//     var chatData ChatData

//     file, err := os.OpenFile("internal/db/chat.json")

//     if err != nil {
//         return err
//     }

//     defer file.Close()

//     decoder := json.NewDecoder(file)

//     if err := decoder.Decode(&chatData); err != nil && err.Error() != "EOF" {
//         return err
//     }

//     if category == "vip" {
//         chatData.VIP = append(chatData.VIP, message)
//     } else {
//         chatData.Pub = append(chatData.Pub, message)
//     }

//     file.Truncate(0) 
//     file.Seek(0, 0) 

//     encoder := json.NewEncoder(file)
//     encoder.SetIndent("", "  ")

//     return encoder.Encode(chatData)
// }