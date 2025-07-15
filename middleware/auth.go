package middleware

import (

	//"library/middleware"

	"os"
	"time"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID  uint `json:"user_id"`
    IsAdmin bool `json:"is_admin"`
    //StandardClaims
}

var secret = os.Getenv("JWT_SECRET")


func GenerateJWT(userID string) (string,error) {

 claims := &jwt.MapClaims{
	 "userID": userID,
 "exp":time.Now().Add(24 * time.Hour).Unix(),
 }

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString,err := token.SignedString(secret)
if err != nil{
	return "", err
}
return tokenString, nil
	
}

func VerifyJWT() {

}

func AuthMiddleware() {

}