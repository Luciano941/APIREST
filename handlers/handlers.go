package handlers

import (
	"encoding/json"
	"net/http"
	"packages/db"
	"packages/models"
	"strconv"

	"github.com/gorilla/mux"
)

func GetUserById(r *http.Request) (models.User, error) {
	vars := mux.Vars(r)
	userId, _ := strconv.Atoi(vars["id"])
	user := models.User{}

	if err := db.Database().First(&user, userId); err != nil {
		return user, err.Error
	} else {
		return user, nil
	}
}

func GetUsers(rw http.ResponseWriter, r *http.Request) {
	users := models.Users{}
	db.Database().Find(&users)
	SendData(rw, users, http.StatusOK)
}

func GetUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := GetUserById(r); err != nil {
		SendError(rw, http.StatusNotFound)
	} else {
		SendData(rw, user, http.StatusOK)
	}
}

func CreateUser(rw http.ResponseWriter, r *http.Request) {
	user := models.User{}
	decoder := json.NewDecoder(r.Body)

	if err := decoder.Decode(&user); err != nil {
		SendError(rw, http.StatusUnprocessableEntity)
	} else {
		db.Database().Save(&user)
		SendData(rw, user, http.StatusOK)
	}
}

func UpdateUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := GetUserById(r); err != nil {
		SendError(rw, http.StatusUnprocessableEntity)
	} else {
		newUser := models.User{}
		decoder := json.NewDecoder(r.Body)

		if err := decoder.Decode(&user); err != nil {
			SendError(rw, http.StatusUnprocessableEntity)
		} else {
			newUser.Id = user.Id
			db.Database().Save(&user)
			SendData(rw, newUser, http.StatusOK)
		}
	}
}

func DeleteUser(rw http.ResponseWriter, r *http.Request) {
	if user, err := GetUserById(r); err != nil {
		SendError(rw, http.StatusNotFound)
	} else {
		db.Database().Delete(&user)
		SendData(rw, user, http.StatusOK)
	}
}
