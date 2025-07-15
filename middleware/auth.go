package middleware

import (
	"fmt"

	"github.com/golang-jwt/jwt/v5"
)

//nahima, alina, and sulema
func GenerateJWT () {

}


func VerifyJWT (tokenString string) (*jwt.Token, error) {
token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
	if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %v\n", t.Header["alg"])
	}
	return []byte(secret), nil 
})
if err != nil {
	return nil, err
}
return token, nil
}


func AuthMiddleware () {



	
}


func GetUserbyUserID () {



}