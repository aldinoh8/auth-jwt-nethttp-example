package controller

import (
	"encoding/json"
	"example/helpers"
	"example/model"
	"example/repository"
	"net/http"

	"github.com/golang-jwt/jwt/v5"
	"github.com/julienschmidt/httprouter"
)

type User struct {
	Repository repository.User
}

func (u User) Register(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
			Detail:  err,
		})
		return
	}

	user := u.Repository.Register(newUser.Username, helpers.HashPassword(newUser.Password))

	WriteJSON(w, http.StatusCreated, map[string]interface{}{
		"message": "success register",
		"newUser": user,
	})
}

func (u User) Login(w http.ResponseWriter, r *http.Request, p httprouter.Params) {
	var newUser model.User
	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		WriteJSON(w, http.StatusBadRequest, ErrorResponse{
			Message: err.Error(),
		})
		return
	}

	// Find user
	user, err := u.Repository.FindByUsername(newUser.Username)
	if err != nil {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
			Message: "invalid username/password",
		})
		return
	}

	// compare hash
	isValidPassword := helpers.ValidatePassword(newUser.Password, user.Password)
	if !isValidPassword {
		WriteJSON(w, http.StatusUnauthorized, ErrorResponse{
			Message: "invalid username/password",
		})
		return
	}

	// token
	token := helpers.GenerateToken(jwt.MapClaims{
		"userId": user.Id,
		"pinAtm": "123456",
	})

	// respone
	WriteJSON(w, http.StatusOK, map[string]interface{}{
		"message": "success login",
		"token":   token,
	})
}
