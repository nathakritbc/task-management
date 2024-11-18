package  utils

import (
	"github.com/golang-jwt/jwt/v5"
	"time"
)

func GenerateToken(userID uint) (string, error) {
	token := jwt.New(jwt.SigningMethodHS256)
	claims := token.Claims.(jwt.MapClaims)
	claims["user_id"] = userID
	claims["exp"] = time.Now().Add(time.Hour * 24).Unix()
	
	return token.SignedString([]byte("your_jwt_secret"))
}