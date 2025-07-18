package middleware

import (
	//"library/middleware"
	"os"
	"time"
	"fmt"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
)

type Claims struct {
    UserID  uint `json:"user_id"`
    IsAdmin bool `json:"is_admin"`
    //StandardClaims
}


var secret = os.Getenv("JWT_SECRET")


func GenerateJWT(userID, userRole string) (string,error) {

 claims := &jwt.MapClaims{
	 "userID": userID,
	 "user_role": userRole,
 "exp":time.Now().Add(24 * time.Hour).Unix(),
 }

token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)
tokenString,err := token.SignedString([]byte(secret))
if err != nil{
	return "", err
}
return tokenString, nil
	
}

  
func VerifyJWT(tokenString string) (*jwt.Token, error) {
	token, err := jwt.Parse(tokenString, func(t *jwt.Token) (any, error) {
		if _, ok := t.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", t.Header["alg"])
		}
		return []byte(secret), nil
	})
	if err != nil {
		return nil, err
	}
	return token, nil
}

func AuthMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		authHeader := r.Header.Get("Authorization")
		if authHeader == "" {
			http.Error(w, "missing token", http.StatusUnauthorized)
			return
		}
		tokenString := strings.TrimPrefix(authHeader, "Bearer ")

		_, err := VerifyJWT(tokenString)
		if err != nil {
			http.Error(w, "invalid token", http.StatusUnauthorized)
			return
		}
		next.ServeHTTP(w, r)
	})

}

