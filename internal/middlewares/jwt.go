package middlewares

import (
    "fmt"
    "time"
    "github.com/golang-jwt/jwt/v5"
    "github.com/spf13/viper"
    "net/http"
)


type JWTService struct {
    SK []byte
}

func NewJWTService() *JWTService {
    viper.SetConfigName("config")        
    viper.SetConfigType("yaml")         
    viper.AddConfigPath("configs/")             

    if err := viper.ReadInConfig(); err != nil {
        return nil
    }

    sk := viper.GetString("jwt.sk")

    return &JWTService{
        SK: []byte(sk),
    }
}

func (s *JWTService) CreateToken(username string) (string, error) {
    claims := jwt.MapClaims{
        "username": username,
        "exp":      time.Now().Add(time.Hour * 24).Unix(), 
    }

    token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
    tokenString, err := token.SignedString(s.SK)

    if err != nil {
        return "", err
    }

    return tokenString, nil
}


func (s *JWTService) VerifyToken(tokenString string) error {
    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return s.SK, nil
    })

    if err != nil {
        return err
    }

    if !token.Valid {
        return fmt.Errorf("Unauthorized: Invalid JWT")
    }

    return nil
}

func (s *JWTService) CheckToken(w http.ResponseWriter, req *http.Request) bool {
    cookie, err := req.Cookie("jwtToken")
    if  err != nil {
        return false
    }
    
    err = s.VerifyToken(cookie.Value) 
    if err != nil {
        return false
    }

    return true
}

func (s *JWTService) GetUsernameFromToken(tokenString string) (string, error) {

    token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
        return s.SK, nil
    })
    if err != nil {
        return "", fmt.Errorf("failed to parse token: %v", err)
    }

    if !token.Valid {
        return "", fmt.Errorf("invalid token")
    }

    claims, ok := token.Claims.(jwt.MapClaims)
    if  !ok {
        return "", fmt.Errorf("invalid token claims")
    }
    username, ok := claims["username"].(string)
    if  !ok {
        return "", fmt.Errorf("username not found in token claims")
    }

    return username, nil
}

func (s *JWTService) GetUsernameFromReq(req *http.Request) (string, error) {
    cookie, err := req.Cookie("jwtToken")
    if err != nil {
        return `Anonymus@localhost`, fmt.Errorf("JWT cookie missing")
    }

    username, err := s.GetUsernameFromToken(cookie.Value)
    if err != nil {
        return `Anonymus@localhost`, fmt.Errorf("invalid JWT: %v", err)
    }

    return fmt.Sprintf("%s@netviewers.ru", username), nil
}

func (s *JWTService) AuthMiddleware(next http.HandlerFunc) http.HandlerFunc {
    return func(w http.ResponseWriter, req *http.Request) {   

        cookie, err := req.Cookie("jwtToken")

        if err != nil {
            http.Error(w, "Unauthorized: JWT cookie missing", http.StatusUnauthorized)
            return
        }

        err = s.VerifyToken(cookie.Value)
        if err != nil {
            http.Error(w, "Unauthorized: Invalid JWT", http.StatusUnauthorized)
            return
        }

        next(w, req)
    }
}