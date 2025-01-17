package auth

import (
	"encoding/json"
	"io"
	"net/http"
	"ramada/api/src/db"
	"ramada/api/src/models"
)

type TokenData struct {
	Token string
}

type LoginData struct {
	Email    string
	Password string
}

func Login(writer http.ResponseWriter, request *http.Request) {
	var (
		db     = db.GetDB()
		user   *models.User
		_error error
	)

	bodyRequest, _error := io.ReadAll(request.Body)

	if _error != nil {
		writer.Write([]byte(_error.Error()))

		return
	}

	var loginData *LoginData

	_error = json.Unmarshal(bodyRequest, &loginData)

	if _error != nil {
		writer.Write([]byte("one or more login fields have errors"))

		return
	}

	db.Find(&user, models.User{Email: loginData.Email})
	_error = CheckPassword(loginData.Password, user.Password)

	if _error != nil {
		writer.Write([]byte("email or password is wrong"))

		return
	}

	token, _error := GenToken(user.ID)

	if _error != nil {
		writer.Write([]byte("login failed"))

		return
	}

	json.NewEncoder(writer).Encode(TokenData{Token: token})
}
