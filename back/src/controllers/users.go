package controllers

import (
	"encoding/json"
	"fmt"
	"net/http"
	"ramada/api/src/auth"
	"ramada/api/src/db"
	"ramada/api/src/models"
	"ramada/api/src/utils"

	"github.com/gorilla/mux"
)

func GetCurrentUser(writer http.ResponseWriter, request *http.Request) {
	var (
		db    = db.GetDB()
		model *models.User
	)

	tokenData, _error := auth.GetTokenData(request)

	if _error != nil {
		http.Error(writer, "error getting user data", http.StatusOK)

		return
	}

	id := fmt.Sprintf("%v", tokenData["userID"])
	db.Find(model, models.User{ID: *utils.StrToUint(id)})
	utils.DBClose(db)
	json.NewEncoder(writer).Encode(model)
}

func GetUser(writer http.ResponseWriter, request *http.Request) {
	var (
		db   = db.GetDB()
		user *models.User
		id   = *utils.StrToUint(mux.Vars(request)["id"])
	)

	if db.Find(&user, models.User{ID: id}); user.ID == 0 {
		utils.DBClose(db)
		http.Error(writer, "user not exists", http.StatusOK)

		return
	}

	json.NewEncoder(writer).Encode(user)
}

func ListUsers(writer http.ResponseWriter, request *http.Request) {
	var (
		db    = db.GetDB()
		users *[]models.User
	)

	if find := db.Find(&users); find.RowsAffected == 0 {
		utils.DBClose(db)
		http.Error(writer, "there are no users here", http.StatusOK)

		return
	}

	json.NewEncoder(writer).Encode(users)
}

func UpinsertUser(writer http.ResponseWriter, request *http.Request) {
	var (
		user        map[string]interface{}
		model       *models.User
		id          uint64 = 0
		db                 = db.GetDB()
		requestJSON        = utils.RequestToJSON(writer, *request)
	)

	hashed, _ := auth.HashIt(fmt.Sprintf("%v", requestJSON["Password"]))

	user = map[string]interface{}{
		"Name":     requestJSON["Name"],
		"Username": requestJSON["Username"],
		"Email":    requestJSON["Email"],
		"Password": hashed,
	}

	if idString := mux.Vars(request)["id"]; idString != "" {
		id = *utils.StrToUint(idString)
	}

	if id != 0 {
		db.Find(&model, models.User{ID: id})
	}

	if model != nil && model.ID != 0 {
		db.Model(model).Updates(user)
		utils.DBClose(db)
		json.NewEncoder(writer).Encode(model)

		return
	}

	db.Model(model).Create(user)
	utils.DBClose(db)
	json.NewEncoder(writer).Encode(model)
}

func DeleteUser(writer http.ResponseWriter, request *http.Request) {
	var (
		db    = db.GetDB()
		model *models.User
		id    = *utils.StrToUint(mux.Vars(request)["id"])
	)

	db.First(&model, models.User{ID: id})

	if model.ID == 0 {
		http.Error(writer, "user not exists", http.StatusOK)
		utils.DBClose(db)

		return
	}

	db.Delete(model, id)
	utils.DBClose(db)

	writer.WriteHeader(http.StatusOK)
	json.NewEncoder(writer).Encode(model)
}
