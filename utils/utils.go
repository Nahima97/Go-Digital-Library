package utils

import (
	"errors"
	"library/middleware"
	"net/http"
	"strings"

	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
)

//hash password with bcrypt
func HashPassword (password string) (string, error) {
hashedPass, err := bcrypt.GenerateFromPassword([]byte(password), bcrypt.DefaultCost)
if err != nil {
	return "", err
}
return string(hashedPass), nil
}

//compare passwords 
func ComparePassword (hashedPass, plainPass string) error {
err := bcrypt.CompareHashAndPassword([]byte(hashedPass), []byte(plainPass))
if err != nil {
	return err
}
return nil 
}

func ExtractUserID(w http.ResponseWriter, r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := middleware.VerifyJWT(tokenString)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userID, ok := claims["userID"].(string)
	if !ok {
		return "", errors.New("userID not found")
	}
	return userID, nil
}

func GetUserRole(w http.ResponseWriter, r *http.Request) (string, error) {
	authHeader := r.Header.Get("Authorization")
	tokenString := strings.TrimPrefix(authHeader, "Bearer ")

	token, err := middleware.VerifyJWT(tokenString)
	if err != nil {
		http.Error(w, "invalid token", http.StatusUnauthorized)
		return "", err
	}
	claims, ok := token.Claims.(jwt.MapClaims)
	if !ok {
		return "", errors.New("invalid token claims")
	}

	userRole, ok := claims["user_role"].(string)
	if !ok {
		return "", errors.New("userID not found")
	}
	return userRole, nil
}