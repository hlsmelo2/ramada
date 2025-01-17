package auth

import (
	"crypto/rand"
	"encoding/base64"
	"errors"
	"fmt"
	"log"
	"net/http"
	"ramada/api/src/config"
	"strings"
	"time"

	jwt "github.com/golang-jwt/jwt/v5"
)

func GetSecret() string {
	key := make([]byte, 64)

	if _, _error := rand.Read(key); _error != nil {
		log.Fatal(_error)
	}

	return base64.StdEncoding.EncodeToString(key)
}

func GenToken(userID uint64) (string, error) {
	data := jwt.MapClaims{}

	data["authorized"] = true
	data["exp"] = time.Now().Add(time.Hour * 6).Unix()
	data["userID"] = userID

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, data)

	return token.SignedString([]byte(config.SECRET_KEY))
}

func InvalidateToken(writer http.ResponseWriter, request *http.Request) {
	stringToken := getHeaderToken(request)
	tokenData, _error := GetTokenData(request)

	if _error != nil {
		writer.Write([]byte(_error.Error()))

		return
	}

	tokenData["exp"] = time.Now().Truncate(time.Hour).Unix()

	token, _error := jwt.ParseWithClaims(stringToken, tokenData, checkSigningMethod)

	if _error != nil {
		writer.Write([]byte(_error.Error()))

		return
	}

	tokenSign, _ := token.SignedString([]byte(config.SECRET_KEY))

	writer.Write([]byte(tokenSign))
}

func getHeaderToken(request *http.Request) string {
	token := request.Header.Get("Authorization")
	splitedToken := strings.Split(token, " ")

	if len(splitedToken) == 2 {
		return splitedToken[1]
	}

	return ""
}

func checkSigningMethod(token *jwt.Token) (interface{}, error) {
	if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
		return nil, fmt.Errorf("unexpected signing method: %s", token.Header["alg"])
	}

	return []byte(config.SECRET_KEY), nil
}

func GetTokenData(request *http.Request) (jwt.MapClaims, error) {
	headerToken := getHeaderToken(request)
	token, _error := jwt.Parse(headerToken, checkSigningMethod)

	if _error != nil {
		return nil, _error
	}

	tokenData, ok := token.Claims.(jwt.MapClaims)

	if !ok || !token.Valid {
		return nil, errors.New("error when getting data")
	}

	return tokenData, nil
}

func ValidateToken(request *http.Request) error {
	if _, _error := GetTokenData(request); _error == nil {
		return nil
	}

	return errors.New("invalid token")
}
